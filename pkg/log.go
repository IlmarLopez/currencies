package log

import (
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
func NewLogger(level string) (*zap.SugaredLogger, error) {
	var _level zap.AtomicLevel

	switch strings.ToLower(level) {
	case "debug":
		_level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		_level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		_level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		_level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		_level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	cfg := NewLoggerConfig(_level)
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
