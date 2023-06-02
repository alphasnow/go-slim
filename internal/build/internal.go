package build

import (
	"github.com/google/wire"
)

var InternalNewSet = wire.NewSet(
	ConfigNewSet,
	CronNewSet,
	QueueNewSet,
	RouteNewSet,
	ServicesNewSet,
)
