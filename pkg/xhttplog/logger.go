package xhttplog

import (
	"go-slim/internal/app"
	"go-slim/pkg/xlog"
	"go.uber.org/zap"
)

type Logger zap.Logger

// NewHttpLogger
// Deprecated: Use LogManager
func NewHttpLogger(xGlobal *app.Config, logConfig *xlog.Config) *Logger {
	xLogConfig := *logConfig
	xLogConfig.Rotation.Filename = "http.log"
	xLog := xlog.NewLogger(&xLogConfig, xGlobal)
	return (*Logger)(xLog)
}
