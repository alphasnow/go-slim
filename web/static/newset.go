package static

import (
	"github.com/google/wire"
	"go-slim/web/static/services/sign_checker"
)

var WebStaticNewSet = wire.NewSet(
	wire.Struct(new(WebStaticRoute), "*"),

	wire.Struct(new(sign_checker.SignChecker), "*"),
	wire.Struct(new(sign_checker.SignUrlChecker), "*"),
)
