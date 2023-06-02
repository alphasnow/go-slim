package xhttp

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Http struct {
	Server *http.Server
}

func (s *Http) Gin() *gin.Engine {
	return s.Server.Handler.(*gin.Engine)
}

func (s *Http) Run() {
	err := s.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (s *Http) Close() {
	ctx := context.Background()
	s.Server.SetKeepAlivesEnabled(false)
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Println("http close error: " + err.Error())
	}
}
