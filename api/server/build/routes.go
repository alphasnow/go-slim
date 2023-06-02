package build

import (
	"github.com/google/wire"
	"go-slim/api/server/routes"
)

var RoutesNewSet = wire.NewSet(
	wire.Struct(new(routes.ApiV1Route), "*"),
	wire.Struct(new(routes.ApiV2Route), "*"),
)
