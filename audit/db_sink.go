package audit

import (
	"encoding/json"
	"fmt"
	"github.com/eosnationftw/eosn-base-api/database"
	"time"
)

type DBSink struct {
	dbPool *database.MysqlConnectionPool
}

func NewDBSink(dbPool *database.MysqlConnectionPool) *DBSink {
	return &DBSink{
		dbPool: dbPool,
	}
}

func (d *DBSink) CreateResource(userId int, resourceId int, resource interface{}, time time.Time) {

	conn := d.dbPool.MustGetConnection()
	resourceJson, err := json.Marshal(resource)

	if err != nil {
		panic(fmt.Sprintf("failed to json marshal resource for audit logging: %e", err))
	}

	stmt, err := conn.Prepare("INSERT INTO audit_log (user_id, resource_id, action_type, resource_type, resource, time) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(fmt.Sprintf("failed to prepare audit log statement: %e", err))
	}

	_, err = stmt.Exec(userId, resourceId, createAction, fmt.Sprintf("%T", resource), resourceJson, time)
	if err != nil {
		panic(fmt.Sprintf("failed to insert audit log: %e", err))
	}
}

func (d *DBSink) UpdateResource(userId int, resourceId int, newData interface{}, prevData interface{}, time time.Time) {

	conn := d.dbPool.MustGetConnection()

	newDataJson, err := json.Marshal(newData)
	if err != nil {
		panic(fmt.Sprintf("failed to json marshal resource for audit logging: %e", err))
	}

	prevDataJson, err := json.Marshal(prevData)
	if err != nil {
		panic(fmt.Sprintf("failed to json marshal resource for audit logging: %e", err))
	}

	stmt, err := conn.Prepare("INSERT INTO audit_log (user_id, resource_id, action_type, resource_type, resource, resource_prev, time) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(fmt.Sprintf("failed to prepare audit log statement: %e", err))
	}

	_, err = stmt.Exec(userId, resourceId, updateAction, fmt.Sprintf("%T", newData), newDataJson, prevDataJson, time)
	if err != nil {
		panic(fmt.Sprintf("failed to insert audit log: %e", err))
	}
}

func (d *DBSink) DeleteResource(userId int, resourceId int, time time.Time) {

	conn := d.dbPool.MustGetConnection()

	stmt, err := conn.Prepare("INSERT INTO audit_log (user_id, resource_id, action_type, time) VALUES (?,?,?,?)")
	if err != nil {
		panic(fmt.Sprintf("failed to prepare audit log statement: %e", err))
	}

	_, err = stmt.Exec(userId, resourceId, deleteAction, time)
	if err != nil {
		panic(fmt.Sprintf("failed to insert audit log: %e", err))
	}
}
