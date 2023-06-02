package build

import (
	"github.com/google/wire"
	"go-slim/api/server/middlewares"
)

var MiddlewaresNewSet = wire.NewSet(
	middlewares.NewLimitReachedHandler,
)
