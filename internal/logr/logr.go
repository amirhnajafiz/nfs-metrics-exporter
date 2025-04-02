package logr

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger creates a new zap logger instance
func NewZapLogger(debug bool) *zap.Logger {
	// create a new zap logger instance based on the debug flag
	var level zapcore.Level
	if debug {
		level = zapcore.DebugLevel
	} else {
		level = zapcore.InfoLevel
	}

	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
