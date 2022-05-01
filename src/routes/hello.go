package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func HelloHandler(c *fiber.Ctx) error {
	fmt.Println("New request")
	return c.SendString("Hello, World!")
}
