package build

import "github.com/google/wire"

var ApiServerNewSet = wire.NewSet(
	ControllersV1NewSet,
	ControllersV2NewSet,
	LogicNewSet,
	MiddlewaresNewSet,
	RoutesNewSet,
	ServicesNewSet,
)
