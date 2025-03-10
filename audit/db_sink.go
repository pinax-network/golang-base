package audit

import (
	"encoding/json"
	"fmt"
	"github.com/pinax-network/golang-base/database"
	"github.com/pinax-network/golang-base/log"
	"go.uber.org/zap"
	"time"
)

type MysqlSink struct {
	dbPool *database.MysqlConnectionPool
}

// NewMysqlSink initializes a new MySQL sink for audit logs. Note that an audit_log table needs to exist, see the
// mysql_audit_log.up.sql migration file on how to set up the necessary table.
func NewMysqlSink(dbPool *database.MysqlConnectionPool) *MysqlSink {
	return &MysqlSink{
		dbPool: dbPool,
	}
}

func (d *MysqlSink) CreateResource(userId int, resourceId int, resource interface{}, time time.Time) {

	conn, err := d.dbPool.GetConnection()
	if err != nil {
		log.Error("failed to get database connection", zap.Error(err))
		return
	}

	resourceJson, err := json.Marshal(resource)
	if err != nil {
		log.Error("failed to json marshal resource for audit logging", zap.Error(err))
		return
	}

	stmt, err := conn.Prepare("INSERT INTO `audit_log` (`user_id`, `resource_id`, `action_type`, `resource_type`, `resource`, `time`) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Error("failed to prepare audit log statement", zap.Error(err))
		return
	}

	_, err = stmt.Exec(userId, resourceId, createAction, fmt.Sprintf("%T", resource), resourceJson, time)
	if err != nil {
		log.Error("failed to insert audit log", zap.Error(err))
		return
	}
}

func (d *MysqlSink) UpdateResource(userId int, resourceId int, newData interface{}, prevData interface{}, time time.Time) {

	conn, err := d.dbPool.GetConnection()
	if err != nil {
		log.Error("failed to get database connection", zap.Error(err))
		return
	}

	newDataJson, err := json.Marshal(newData)
	if err != nil {
		log.Error("failed to json marshal resource for audit logging", zap.Error(err))
		return
	}

	prevDataJson, err := json.Marshal(prevData)
	if err != nil {
		log.Error("failed to json marshal resource for audit logging", zap.Error(err))
		return
	}

	stmt, err := conn.Prepare("INSERT INTO `audit_log` (`user_id`, `resource_id`, `action_type`, `resource_type`, `resource`, `resource_prev`, `time`) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Error("failed to prepare audit log statement", zap.Error(err))
		return
	}

	_, err = stmt.Exec(userId, resourceId, updateAction, fmt.Sprintf("%T", newData), newDataJson, prevDataJson, time)
	if err != nil {
		log.Error("failed to insert audit log", zap.Error(err))
		return
	}
}

func (d *MysqlSink) DeleteResource(userId int, resourceId int, resource interface{}, time time.Time) {

	conn, err := d.dbPool.GetConnection()
	if err != nil {
		log.Error("failed to get database connection", zap.Error(err))
		return
	}

	resourceJson, err := json.Marshal(resource)
	if err != nil {
		log.Error("failed to json marshal resource for audit logging", zap.Error(err))
		return
	}

	stmt, err := conn.Prepare("INSERT INTO `audit_log` (`user_id`, `resource_id`, `action_type`, `resource_type`, `resource`, `time`) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Error("failed to prepare audit log statement", zap.Error(err))
		return
	}

	_, err = stmt.Exec(userId, resourceId, deleteAction, fmt.Sprintf("%T", resource), resourceJson, time)
	if err != nil {
		log.Error("failed to insert audit log", zap.Error(err))
		return
	}
}
