package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-slim/api/server/controllers/v1/common"
	"go-slim/internal/constants"
	"gorm.io/gorm"
	"time"
)

// Server
// @basePath /
type Server struct {
	Redis *redis.Client
	Mysql *gorm.DB
}

// Timestamp
// @Summary sync time
// @Description get server time
// @Tags       Server
// @Accept      json
// @Produce     json
// @Param       timezone  query     string false "timezone"
// @Success     200 {object} string
// @Router      /api/v1/server/timestamp [get]
func (s *Server) Timestamp(c *gin.Context) {
	now := time.Now()
	common.Success(c, gin.H{
		"timestamp": now.Unix(),
		"datetime":  now.Format(constants.TimeFormat),
	})
}

func (s *Server) Health(c *gin.Context) {
	var rd, my int
	if _, err := s.Redis.Ping(c).Result(); err == nil {
		rd = 1
	}
	db, _ := s.Mysql.DB()
	if err := db.Ping(); err == nil {
		my = 1
	}
	common.Success(c, gin.H{
		"redis": rd,
		"mysql": my,
	})
}
