package home

import (
	"github.com/google/wire"
	"go-slim/web/home/controllers"
)

type Controllers struct {
	Home *controllers.Home
}

var WebHomeNewSet = wire.NewSet(

	wire.Struct(new(Controllers), "*"),
	wire.Struct(new(controllers.Home), "*"),

	wire.Struct(new(WebHomeRoute), "*"),
)
