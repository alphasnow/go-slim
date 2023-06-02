package build

import (
	"github.com/google/wire"
	"go-slim/internal/services"
)

var ServicesNewSet = wire.NewSet(
	services.NewUrlGen,
)
