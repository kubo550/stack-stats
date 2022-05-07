package routes

import (
	"github.com/gofiber/fiber/v2"
	"stats/src/log"
)

func HealthHandler(c *fiber.Ctx) error {
	log.Info("HealthHandler - Health check")
	return c.SendStatus(200)
}
