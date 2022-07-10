package log

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLoggerConfig returns a new LoggerConfig with the given level.
func NewLoggerConfig(level zap.AtomicLevel) zap.Config {
	cfg := zap.Config{
		Level:            level,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    "function",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	return cfg
}

// NewLogger returns a new Logger with the given level.
func NewLogger() (*zap.SugaredLogger, error) {
	var level zap.AtomicLevel

	if logLevel, ok := os.LookupEnv("LOG_LEVEL"); ok {
		switch strings.ToLower(logLevel) {
		case "debug":
			level = zap.NewAtomicLevelAt(zap.DebugLevel)
		case "info":
			level = zap.NewAtomicLevelAt(zap.InfoLevel)
		case "warn":
			level = zap.NewAtomicLevelAt(zap.WarnLevel)
		case "error":
			level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		default:
			level = zap.NewAtomicLevelAt(zap.InfoLevel)
		}
	}

	cfg := NewLoggerConfig(level)
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
