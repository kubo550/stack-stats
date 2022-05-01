package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"stats/src/structs"
	"stats/src/utils"
)

func StatsHandler(c *fiber.Ctx) error {
	userId := c.Query("id")
	fmt.Println("User id:", userId)

	stackStats := utils.GetStackStats(userId)
	fmt.Println("Stack stats:", stackStats)

	theme := structs.Theme{Gold: "#F0B400", Silver: "#999C9F", Bronze: "#AB8A5F", BgColor: "#2D2D2D", TextColor: "#C4CCBC"}

	svg := utils.GenerateSVG(stackStats, theme)

	c.Set(fiber.HeaderContentType, "image/svg+xml; charset=utf-8")

	return c.SendString(svg)
}
