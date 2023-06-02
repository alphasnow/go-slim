package build

import (
	"github.com/google/wire"
	"go-slim/pkg/xaliyun"
	"go-slim/pkg/xbaidu"
	"go-slim/pkg/xcaptcha"
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
	"go-slim/pkg/xsnowflake"
	"go-slim/pkg/xtencent"
)

var PkgNewSet = wire.NewSet(

	// http
	xhttp.NewGin,
	xhttp.NewHttp,
	wire.Struct(new(xhttp.Http), "*"),
	// redis
	xredis.NewRedis,
	// log
	xlog.NewLogger,
	xlog.NewManager,
	// mysql
	xmysql.NewMysql,
	xmysql.NewGenerator,
	// jwt
	xjwt.NewToken,
	// limiter
	xlimiter.NewLimiter,
	// crypt
	xcrypt.NewAesCbc,
	// snowflake
	xsnowflake.NewSnowFlake,
	// optimus
	xoptimus.NewOptimus,
	// paddle
	xpaddle.NewClient,
	wire.Struct(new(xpaddle.HumansegMobileClient), "*"),
	// captcha
	xcaptcha.NewStringCaptcha,
	// baidu
	xbaidu.NewOauth,
	xbaidu.NewClient,
	// tencent
	xtencent.NewClientProfile,
	xtencent.NewCredential,
	xtencent.NewFtClient,
	// aliyun
	xaliyun.NewOpenApiConfig,
	xaliyun.NewImageSeg,
	// aliyun oss
	xoss.NewClient,
	xoss.NewBucket,
)
