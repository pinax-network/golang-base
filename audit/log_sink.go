package audit

import (
	"fmt"
	"github.com/pinax-network/golang-base/log"
	"go.uber.org/zap"
	"time"
)

type LogSink struct {
}

func NewLogSink() *LogSink {
	return &LogSink{}
}

func (f *LogSink) CreateResource(userId int, resourceId int, resource interface{}, time time.Time) {
	log.Info(
		"[created resource]",
		zap.Int("user_id", userId),
		zap.Int("resource_id", resourceId),
		zap.String("resource_type", fmt.Sprintf("%T", resource)),
		zap.Any("resource_data", resource),
		zap.Time("time", time),
	)
}

func (f *LogSink) UpdateResource(userId int, resourceId int, newData interface{}, prevData interface{}, time time.Time) {
	log.Info(
		"[updated resource]",
		zap.Int("user_id", userId),
		zap.Int("resource_id", resourceId),
		zap.String("resource_type", fmt.Sprintf("%T", newData)),
		zap.Any("new_data", newData),
		zap.Any("old_data", prevData),
		zap.Time("time", time),
	)
}

func (f *LogSink) DeleteResource(userId int, resourceId int, resource interface{}, time time.Time) {
	log.Info(
		"[deleted resource]",
		zap.Int("user_id", userId),
		zap.Int("resource_id", resourceId),
		zap.String("resource_type", fmt.Sprintf("%T", resource)),
		zap.Any("resource_data", resource),
		zap.Time("time", time),
	)
}
