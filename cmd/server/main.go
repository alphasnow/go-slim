package main

import (
	"context"
	"fmt"
	"go-slim/internal/app"
	"go-slim/internal/build"
	"go-slim/internal/config"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func main() {
	defer catchPanic()

	log.Println("[SERVER]", "application create")
	application := createApp()

	log.Println("[SERVER]", "server running", application.Http.Server.Addr)
	go application.Http.Run()

	<-catchExit()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	log.Println("[SERVER]", "server wait 3 seconds")

	go closeApp(application, cancel)

	<-ctx.Done()
	log.Println("[SERVER]", "server closed")
}

func catchPanic() {
	if e := recover(); e != nil {
		err := fmt.Sprintf("%v \n %s", e, debug.Stack())

		_ = ioutil.WriteFile(app.LogsPath+"/server.log", []byte(err), 0666)

		log.Println("[SERVER]", "server panic", err)

		log.Println("[SERVER]", "server will close in a minute")
		time.Sleep(60 * time.Second)
	}
}

func createApp() *build.App {
	cfg := config.NewConfig(app.ConfigsPath, app.RootPath)
	application, err := build.BuildApp(cfg)
	if err != nil {
		panic(err)
	}
	application.Initialize()
	return application
}

func closeApp(application *build.App, cancel context.CancelFunc) {
	defer cancel()

	// http
	application.Http.Close()

	// db
	if db, err := application.DB.DB(); err == nil {
		if err = db.Close(); err != nil {
			log.Println("db close error: " + err.Error())
		}
	} else {
		log.Println("db close error: " + err.Error())
	}

	// redis
	if err := application.Redis.Close(); err != nil {
		log.Println("redis close error: " + err.Error())
	}

	// queue
	application.Queue.Close()

	// cron
	application.Cron.Close()

	log.Println("[SERVER]", "server close all service")
}

func catchExit() chan os.Signal {
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	return quit
}
