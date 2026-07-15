package database

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/friendsofgo/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pinax-network/golang-base/log"
	"go.uber.org/zap"
)

const (
	// dialTimeout bounds how long establishing a TCP connection to a node may take.
	// Without it a dial to a dead/blackholed node blocks until the OS TCP timeout
	// (minutes), which stalls both queries and the health-check loop.
	dialTimeout = 5 * time.Second
	// ioTimeout bounds individual read/write operations on an established connection.
	ioTimeout = 5 * time.Second
	// healthCheckTimeout bounds a single node health check so the ping loop can never
	// block indefinitely on an unresponsive node.
	healthCheckTimeout = 3 * time.Second
	// maxConnLifetime recycles pooled connections so a broken connection to a node that
	// went away is not reused indefinitely.
	maxConnLifetime = 5 * time.Minute
	// maxConnIdleTime closes idle connections, forcing a fresh (timeout-bounded) dial
	// on the next use rather than reusing a possibly-dead idle connection.
	maxConnIdleTime = 1 * time.Minute
)

type MysqlConnectionPool struct {
	Connections []*MysqlConnection
	Mutex       *sync.Mutex
	PingsTicker *time.Ticker
	PingsDone   chan bool
	Config      *ClusterConfig
}

type MysqlConnection struct {
	Name   string
	Dsn    string
	Config *MysqlConnectionOptions
	// DB and IsActive are written by the health-check goroutine and read concurrently
	// by request goroutines via GetConnection/getActive, so both are accessed atomically
	// rather than under Mutex to stay race-free without holding a lock across network I/O.
	DB       atomic.Pointer[sql.DB]
	IsActive atomic.Bool
}

type MysqlConnectionOptions struct {
	User     string
	Password string
	Database string
	Host     string
	Port     int
}

var (
	ErrNoHealthyConn            = errors.New("no healthy mysql connection available")
	ErrUnsupportedBalancingMode = errors.New("unsupported balancing mode")
)

// NewMysqlConnectionPool builds a connection pool for the configured cluster and starts
// a background goroutine that periodically re-checks node health.
//
// If no node is reachable at startup it returns a non-nil pool together with
// ErrNoHealthyConn: this is a recoverable condition, not a fatal one. The background
// pinger keeps running and will mark nodes active again once they come back, so callers
// that want degraded/emergency operation can keep using the returned pool. Because the
// pinger owns a ticker and a goroutine, the caller must call Close() on the returned pool
// even when it received ErrNoHealthyConn.
func NewMysqlConnectionPool(config *ClusterConfig) (*MysqlConnectionPool, error) {
	connPool := &MysqlConnectionPool{}
	connPool.Connections = make([]*MysqlConnection, 0, len(config.Connections))
	connPool.Mutex = &sync.Mutex{}
	connPool.Config = config

	for _, connection := range config.Connections {
		conn := &MysqlConnection{}
		conn.Name = connection.Host
		conn.Config = &MysqlConnectionOptions{
			User:     config.User,
			Password: config.Password,
			Database: config.Database,
			Host:     connection.Host,
			Port:     connection.Port,
		}
		conn.Dsn = GetMysqlDsn(conn.Config, false)

		db, err := connect(conn.Dsn)
		conn.DB.Store(db)

		var isReachable bool
		switch {
		case err != nil:
			log.Error("failed to open database handle", zap.String("name", conn.Name), zap.Error(err))
		case !connPool.checkIsReachable(conn):
			log.Warn("database node is not reachable at startup", zap.String("name", conn.Name))
		default:
			isReachable = true
		}
		conn.IsActive.Store(isReachable)

		connPool.Connections = append(connPool.Connections, conn)
	}

	// Snapshot construction-time reachability to decide the returned error. IsActive is
	// atomic, so this remains correct regardless of the pinger; we read it here so the
	// result reflects startup state rather than a value the first tick may have changed.
	hasHealthyConn := false
	for _, connection := range connPool.Connections {
		if connection.IsActive.Load() {
			hasHealthyConn = true
			break
		}
	}

	connPool.PingsTicker = time.NewTicker(10 * time.Second)
	connPool.PingsDone = make(chan bool)
	go connPool.startDatabasePinging()

	if hasHealthyConn {
		return connPool, nil
	}

	return connPool, ErrNoHealthyConn
}

func GetMysqlDsn(connection *MysqlConnectionOptions, multiStatements bool) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&multiStatements=%t&timeout=%s&readTimeout=%s&writeTimeout=%s",
		connection.User,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.Database,
		multiStatements,
		dialTimeout,
		ioTimeout,
		ioTimeout,
	)
}

func connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %v", err)
	}

	// Recycle connections so a stale/broken connection to a node that went away is not
	// reused indefinitely; the next use then triggers a fresh, timeout-bounded dial.
	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetConnMaxIdleTime(maxConnIdleTime)

	return db, nil
}

func (m *MysqlConnectionPool) checkIsReachable(conn *MysqlConnection) bool {

	db := conn.DB.Load()
	if db == nil {
		return false
	}

	// Bound the health check so the ping loop can never block indefinitely on an
	// unresponsive node, even if the DSN-level timeouts were somehow not applied.
	ctx, cancel := context.WithTimeout(context.Background(), healthCheckTimeout)
	defer cancel()

	// if it's not a cluster we can just ping the database
	if !*m.Config.IsGaleraCluster {
		err := db.PingContext(ctx)
		log.WarnIfError("failed to ping database", err, zap.String("name", conn.Name))
		return err == nil
	} else {
		// otherwise we need to check the global wsrep_ready state
		var variableName string
		var wsrepStatus string

		err := db.QueryRowContext(ctx, "SHOW GLOBAL STATUS LIKE 'wsrep_ready'").Scan(&variableName, &wsrepStatus)
		if err != nil {
			log.Warn("failed to check database connection", zap.Error(err), zap.String("name", conn.Name))
			return false
		}

		return wsrepStatus == "ON"
	}
}

