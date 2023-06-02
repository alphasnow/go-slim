package routes

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	v1 "go-slim/api/server/controllers/v1"
)

type ApiV1Route struct {
	Gin *gin.Engine
	Ctl *v1.Controllers
}

func (r *ApiV1Route) Register() {

	api1 := r.Gin.Group("/api/v1")
	r.groupMiddleware(api1)

	server := api1.Group("/server")
	{
		server.GET("timestamp", r.Ctl.Server.Timestamp)
		server.GET("health", r.Ctl.Server.Health)
	}

}

func (r *ApiV1Route) groupMiddleware(route *gin.RouterGroup) {

	// REQUEST-ID
	route.Use(requestid.New())
}
