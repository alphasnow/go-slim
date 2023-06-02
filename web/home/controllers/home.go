package controllers

import (
	"github.com/gin-gonic/gin"
	"go-slim/internal/services"
	"net/http"
	"time"
)

type Home struct {
	UrlGen *services.UrlGen
}

func (ctl *Home) Index(c *gin.Context) {
	bs := ctl.UrlGen.Assets("bootstrap.min.css")
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"now": time.Now().Format("2006-01-02 15:04:05"),
		"bs":  bs,
	})
}
