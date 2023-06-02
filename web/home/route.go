package home

import (
	"github.com/gin-gonic/gin"
	"go-slim/internal/app"
)

type WebHomeRoute struct {
	Gin         *gin.Engine
	Controllers *Controllers
	Global      *app.Config
}

func (r *WebHomeRoute) Register() {
	ginWeb := r.Gin.Group("")

	ginWeb.GET("/", r.Controllers.Home.Index)
}
