package build

import (
	"github.com/google/wire"
	"go-slim/internal/cron"
)

var CronNewSet = wire.NewSet(
	wire.Struct(new(cron.Tasks), "*"),
	wire.Struct(new(cron.Manager), "Tasks"),
	//cron.NewCleanTmpFileTask,
	wire.Struct(new(cron.DemoTask), "*"),
)
