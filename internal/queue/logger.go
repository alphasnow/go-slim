package queue

import (
	"fmt"
	"github.com/hibiken/asynq"
	"go-slim/pkg/xlog"
	"go.uber.org/zap"
)

type Logger struct {
	Zap *zap.Logger
}

func NewLogger(logManager *xlog.Manager) *Logger {
	zapLog := logManager.Logger(xlog.QUEUE)
	return &Logger{Zap: zapLog}
}

func (l *Logger) argsToMsg(args ...interface{}) string {
	msg := ""
	for _, v := range args {
		msg += fmt.Sprintf("%v ", v)
	}
	return msg
}
func (l *Logger) Debug(args ...interface{}) {
	l.Zap.Debug(l.argsToMsg(args))
}

func (l *Logger) Info(args ...interface{}) {
	l.Zap.Info(l.argsToMsg(args))
}

func (l *Logger) Warn(args ...interface{}) {
	l.Zap.Warn(l.argsToMsg(args))
}

func (l *Logger) Error(args ...interface{}) {
	l.Zap.Error(l.argsToMsg(args))
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Zap.Fatal(l.argsToMsg(args))
}

var _ asynq.Logger = (*Logger)(nil)
