package log

import (
	"fmt"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LogLevelType = int

var LogLevel = struct {
	DEBUG   LogLevelType
	INFO    LogLevelType
	WARNING LogLevelType
	ERROR   LogLevelType
	FATAL   LogLevelType
}{
	DEBUG:   0,
	INFO:    1,
	WARNING: 2,
	ERROR:   3,
	FATAL:   4,
}

var ZapLogger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func InitializeLogger(logDebug bool) error {

	var consoleEncoder zapcore.Encoder
	var minLevel zapcore.Level

	if logDebug {
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		minLevel = zapcore.DebugLevel
	} else {
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
		minLevel = zapcore.InfoLevel
	}

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && lvl >= minLevel
	})

	consoleOut := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleOut, lowPriority),
	)

	logger := zap.New(core)

	ZapLogger = logger
	SugaredLogger = logger.Sugar()

	return nil
}

func Log(logLevel LogLevelType, message string, additionalFields ...zap.Field) {

	if ZapLogger == nil {
		log.Error("zap logger isn't initialized yet!")
		log.Error(message)

		for _, f := range additionalFields {
			log.Error(fmt.Sprintf("'%s': %s", f.Key, f.String))
		}
		if logLevel == LogLevel.FATAL {
			os.Exit(1)
		}
		return
	}

	switch logLevel {
	case LogLevel.DEBUG:
		ZapLogger.Debug(message, additionalFields...)
		break
	case LogLevel.INFO:
		ZapLogger.Info(message, additionalFields...)
		break
	case LogLevel.WARNING:
		incWarnCounter()
		ZapLogger.Warn(message, additionalFields...)
		break
	case LogLevel.ERROR:
		incErrorCounter()
		ZapLogger.Error(message, additionalFields...)
		break
	case LogLevel.FATAL:
		incFatalCounter()
		ZapLogger.Fatal(message, additionalFields...)
		break
	}
}

func Debug(message string, additionalFields ...zap.Field) {
	Log(LogLevel.DEBUG, message, additionalFields...)
}

func Info(message string, additionalFields ...zap.Field) {
	Log(LogLevel.INFO, message, additionalFields...)
}

func Warn(message string, additionalFields ...zap.Field) {
	Log(LogLevel.WARNING, message, additionalFields...)
}

func Error(message string, additionalFields ...zap.Field) {
	Log(LogLevel.ERROR, message, additionalFields...)
}

func Fatal(message string, additionalFields ...zap.Field) {
	Log(LogLevel.FATAL, message, additionalFields...)
}

func LogIfError(logLevel LogLevelType, message string, err error, additionalFields ...zap.Field) bool {
	if err != nil {
		fields := append([]zap.Field{zap.Error(err)}, additionalFields...)
		Log(logLevel, message, fields...)
		return true
	}
	return false
}

// FatalIfError logs the given messages plus additional fields and exits the application afterwards.
func FatalIfError(message string, err error, additionalFields ...zap.Field) {
	LogIfError(LogLevel.FATAL, message, err, additionalFields...)
}

// PanicIfError logs the given messages plus additional fields and exits the application afterwards.
func PanicIfError(message string, err error, additionalFields ...zap.Field) {
	if LogIfError(LogLevel.ERROR, message, err, additionalFields...) {
		panic(err)
	}
}

// CriticalIfError logs the given message and fields as critical if the given error is not null. Returns true if an error occurred or false if not.
func CriticalIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(LogLevel.ERROR, message, err, additionalFields...)
}

// WarnIfError logs the given message and fields as warning if the given error is not null. Returns true if an error occurred or false if not.
func WarnIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(LogLevel.WARNING, message, err, additionalFields...)
}

// InfoIfError logs the given message and fields as info if the given error is not null. Returns true if an error occurred or false if not.
func InfoIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(LogLevel.INFO, message, err, additionalFields...)
}

// DebugIfError logs the given message and fields as debug if the given error is not null. Returns true if an error occurred or false if not.
func DebugIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(LogLevel.DEBUG, message, err, additionalFields...)
}
