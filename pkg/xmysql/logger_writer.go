package xmysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Writer struct {
	Zap *zap.Logger
}

func (w *Writer) Printf(s string, i ...interface{}) {
	res := fmt.Sprintf(s, i...)
	w.Zap.Info(res)
}

var _ logger.Writer = (*Writer)(nil)
