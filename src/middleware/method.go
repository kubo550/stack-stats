package middleware

import "github.com/gofiber/fiber/v2"

func MethodMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Method() != "GET" {
			c.Status(fiber.StatusMethodNotAllowed)
			return c.SendString("Sorry, only GET requests are allowed.")
		}
		return c.Next()
	}
}
