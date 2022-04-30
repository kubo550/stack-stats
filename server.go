package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func healthHandler(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		if c.Method() != "GET" {
			c.Status(fiber.StatusMethodNotAllowed)
			return c.SendString("Sorry, only GET requests are allowed.")
		}
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("New request")

		return c.SendString("Hello, World!")
	})

	app.Get("/health", healthHandler)

	app.Use(statsMiddleware())

	app.Get("/stats", statsHandler)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}

func statsMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Query("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "id is required",
			})
		}
		return c.Next()
	}
}

type stats struct {
	ID         string `json:"id"`
	Name       string `json:"fullName"`
	Reputation int    `json:"reputation"`
	Gold       int    `json:"gold"`
	Silver     int    `json:"silver"`
	Bronze     int    `json:"bronze"`
	ImageUrl   string `json:"imageUrl"`
}

type StackStats struct {
	Items []struct {
		DisplayName  string `json:"display_name"`
		ProfileImage string `json:"profile_image"`
		Reputation   int    `json:"reputation"`
		BadgeCounts  struct {
			Bronze int `json:"bronze"`
			Gold   int `json:"gold"`
			Silver int `json:"silver"`
		} `json:"badge_counts"`
	} `json:"items"`
}

type Theme struct {
	TitleColor string `json:"title_color"`
	IconColor  string `json:"icon_color"`
	TextColor  string `json:"text_color"`
	BgColor    string `json:"bg_color"`
}

func statsHandler(c *fiber.Ctx) error {
	userId := c.Query("id")
	fmt.Println("User id:", userId)

	stackStats := getStackStats(userId)
	theme := Theme{"9745f5", "9f4bff", "ffffff", "000000"}

	svg := generateSVG(stackStats, theme)

	//c.Set(fiber.HeaderContentType, "image/svg+xml; charset=utf-8")

	return c.JSON(stackStats)
}



func getStackStats(userId string) stats {
	client := &http.Client{}
	req, err := http.NewRequest("GET", getStackApiUrl(userId), nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var stackStats StackStats
	err = json.Unmarshal(responseBody, &stackStats)
	if err != nil {
		log.Fatal(err)
	}

	return stats{
		ID:         userId,
		Name:       stackStats.Items[0].DisplayName,
		Reputation: stackStats.Items[0].Reputation,
		Gold:       stackStats.Items[0].BadgeCounts.Gold,
		Silver:     stackStats.Items[0].BadgeCounts.Silver,
		Bronze:     stackStats.Items[0].BadgeCounts.Bronze,
		ImageUrl:   stackStats.Items[0].ProfileImage,
	}
}

func getStackApiUrl(userId string) string {
	return fmt.Sprintf("https://api.stackexchange.com/2.3/users/%s?order=desc&sort=reputation&site=stackoverflow", userId)
}

