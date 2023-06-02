package build

import (
	"github.com/google/wire"
	"go-slim/api/admin/middlewares"
	"go-slim/api/admin/routes"
	"go-slim/api/admin/services"
)

var ApiAdminNewSet = wire.NewSet(
	wire.Struct(new(routes.ApiAdminRoute), "*"),

	middlewares.NewTokenChecker,

	routes.ControllersNewSet,

	// services
	wire.Struct(new(services.FormFile), "*"),
)
