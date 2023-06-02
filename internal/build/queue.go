package build

import (
	"github.com/google/wire"
	"go-slim/internal/queue"
)

var QueueNewSet = wire.NewSet(
	wire.Struct(new(queue.Manager), "RdsOpt", "Conf", "Tasks"),
	wire.Struct(new(queue.Tasks), "*"),
	queue.GetRedisClientOpt,
	queue.GetConfig,
	queue.NewLogger,
	wire.Struct(new(queue.DemoTask), "*"),
)
