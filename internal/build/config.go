package build

import (
	"github.com/google/wire"
	"go-slim/internal/config"
)

var ConfigNewSet = wire.NewSet(
	wire.FieldsOf(
		new(*config.Config),
		"Http",
		"Redis",
		"Log",
		"Mysql",
		"JWT",
		"Limiter",
		"Session",
		"Crypt",
		"SnowFlake",
		"Paddle",
		"Optimus",
		"Baidu",
		"Aliyun",
		"Tencent",
		"Oss",
		"App",
	),
)
