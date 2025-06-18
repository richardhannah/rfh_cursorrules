package logger

import (
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// Init initializes the logger with JSON format
func Init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.CallerKey = "file"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var err error
	log, err = config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}

// getCallerInfo returns the function information
func getCallerInfo() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	return fn.Name()
}

// Info logs an info level message with optional fields
func Info(msg string, fields ...zap.Field) {
	if log == nil {
		Init()
	}
	function := getCallerInfo()
	fields = append(fields, zap.String("function", function))
	log.Info(msg, fields...)
}

// Error logs an error level message with optional fields
func Error(msg string, err error, fields ...zap.Field) {
	if log == nil {
		Init()
	}
	function := getCallerInfo()
	fields = append(fields,
		zap.String("function", function),
		zap.Error(err),
	)
	log.Error(msg, fields...)
}

// Warn logs a warning level message with optional fields
func Warn(msg string, fields ...zap.Field) {
	if log == nil {
		Init()
	}
	function := getCallerInfo()
	fields = append(fields, zap.String("function", function))
	log.Warn(msg, fields...)
}

// Debug logs a debug level message with optional fields
func Debug(msg string, fields ...zap.Field) {
	if log == nil {
		Init()
	}
	function := getCallerInfo()
	fields = append(fields, zap.String("function", function))
	log.Debug(msg, fields...)
}

// Fatal logs a fatal level message and exits
func Fatal(msg string, err error, fields ...zap.Field) {
	if log == nil {
		Init()
	}
	function := getCallerInfo()
	fields = append(fields,
		zap.String("function", function),
		zap.Error(err),
	)
	log.Fatal(msg, fields...)
}

// Sync flushes any buffered log entries
func Sync() error {
	if log != nil {
		return log.Sync()
	}
	return nil
}

// WithContext creates a logger with additional context fields
func WithContext(fields ...zap.Field) *zap.Logger {
	if log == nil {
		Init()
	}
	return log.With(fields...)
}

// Helper functions for common field types
func String(key, val string) zap.Field {
	return zap.String(key, val)
}

func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func Bool(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}

func Any(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}
