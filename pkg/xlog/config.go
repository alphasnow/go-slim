package xlog

import (
	"go-slim/pkg/xlog/rotation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	zap.Config `yaml:",inline"`
	Rotation   *rotation.Config `yaml:"rotation"`
}

func NewConfig() *Config {
	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg := &Config{Config: zapCfg, Rotation: &rotation.Config{
		Filename:   DEFAULT + ".log",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 100,
		LocalTime:  true,
		Compress:   true,
		Filepath:   "./logs",
	}}
	return cfg
}

func (c *Config) mergeEnv() {
	level := os.Getenv("LOG_LEVEL")
	if level != "" {
		l := zap.AtomicLevel{}
		_ = l.UnmarshalText([]byte(level))
		c.Config.Level = l
	}
}
