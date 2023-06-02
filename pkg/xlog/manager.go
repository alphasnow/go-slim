package xlog

import (
	"go-slim/internal/app"
	"go.uber.org/zap"
)

const (
	DEFAULT = "app"
	QUEUE   = "queue"
	CRON    = "cron"
	HTTP    = "http"
	SQL     = "sql"
)

type Manager struct {
	Cfg     *Config
	LogPath string
	loggers map[string]*zap.Logger
}

func NewManager(cfg *Config, g *app.Config) *Manager {
	return &Manager{
		cfg,
		g.LogsPath,
		map[string]*zap.Logger{},
	}
}

func (m *Manager) NewLogger(name string) *zap.Logger {
	xLogConfig := *m.Cfg
	xLogConfig.Rotation.Filename = name + ".log"
	xGlobal := app.Config{LogsPath: m.LogPath}
	xLog := NewLogger(&xLogConfig, &xGlobal)
	return xLog
}

func (m *Manager) Logger(name string) *zap.Logger {
	_, ok := m.loggers[name]
	if ok == false {
		m.loggers[name] = m.NewLogger(name)
	}
	return m.loggers[name]
}
