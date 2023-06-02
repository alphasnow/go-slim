package middlewares

import (
	libgin "github.com/gin-gonic/gin"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"go-slim/api/server/schema"
	"net/http"
)

func NewLimitReachedHandler() mgin.LimitReachedHandler {
	return func(c *libgin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, schema.NewRes(schema.TooManyRequests))
	}
}
