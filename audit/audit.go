package audit

import (
	"github.com/eosnationftw/eosn-base-api/log"
	"time"
)

var auditLogger *AuditLogger

var (
	createAction = "create"
	updateAction = "update"
	deleteAction = "delete"
)

type AuditLogger struct {
	sinks []Sink
}

type Sink interface {
	CreateResource(userId int, resourceId int, resource interface{}, time time.Time)
	UpdateResource(userId int, resourceId int, newData interface{}, prevData interface{}, time time.Time)
	DeleteResource(userId int, resourceId int, resource interface{}, time time.Time)
}

func InitializeAuditLog(sinks ...Sink) {

	if len(sinks) == 0 {
		log.Warn("no sink is set for the audit log!")
	}

	auditLogger = &AuditLogger{sinks: sinks}
}

func LogCreateResource(userId int, resourceId int, resource interface{}, time time.Time) {
	incCreateResourceCounter()
	for _, s := range auditLogger.sinks {
		s.CreateResource(userId, resourceId, resource, time)
	}
}

func LogUpdateResource(userId int, resourceId int, newData interface{}, prevData interface{}, time time.Time) {
	incUpdateResourceCounter()
	for _, s := range auditLogger.sinks {
		s.UpdateResource(userId, resourceId, newData, prevData, time)
	}
}

func LogDeleteResource(userId int, resourceId int, resource interface{}, time time.Time) {
	incDeleteResourceCounter()
	for _, s := range auditLogger.sinks {
		s.DeleteResource(userId, resourceId, resource, time)
	}
}
