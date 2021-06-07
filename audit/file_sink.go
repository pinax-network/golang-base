package audit

import (
	"encoding/json"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"time"
)

type FileSink struct {
	logger *zap.Logger
}

func NewFileSink(logDir string) (*FileSink, error) {
	res := &FileSink{}

	logFile := path.Join(logDir, "audit-%Y-%m-%d.log")
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		return nil, err
	}

	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	if err := json.Unmarshal(data, &encCfg); err != nil {
		return nil, err
	}

	// add the encoder config and rotator to create a new zap logger
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), w, zap.InfoLevel)
	res.logger = zap.New(core)

	return res, nil
}

func (f *FileSink) CreateResource(userId int, resourceId int, resource interface{}, time time.Time) {
	f.logger.Info(
		"[created resource]",
		zap.Int("user_id", userId),
		zap.Int("resource_id", resourceId),
		zap.String("resource_type", fmt.Sprintf("%T", resource)),
		zap.Any("resource_data", resource),
		zap.Time("time", time),
	)
}

func (f *FileSink) UpdateResource(userId int, resourceId int, newData interface{}, prevData interface{}, time time.Time) {
	f.logger.Info(
		"[updated resource]",
		zap.Int("user_id", userId),
		zap.Int("resource_id", resourceId),
		zap.String("resource_type", fmt.Sprintf("%T", newData)),
		zap.Any("new_data", newData),
		zap.Any("old_data", prevData),
		zap.Time("time", time),
	)
}

func (f *FileSink) DeleteResource(userId int, resourceId int, time time.Time) {
	f.logger.Info(
		"[deleted resource]",
		zap.Int("user_id", userId),
		zap.Int("resource_id", resourceId),
		zap.Time("time", time),
	)
}
