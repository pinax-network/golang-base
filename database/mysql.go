package database

import (
	"database/sql"
	"fmt"
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/friendsofgo/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
	"math/rand"
	"sync"
	"time"
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
	ErrUnsupportedBalancingMode = errors.New("unsupported balancing mode given")
)

func NewMysqlConnectionPool(config *ClusterConfig) (connPool *MysqlConnectionPool, err error) {
	connPool = &MysqlConnectionPool{}
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

	for _, connection := range connPool.Connections {
		if connection.IsActive {
			connPool.startDatabasePinging()
			err = nil
			return
		}
	}

	err = ErrNoHealthyConn
	return
}

func GetMysqlDsn(connection *MysqlConnectionOptions, multiStatements bool) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=%t",
		connection.User,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.Database,
		multiStatements,
	)
}

func connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %v", err)
	}

	return db, nil
}

func (m *MysqlConnectionPool) checkIsReachable(conn *MysqlConnection) bool {

	// if it's not a cluster we can just ping the database
	if !*m.Config.IsGaleraCluster {
		err := conn.DB.Ping()
		log.WarnIfError("failed to ping database", err, zap.String("name", conn.Name))
		return err == nil
	} else {
		// otherwise we need to check the global wsrep_ready state
		var variableName string
		var wsrepStatus string

		err := conn.DB.QueryRow("SHOW GLOBAL STATUS LIKE 'wsrep_ready'").Scan(&variableName, &wsrepStatus)
		if err != nil {
			log.Warn("failed to check database connection", zap.Error(err), zap.String("name", conn.Name))
			return false
		}

		return wsrepStatus == "ON"
	}
}

func (m *MysqlConnectionPool) startDatabasePinging() {

	m.PingsTicker = time.NewTicker(10 * time.Second)
	m.PingsDone = make(chan bool)

	go func() {
		for {
			select {
			case <-m.PingsDone:
				log.Log(log.INFO, "stop pinging database connections")
				return
			case <-m.PingsTicker.C:

				numHealthy := 0
				numUnhealthy := 0

				for _, conn := range m.Connections {
					isReachable := true

					if conn.DB != nil {
						if !m.checkIsReachable(conn) {
							isReachable = false
						} else if !conn.IsActive { // conn was previously not reachable but now is again
							log.Info("successfully reconnected to database", zap.String("name", conn.Name))
							m.Mutex.Lock()
							conn.IsActive = isReachable
							m.Mutex.Unlock()
						}
					}
					// try to reconnect to database
					if conn.DB == nil || !isReachable {
						db, err := connect(conn.Dsn)
						if log.WarnIfError("failed to (re-)connect to database", err, zap.String("name", conn.Name)) {
							isReachable = false
						} else {
							if !m.checkIsReachable(conn) {
								isReachable = false
							}
						}

						m.Mutex.Lock()
						conn.DB = db
						conn.IsActive = isReachable
						m.Mutex.Unlock()
					}

					if isReachable {
						numHealthy++
					} else {
						numUnhealthy++
					}
				}

				recordConnStats(numHealthy, numUnhealthy)
			}
		}
	}()
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
		rand.Seed(time.Now().UnixNano())
		randConn := rand.Intn(len(m.Connections))

		// check if random connection is active
		if m.Connections[randConn].IsActive {
			return m.Connections[randConn], nil
		}

		// if the random connection is not active we fall through here and get the first active one
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
