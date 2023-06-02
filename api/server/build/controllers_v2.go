package build

import (
	"github.com/google/wire"
	v2 "go-slim/api/server/controllers/v2"
	"go-slim/api/server/controllers/v2/auth"
	"go-slim/api/server/controllers/v2/open"
	"go-slim/api/server/controllers/v2/server"
	"go-slim/api/server/controllers/v2/upload"
	"go-slim/api/server/controllers/v2/verify"
)

var ControllersV2NewSet = wire.NewSet(
	wire.Struct(new(v2.Controllers), "*"),

	wire.Struct(new(upload.SingleFile), "*"),
	wire.Struct(new(upload.RangeFile), "*"),
	wire.Struct(new(upload.Check), "*"),
	wire.Struct(new(upload.Rules), "*"),

	wire.Struct(new(server.Server), "*"),

	wire.Struct(new(open.Paddle), "*"),
	wire.Struct(new(open.Aliyun), "*"),
	wire.Struct(new(open.Baidu), "*"),
	wire.Struct(new(open.Tencent), "*"),

	wire.Struct(new(auth.AppClient), "*"),
	wire.Struct(new(auth.Username), "*"),

	wire.Struct(new(verify.Captcha), "*"),
)
