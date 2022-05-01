package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"stats/src/middleware"
	"stats/src/routes"
)

func main() {
	app := fiber.New()

	app.Use(middleware.MethodMiddleware())

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	app.Get("/", routes.HelloHandler)
	app.Get("/health", routes.HealthHandler)

	app.Use(middleware.StatsMiddleware())
	app.Get("/stats", routes.StatsHandler)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
