package xhttp

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func NewGin(cfg *Config) *gin.Engine {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	return gin.New()
}

func NewHttp(cfg *Config, handler *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:         cfg.GetAddr(),
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	return server
}
