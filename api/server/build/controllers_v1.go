package build

import (
	"github.com/google/wire"
	"go-slim/api/server/controllers/v1"
	"go-slim/api/server/controllers/v1/server"
)

var ControllersV1NewSet = wire.NewSet(
	wire.Struct(new(v1.Controllers), "*"),

	wire.Struct(new(server.Server), "*"),
)
