package routes

import (
	"github.com/gofiber/fiber/v2"
	"stats/src/consoleLog"
)

func HealthHandler(c *fiber.Ctx) error {
	consoleLog.Info("HealthHandler - Health check")
	return c.SendStatus(200)
}
