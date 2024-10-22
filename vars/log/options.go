package log

import (
	"go.uber.org/zap/zapcore"
)

const (
	FormatConsole = "console"
	FormatJSON    = "json"

	OutputStd    = "stdout"
	OutputStdErr = "stderr"
)

type Heheh interface {
}

type Options struct {
	OutputPaths      []string `json:"output_paths" mapstructure:"output_paths"`
	ErrorOutputPaths []string `json:"error_output_paths" mapstructure:"error_output_paths"`
	Level            string   `json:"level" mapstructure:"level"`
	Format           string   `json:"format" mapstructure:"format"`
	Name             string   `json:"name" mapstructure:"name"`
	EnableColor      bool     `json:"enable_color" mapstructure:"enable_color"`
	Develop          bool     `json:"develop" mapstructure:"develop"`
	EnableTraceID    bool     `json:"enable_trace_id" mapstructure:"enable_trace_id"`
	EnableTraceStack bool     `json:"enable_trace_stack" mapstructure:"enable_trace_stack"`
}

type Option func(*Options)

func WithLevel(level string) Option {
	return func(ops *Options) {
		ops.Level = level
	}
}

func WithFormat(format string) Option {
	return func(ops *Options) {
		ops.Format = format
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		Level:            zapcore.InfoLevel.String(),
		Format:           FormatConsole,
		OutputPaths:      []string{OutputStd},
		ErrorOutputPaths: []string{OutputStdErr},
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
