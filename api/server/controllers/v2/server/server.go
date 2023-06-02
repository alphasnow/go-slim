package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-slim/api/server/controllers/v2/common"
	"gorm.io/gorm"
	"time"
)

type Server struct {
	Redis *redis.Client
	Mysql *gorm.DB
}

type ServerTimestampResp struct {
	Date string `json:"date"`
	Time int64  `json:"timestamp"`
}

func (ctl *Server) Timestamp(c *gin.Context) {
	now := time.Now()
	resp := &ServerTimestampResp{
		Date: now.Format("2006-01-02 15:04:05"),
		Time: now.Unix(),
	}
	common.Success(c, resp)
}

type ServerHealthResp struct {
	Redis bool `json:"redis"`
	Mysql bool `json:"mysql"`
}

func (ctl *Server) Health(c *gin.Context) {
	var resp ServerHealthResp

	if _, err := ctl.Redis.Ping(c).Result(); err == nil {
		resp.Redis = true
	}

	db, _ := ctl.Mysql.DB()
	if err := db.Ping(); err == nil {
		resp.Mysql = true
	}

	common.Success(c, resp)
}
