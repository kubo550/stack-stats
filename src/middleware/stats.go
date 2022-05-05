package middleware

import (
	"github.com/gofiber/fiber/v2"
	"stats/src/consoleLog"
)

func StatsMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Query("id")
		if id == "" {
			consoleLog.Warning("StatsMiddleware - No id provided")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "id is required",
			})
		}
		return c.Next()
	}
}
