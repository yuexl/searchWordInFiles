package main

import (
	"net/http"
	"net/http/pprof"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"

	"fileSearch/log"

	"fileSearch/api/config"
	"fileSearch/api/handlers"
	"fileSearch/api/rpc"
)

func SetPprof() {
	go func() {
		http.HandleFunc("/debug/pprof/block", pprof.Index)
		http.HandleFunc("/debug/pprof/goroutine", pprof.Index)
		http.HandleFunc("/debug/pprof/heap", pprof.Index)
		http.HandleFunc("/debug/pprof/threadcreate", pprof.Index)

		http.ListenAndServe("localhost:8888", nil)
	}()
}

func main() {
	//SetPprof()

	app := fiber.New()

	setupRouters(app)

	rpc.InitRpcClient()

	log.GLogger.WithFields(logrus.Fields{
		"host": config.GConfig.Api.Host,
		"port": config.GConfig.Api.Port,
	}).Infoln("api start listing")

	//app.Listen(config.GConfig.Api.Host + ":" + config.GConfig.Api.Port)
	app.Listen(":" + config.GConfig.Api.Port)
}

func setupRouters(app *fiber.App) {
	app.Get("/", handlers.IndexHander)
	group := app.Group("/api/v1")
	{
		group.Get("/search/:word", handlers.GetSearchHandle)
		group.Get("/echo/:word", handlers.EchoHandle)
	}
}
