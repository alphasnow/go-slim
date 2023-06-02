package config

import (
	"go-slim/internal/app"
	"go-slim/pkg/xaliyun"
	"go-slim/pkg/xbaidu"
	"go-slim/pkg/xcrypt"
	"go-slim/pkg/xhttp"
	"go-slim/pkg/xjwt"
	"go-slim/pkg/xlimiter"
	"go-slim/pkg/xlog"
	"go-slim/pkg/xmysql"
	"go-slim/pkg/xoptimus"
	"go-slim/pkg/xoss"
	"go-slim/pkg/xpaddle"
	"go-slim/pkg/xredis"
	"go-slim/pkg/xsession"
	"go-slim/pkg/xsnowflake"
	"go-slim/pkg/xtencent"
)

type Config struct {
	Http      *xhttp.Config      `yaml:"http"`
	Redis     *xredis.Config     `yaml:"redis"`
	Log       *xlog.Config       `yaml:"log"`
	Mysql     *xmysql.Config     `yaml:"mysql"`
	JWT       *xjwt.Config       `yaml:"jwt"`
	Limiter   *xlimiter.Config   `yaml:"limiter"`
	Session   *xsession.Config   `yaml:"session"`
	Crypt     *xcrypt.Config     `yaml:"crypt"`
	SnowFlake *xsnowflake.Config `yaml:"snowFlake"`
	Paddle    *xpaddle.Config    `yaml:"paddle"`
	Optimus   *xoptimus.Config   `yaml:"optimus"`
	Baidu     *xbaidu.Config     `yaml:"baidu"`
	Tencent   *xtencent.Config   `yaml:"tencent"`
	Aliyun    *xaliyun.Config    `yaml:"aliyun"`
	Oss       *xoss.Config       `yaml:"oss"`
	App       *app.Config        `yaml:"app"`
}

func DefaultConfig() *Config {
	return &Config{
		Http:      xhttp.NewConfig(),
		Redis:     xredis.NewConfig(),
		Log:       xlog.NewConfig(),
		Mysql:     xmysql.NewConfig(),
		JWT:       xjwt.NewConfig(),
		Limiter:   xlimiter.NewConfig(),
		Session:   xsession.NewConfig(),
		Crypt:     xcrypt.NewConfig(),
		SnowFlake: xsnowflake.NewConfig(),
		Paddle:    xpaddle.NewConfig(),
		Optimus:   xoptimus.NewConfig(),
		Baidu:     xbaidu.NewConfig(),
		Tencent:   xtencent.NewConfig(),
		Aliyun:    xaliyun.NewConfig(),
		Oss:       xoss.NewConfig(),
		App:       app.NewConfig(),
	}
}
