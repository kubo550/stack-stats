package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"stats/src/consoleLog"
	"stats/src/middleware"
	"stats/src/routes"
)

func main() {
	app := fiber.New()

	//app.Use(middleware.MethodMiddleware())

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Europe/Warsaw",
	}))

	app.Get("/", routes.HelloHandler)
	app.Get("/health", routes.HealthHandler)

	app.Use(middleware.StatsMiddleware())
	app.Get("/stats", routes.StatsHandler)

	// default port from env
	port := os.Getenv("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		consoleLog.Error(err)
		panic(err)
	}
}
