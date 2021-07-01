package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/eosnationftw/eosn-base-api/log"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"sync"
	"time"
)

type MysqlConnectionPool struct {
	Connections []*MysqlConnection
	Mutex       *sync.Mutex
	PingsTicker *time.Ticker
	PingsDone   chan bool
}

type MysqlConnection struct {
	Name     string
	Dsn      string
	DB       *sql.DB
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
	ErrNoHealthyConn = errors.New("no healthy mysql connection available")
)

func NewMysqlConnectionPool(connections []MysqlConnectionOptions) (connPool *MysqlConnectionPool, err error) {
	connPool = &MysqlConnectionPool{}
	connPool.Connections = make([]*MysqlConnection, 0, len(connections))
	connPool.Mutex = &sync.Mutex{}

	for _, connection := range connections {
		conn := &MysqlConnection{}
		conn.Name = connection.Host
		conn.Dsn = GetMysqlDsn(connection, false)

		db, err := connect(conn.Dsn)
		conn.DB = db
		conn.IsActive = true

		if log.CriticalIfError("could not connect to database", err) {
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

func GetMysqlDsn(connection MysqlConnectionOptions, multiStatements bool) string {
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

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping DB: %v", err)
	}

	return db, nil
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
						err := conn.DB.Ping()
						if err != nil {
							log.Warn("failed to ping database", zap.Error(err), zap.String("name", conn.Name))
							isReachable = false
						} else if !conn.IsActive { // conn was previously not reachable but can be pinged again
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
							err = db.Ping()
							if !log.WarnIfError("failed to ping database", err, zap.String("name", conn.Name)) {
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

func (m *MysqlConnectionPool) MustGetConnection() *sql.DB {

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for _, db := range m.Connections {
		if db.IsActive {
			return db.DB
		}
	}
	incNoHealthyConnError()
	panic(ErrNoHealthyConn)
}

func (m *MysqlConnectionPool) GetConnection() (*sql.DB, error) {

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for _, db := range m.Connections {
		if db.IsActive {
			return db.DB, nil
		}
	}
	incNoHealthyConnError()
	return nil, ErrNoHealthyConn
}

func (m *MysqlConnectionPool) BeginTx() *sql.Tx {
	db := m.MustGetConnection()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	return tx
}

func MustRollbackTx(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		panic(fmt.Errorf("failed to rollback transaction: %e", err))
	}
}

func MustCommit(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		panic(fmt.Errorf("failed to commit transaction: %e", err))
	}
}

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
