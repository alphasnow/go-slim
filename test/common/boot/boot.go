package boot

import (
	"go-slim/internal/app"
	"go-slim/internal/build"
	"go-slim/internal/config"
	"go-slim/test/common"
	"math/rand"
	"net/http"
	"time"
)

var App *build.App

func init() {
	var err error
	rand.Seed(time.Now().UnixNano())
	cfg := config.NewConfig(app.ConfigsPath, app.RootPath)
	cfg.App.AppEnv = app.EnvTest
	App, err = build.BuildApp(cfg)
	if err != nil {
		panic(err)
	}
	App.Initialize()
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	App.Http.Gin().ServeHTTP(w, req)
}

func Client() *common.HttpClient {
	return common.NewClient(App.Http.Gin())
}