func (m *MysqlConnectionPool) startDatabasePinging() {
	for {
		select {
		case <-m.PingsDone:
			log.Log(log.INFO, "stop pinging database connections")
			return
		case <-m.PingsTicker.C:

			// Refresh all connections concurrently so a single slow/dead node cannot
			// stall detection for the others (health checks are individually bounded).
			var wg sync.WaitGroup
			var numHealthy, numUnhealthy int64

			for _, conn := range m.Connections {
				wg.Add(1)
				go func(conn *MysqlConnection) {
					defer wg.Done()
					if m.refreshConnection(conn) {
						atomic.AddInt64(&numHealthy, 1)
					} else {
						atomic.AddInt64(&numUnhealthy, 1)
					}
				}(conn)
			}
			wg.Wait()

			recordConnStats(int(numHealthy), int(numUnhealthy))
		}
	}
}

// refreshConnection re-checks a single connection's health and updates its IsActive
// flag accordingly. It returns whether the connection is currently reachable.
//
// The underlying *sql.DB is created once (at pool construction, or lazily here if that
// ever failed) and reused across outages: database/sql owns its own connection pool and
// transparently re-dials when a node recovers, so the handle is never rebuilt on a mere
// health-check failure -- doing so would churn connection-pool goroutines every tick for
// the duration of an outage.
func (m *MysqlConnectionPool) refreshConnection(conn *MysqlConnection) bool {
	// Ensure a handle exists. connect() only fails on a malformed DSN, so this is a
	// one-off cost that effectively never recurs during normal operation.
	if conn.DB.Load() == nil {
		db, err := connect(conn.Dsn)
		if log.WarnIfError("failed to open database handle", err, zap.String("name", conn.Name)) {
			conn.IsActive.Store(false)
			return false
		}
		conn.DB.Store(db)
	}

	wasActive := conn.IsActive.Load()
	isReachable := m.checkIsReachable(conn)

	if isReachable && !wasActive { // conn was previously not reachable but now is again
		log.Info("successfully reconnected to database", zap.String("name", conn.Name))
	}

	conn.IsActive.Store(isReachable)
	return isReachable
}

// MustGetConnection returns an active connection of panics if none of the connections from the pool is healthy
func (m *MysqlConnectionPool) MustGetConnection() *sql.DB {

	conn, err := m.GetConnection()

	if err != nil {
		panic(err)
	}

	return conn
}

// GetConnection returns an active connection or ErrNoHealthyConn if none of the connections from the pool is healthy
func (m *MysqlConnectionPool) GetConnection() (*sql.DB, error) {

	active, err := m.getActive()

	if err != nil {
		return nil, err
	}

	return active.DB.Load(), err
}

func (m *MysqlConnectionPool) GetActiveConfig() (*MysqlConnectionOptions, error) {

	active, err := m.getActive()

	if err != nil {
		return nil, err
	}

	return active.Config, err
}

func (m *MysqlConnectionPool) getActive() (*MysqlConnection, error) {

	// IsActive is read atomically, so no pool-level lock is needed here; this keeps
	// connection selection off the hot path's lock while remaining race-free.
	if len(m.Connections) == 0 {
		incNoHealthyConnError()
		return nil, ErrNoHealthyConn
	}

	switch m.Config.BalancingMode {
	case Random:
		randConn := rand.Intn(len(m.Connections))

		// check if this random connection is active
		if m.Connections[randConn].IsActive.Load() {
			return m.Connections[randConn], nil
		}

		// if the random connection is not active, we fall through here and get the first active one
		fallthrough

	case Ordered:

		// get the first active connection and return it
		for _, db := range m.Connections {
			if db.IsActive.Load() {
				return db, nil
			}
		}

		// could not find any healthy connection, report and return ErrNoHealthyConn
		incNoHealthyConnError()
		return nil, ErrNoHealthyConn

	default:
		return nil, errors.WithMessage(ErrUnsupportedBalancingMode, string(m.Config.BalancingMode))
	}
}

// MustBeginTx starts a database transaction or panics if an error occurs
func (m *MysqlConnectionPool) MustBeginTx() *sql.Tx {
	db := m.MustGetConnection()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	return tx
}

// BeginTx starts a database transaction
func (m *MysqlConnectionPool) BeginTx() (*sql.Tx, error) {
	db, err := m.GetConnection()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// IsTx checks whether the given executor is a transaction
func IsTx(executor boil.ContextExecutor) bool {
	_, ok := executor.(*sql.Tx)
	return ok
}

// MustRollbackTx rolls back the given transaction or panics if an error occurs
func MustRollbackTx(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		panic(fmt.Errorf("failed to rollback transaction: %e", err))
	}
}

// MustCommit commits the given transaction or panics if an error occurs
func MustCommit(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		panic(fmt.Errorf("failed to commit transaction: %e", err))
	}
}

// Close closes all database connections from the pool
func (m *MysqlConnectionPool) Close() {

	m.PingsTicker.Stop()
	close(m.PingsDone)

	for _, conn := range m.Connections {
		// Close every handle we hold, not only the active ones: a node that is down at
		// shutdown still has an open *sql.DB whose pool goroutines would otherwise leak.
		if db := conn.DB.Load(); db != nil {
			err := db.Close()
			log.CriticalIfError("failed to close database connection", err, zap.String("connection_name", conn.Name))
		}
	}

	log.Info("closed all database connections")
}
