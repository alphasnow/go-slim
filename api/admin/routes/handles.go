package routes

import (
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/schema"
	"net/http"
)

func HandleNoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Response{
		Success:      false,
		ErrorCode:    404,
		ErrorMessage: "No Route",
	})
}
func HandleNoMethod(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Response{
		Success:      false,
		ErrorCode:    405,
		ErrorMessage: "No Method",
	})
}
