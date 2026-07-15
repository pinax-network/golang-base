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
	Name     string
	Dsn      string
	DB       *sql.DB
	Config   *MysqlConnectionOptions
	IsActive bool
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
		conn.DB = db
		conn.IsActive = true

		if err != nil || !connPool.checkIsReachable(conn) {
			log.Error("could not connect to database", zap.Any("conn", conn))
			conn.IsActive = false
		}

		connPool.Connections = append(connPool.Connections, conn)
	}

	connPool.PingsTicker = time.NewTicker(10 * time.Second)
	connPool.PingsDone = make(chan bool)
	go connPool.startDatabasePinging()
	for _, connection := range connPool.Connections {
		if connection.IsActive {
			return connPool, nil
		}
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

	// Bound the health check so the ping loop can never block indefinitely on an
	// unresponsive node, even if the DSN-level timeouts were somehow not applied.
	ctx, cancel := context.WithTimeout(context.Background(), healthCheckTimeout)
	defer cancel()

	// if it's not a cluster we can just ping the database
	if !*m.Config.IsGaleraCluster {
		err := conn.DB.PingContext(ctx)
		log.WarnIfError("failed to ping database", err, zap.String("name", conn.Name))
		return err == nil
	} else {
		// otherwise we need to check the global wsrep_ready state
		var variableName string
		var wsrepStatus string

		err := conn.DB.QueryRowContext(ctx, "SHOW GLOBAL STATUS LIKE 'wsrep_ready'").Scan(&variableName, &wsrepStatus)
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

// refreshConnection checks a single connection's health, rebuilding the underlying
// handle if it is not reachable, and updates its IsActive flag under the pool mutex.
// It returns whether the connection is currently reachable.
func (m *MysqlConnectionPool) refreshConnection(conn *MysqlConnection) bool {
	// Fast path: an existing, healthy connection needs no churn.
	if conn.DB != nil && m.checkIsReachable(conn) {
		if !conn.IsActive { // conn was previously not reachable but now is again
			log.Info("successfully reconnected to database", zap.String("name", conn.Name))
			m.Mutex.Lock()
			conn.IsActive = true
			m.Mutex.Unlock()
		}
		return true
	}

	// The connection just failed its health check (or has no handle yet). Mark it
	// inactive right away so getActive() stops routing to it while we rebuild and
	// re-validate the handle below.
	m.Mutex.Lock()
	conn.IsActive = false
	m.Mutex.Unlock()

	// Rebuild the handle. connect() only fails on a malformed DSN; if it ever does,
	// keep the existing handle in place rather than replacing it with a nil one.
	newDB, err := connect(conn.Dsn)
	if log.WarnIfError("failed to (re-)connect to database", err, zap.String("name", conn.Name)) {
		return false
	}

	// Swap in the fresh handle, validating the new one rather than the old.
	m.Mutex.Lock()
	oldDB := conn.DB
	conn.DB = newDB
	m.Mutex.Unlock()

	// Close the old handle to avoid leaking it across reconnects. Close() waits for
	// in-flight queries, so do it asynchronously to avoid stalling the ping cycle; the
	// handle pointed at a node we already deemed unreachable.
	if oldDB != nil {
		go func() {
			log.WarnIfError("failed to close stale database connection", oldDB.Close(), zap.String("name", conn.Name))
		}()
	}

	isReachable := m.checkIsReachable(conn)

	m.Mutex.Lock()
	conn.IsActive = isReachable
	m.Mutex.Unlock()

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

	return active.DB, err
}

func (m *MysqlConnectionPool) GetActiveConfig() (*MysqlConnectionOptions, error) {

	active, err := m.getActive()

	if err != nil {
		return nil, err
	}

	return active.Config, err
}

func (m *MysqlConnectionPool) getActive() (*MysqlConnection, error) {

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	switch m.Config.BalancingMode {
	case Random:
		randConn := rand.Intn(len(m.Connections))

		// check if this random connection is active
		if m.Connections[randConn].IsActive {
			return m.Connections[randConn], nil
		}

		// if the random connection is not active, we fall through here and get the first active one
		fallthrough

	case Ordered:

		// get the first active connection and return it
		for _, db := range m.Connections {
			if db.IsActive {
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
		if conn.IsActive && conn.DB != nil {
			err := conn.DB.Close()
			log.CriticalIfError("failed to close database connection", err, zap.String("connection_name", conn.Name))
		}
	}

	log.Info("closed all database connections")
}
