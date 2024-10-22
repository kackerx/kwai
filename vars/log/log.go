package log

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(ctx context.Context, msg string)
	Debugf(ctx context.Context, format string, args ...any)
	DebugW(ctx context.Context, keysAndValues ...any)
}

type ZapLogger struct {
	zapLogger *zap.Logger
}

var (
	logger *ZapLogger
	once   sync.Once
)

func Init(opts *Options) {
	once.Do(func() {
		logger = NewZapLogger(opts)
	})
}

func NewZapLogger(opts *Options) *ZapLogger {
	// zap实例化
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		fmt.Println("default debug")
		zapLevel = zapcore.DebugLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		Development:      true,
		Encoding:         "console", // 或 "json"
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel))
	if err != nil {
		panic(err)
	}

	return &ZapLogger{l}
}

func (z *ZapLogger) Debug(ctx context.Context, msg string) {
	z.zapLogger.Debug("level: " + z.zapLogger.Level().String())
}

func (z *ZapLogger) Debugf(ctx context.Context, format string, args ...any) {
	// TODO implement me
	panic("implement me")
}

func (z *ZapLogger) DebugW(ctx context.Context, keysAndValues ...any) {
	// TODO implement me
	panic("implement me")
}

func Debug(ctx context.Context, msg string) {
	logger.Debug(ctx, msg)
}
