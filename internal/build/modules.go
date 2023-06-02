package build

import (
	"github.com/google/wire"
	admin "go-slim/api/admin/build"
	server "go-slim/api/server/build"
	"go-slim/web/home"
	"go-slim/web/static"
)

var ModulesNewSet = wire.NewSet(
	admin.ApiAdminNewSet,
	server.ApiServerNewSet,

	home.WebHomeNewSet,
	static.WebStaticNewSet,
)
