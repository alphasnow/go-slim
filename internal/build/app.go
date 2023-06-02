package build

import (
	"github.com/go-redis/redis/v8"
	"go-slim/internal/cron"
	"go-slim/internal/models"
	"go-slim/internal/queue"
	"go-slim/internal/routes"
	"go-slim/pkg/xhttp"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type App struct {
	Http   *xhttp.Http
	Routes *routes.Routes
	Queue  *queue.Manager
	Cron   *cron.Manager

	DB    *gorm.DB
	Redis *redis.Client
}

func (a *App) Initialize() {
	if err := models.AutoMigrate(a.DB); err != nil {
		panic(err)
	}
	if err := models.AutoSeeder(a.DB); err != nil {
		panic(err)
	}

	a.Routes.Register()
	a.Cron.StartAsync()
	a.Queue.Start()
}

type Cli struct {
	Queue   *queue.Manager
	Cron    *cron.Manager
	GormGen *gen.Generator
}
