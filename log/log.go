package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	PANIC
	FATAL
)

var ZapLogger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func InitializeLogger(logDebug bool) error {

	var consoleEncoder zapcore.Encoder
	var minLevel zapcore.Level

	if logDebug {
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		minLevel = zapcore.DebugLevel
	} else {
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder = zapcore.NewConsoleEncoder(cfg)
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

func InitializeFileLogger(logDebug bool, logLevel zapcore.Level, file *os.File) error {

	var consoleEncoder zapcore.Encoder
	var minLevel zapcore.Level

	if logDebug {
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		minLevel = zapcore.DebugLevel
	} else {
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder = zapcore.NewConsoleEncoder(cfg)
		minLevel = logLevel
	}

	core := zapcore.NewCore(consoleEncoder, file, minLevel)

	logger := zap.New(core)

	ZapLogger = logger
	SugaredLogger = logger.Sugar()

	return nil
}

func Log(logLevel LogLevel, message string, additionalFields ...zap.Field) {

	if ZapLogger == nil {
		log.Println("zap logger isn't initialized yet!")
		log.Println(message)

		for _, f := range additionalFields {
			log.Println(fmt.Sprintf("'%s': %+v", f.Key, f))
		}
		if logLevel == FATAL {
			os.Exit(1)
		}
		if logLevel == PANIC {
			panic(message)
		}
		return
	}

	switch logLevel {
	case DEBUG:
		ZapLogger.Debug(message, additionalFields...)
		break
	case INFO:
		ZapLogger.Info(message, additionalFields...)
		break
	case WARNING:
		incWarnCounter()
		ZapLogger.Warn(message, additionalFields...)
		break
	case ERROR:
		incErrorCounter()
		ZapLogger.Error(message, additionalFields...)
		break
	case PANIC:
		incPanicCounter()
		ZapLogger.Panic(message, additionalFields...)
		break
	case FATAL:
		incFatalCounter()
		ZapLogger.Fatal(message, additionalFields...)
		break
	}
}

func Debug(message string, additionalFields ...zap.Field) {
	Log(DEBUG, message, additionalFields...)
}

func Debugf(template string, args ...interface{}) {
	SugaredLogger.Debugf(template, args)
}

func Info(message string, additionalFields ...zap.Field) {
	Log(INFO, message, additionalFields...)
}

func Infof(template string, args ...interface{}) {
	SugaredLogger.Infof(template, args)
}

func Warn(message string, additionalFields ...zap.Field) {
	Log(WARNING, message, additionalFields...)
}

func Warnf(template string, args ...interface{}) {
	SugaredLogger.Warnf(template, args)
}

func Error(message string, additionalFields ...zap.Field) {
	Log(ERROR, message, additionalFields...)
}

func Errorf(template string, args ...interface{}) {
	SugaredLogger.Errorf(template, args)
}

func Panic(message string, additionalFields ...zap.Field) {
	Log(PANIC, message, additionalFields...)
}

func Panicf(template string, args ...interface{}) {
	SugaredLogger.Panicf(template, args)
}

func Fatal(message string, additionalFields ...zap.Field) {
	Log(FATAL, message, additionalFields...)
}

func Fatalf(template string, args ...interface{}) {
	SugaredLogger.Fatalf(template, args)
}

func LogIfError(logLevel LogLevel, message string, err error, additionalFields ...zap.Field) bool {
	if err != nil {
		fields := append([]zap.Field{zap.Error(err)}, additionalFields...)
		Log(logLevel, message, fields...)
		return true
	}
	return false
}

// FatalIfError logs the given messages plus additional fields and exits the application afterwards.
func FatalIfError(message string, err error, additionalFields ...zap.Field) {
	LogIfError(FATAL, message, err, additionalFields...)
}

// PanicIfError logs the given messages plus additional fields and exits the application afterwards.
func PanicIfError(message string, err error, additionalFields ...zap.Field) {
	if LogIfError(PANIC, message, err, additionalFields...) {
		panic(err)
	}
}

// CriticalIfError logs the given message and fields as critical if the given error is not null. Returns true if an error occurred or false if not.
func CriticalIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(ERROR, message, err, additionalFields...)
}

// WarnIfError logs the given message and fields as warning if the given error is not null. Returns true if an error occurred or false if not.
func WarnIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(WARNING, message, err, additionalFields...)
}

// InfoIfError logs the given message and fields as info if the given error is not null. Returns true if an error occurred or false if not.
func InfoIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(INFO, message, err, additionalFields...)
}

// DebugIfError logs the given message and fields as debug if the given error is not null. Returns true if an error occurred or false if not.
func DebugIfError(message string, err error, additionalFields ...zap.Field) bool {
	return LogIfError(DEBUG, message, err, additionalFields...)
}
