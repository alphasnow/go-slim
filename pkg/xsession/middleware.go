package xsession

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewMiddleware(config *Config) gin.HandlerFunc {
	mid := func(c *gin.Context) {
		c.Next()
	}
	var store sessions.Store
	switch config.Store {
	// todo: redis
	case "cookie":
		store = cookie.NewStore([]byte(config.Cookie.Secret))
	}
	mid = sessions.Sessions(config.Name, store)
	return mid
}
