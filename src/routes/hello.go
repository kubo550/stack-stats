package routes

import (
	"github.com/gofiber/fiber/v2"
	"stats/src/consoleLog"
)

func HelloHandler(c *fiber.Ctx) error {
	consoleLog.Info("HelloHandler - new request")

	return c.Status(fiber.StatusOK).SendFile("README.md")
}
