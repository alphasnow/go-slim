package xmysql

import (
	"go.uber.org/zap"
)

// Logger
type Logger struct {
	Zap *zap.Logger
}

// var _ logger.Interface = (*Logger)(nil)
