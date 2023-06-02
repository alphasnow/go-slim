package cron

import (
	"go-slim/internal/app"
	"go-slim/pkg/xlog"
	"go.uber.org/zap"
)

type Logger zap.Logger

// NewCronLogger
// Deprecated: Use LogManager
func NewCronLogger(xGlobal *app.Config, logConfig *xlog.Config) *Logger {
	xLogConfig := *logConfig
	xLogConfig.Rotation.Filename = "cron.log"
	xLog := xlog.NewLogger(&xLogConfig, xGlobal)
	return (*Logger)(xLog)
}

func Zap(log *Logger) *zap.Logger {
	return (*zap.Logger)(log)
}
