package main

import (
	"fileSearch/api/config"
	"fileSearch/api/handlers"
	"fileSearch/api/rpc"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	setupRouters(app)

	rpc.InitRpcClient()

	app.Listen(config.GConfig.Api.Host + ":" + config.GConfig.Api.Port)
}

func setupRouters(app *fiber.App) {
	app.Get("/", handlers.IndexHander)
	group := app.Group("/api/v1")
	{
		group.Get("/search/:word", handlers.GetSearchHandle)
		group.Get("/echo/:word", handlers.EchoHandle)
	}
}
