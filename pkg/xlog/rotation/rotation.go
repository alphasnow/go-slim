package rotation

import (
	"github.com/natefinch/lumberjack"
)

func NewRotation(cfg *Config) *lumberjack.Logger {

	jackLogger := &lumberjack.Logger{
		Filename:   cfg.File(),
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge, // days
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}
	return jackLogger
}
