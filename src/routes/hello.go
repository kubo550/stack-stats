package routes

import (
	"github.com/gofiber/fiber/v2"
	"stats/src/log"
)

func HelloHandler(c *fiber.Ctx) error {
	log.Info("HelloHandler - new request")

	return c.Status(fiber.StatusOK).SendString("Follow this repo to see the code: https://github.com/kubo550/stack-stats")
}
