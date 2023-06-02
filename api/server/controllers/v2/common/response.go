package common

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go-slim/api/server/schema"
	"net/http"
)

func Error(c *gin.Context, code schema.Code, err ...error) {
	response := schema.NewRes(code, err...)
	response.RequestID = requestid.Get(c)
	c.AbortWithStatusJSON(http.StatusOK, response)
}
func Success(c *gin.Context, data interface{}) {
	response := schema.NewDataRes(data)
	response.RequestID = requestid.Get(c)
	c.JSON(http.StatusOK, response)
}
