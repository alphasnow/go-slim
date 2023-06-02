package build

import (
	"github.com/google/wire"
	"go-slim/internal/routes"
)

var RouteNewSet = wire.NewSet(
	wire.Struct(new(routes.Routes), "*"),
)
