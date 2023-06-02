package xlog

import (
	"go-slim/internal/app"
	"go-slim/pkg/xlog/rotation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sort"
	"time"
)

func NewLogger(cfg *Config, g *app.Config) *zap.Logger {
	cfg.Rotation.Filepath = g.LogsPath
	cfg.mergeEnv()

	// https://github.com/uber-go/zap/blob/master/FAQ.md#does-zap-support-log-rotation
	encoder := newEncoder(cfg.Encoding, cfg.EncoderConfig)
	jack := rotation.NewRotation(cfg.Rotation)
	writer := zapcore.AddSync(jack)

	core := zapcore.NewCore(
		encoder,
		writer,
		cfg.Level,
	)
	opts := buildOptions(cfg)
	logger := zap.New(core, opts...)

	return logger
}

func newEncoder(name string, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	var encoder zapcore.Encoder

	switch name {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig) // json格式
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	}

	return encoder
}

func buildOptions(cfg *Config) []zap.Option {
	var opts []zap.Option

	if cfg.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zap.ErrorLevel
	if cfg.Development {
		stackLevel = zap.WarnLevel
	}
	if !cfg.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	if scfg := cfg.Sampling; scfg != nil {
		opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			var samplerOpts []zapcore.SamplerOption
			if scfg.Hook != nil {
				samplerOpts = append(samplerOpts, zapcore.SamplerHook(scfg.Hook))
			}
			return zapcore.NewSamplerWithOptions(
				core,
				time.Second,
				cfg.Sampling.Initial,
				cfg.Sampling.Thereafter,
				samplerOpts...,
			)
		}))
	}

	if len(cfg.InitialFields) > 0 {
		fs := make([]zap.Field, 0, len(cfg.InitialFields))
		keys := make([]string, 0, len(cfg.InitialFields))
		for k := range cfg.InitialFields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fs = append(fs, zap.Any(k, cfg.InitialFields[k]))
		}
		opts = append(opts, zap.Fields(fs...))
	}

	return opts
}
