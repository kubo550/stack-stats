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

	app.Get("/Stats", statsHandler)

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

type Stats struct {
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
	theme := Theme{"#9745f5", "#9f4bff", "#ffffff", "#000000"}

	svg := generateSVG(stackStats, theme)

	c.Set(fiber.HeaderContentType, "image/svg+xml; charset=utf-8")

	return c.SendString(svg)
}

func generateSVG(stackStats Stats, theme Theme) string {
	const width = "158"
	const height = "47"

	var svg string

	// data-testId="gold"

	svg += `<svg width="` + width + `" height="` + height + `" viewBox="0 0 ` + width + ` ` + height + `" fill="none" xmlns="http://www.w3.org/2000/svg">`
	svg += `<rect width="` + width + `" height="` + height + `" fill="#2D2D2D"/>`
	svg += `<path data-testId="gold" d="M47.5577 19.2727V28H45.9767V20.8111H45.9256L43.8844 22.1151V20.6662L46.0534 19.2727H47.5577ZM51.1213 26.8068L51.0659 27.2756C51.0261 27.6335 50.9551 27.9972 50.8528 28.3665C50.7534 28.7386 50.6483 29.081 50.5375 29.3935C50.4267 29.706 50.3372 29.9517 50.269 30.1307H49.2292C49.269 29.9574 49.323 29.7216 49.3912 29.4233C49.4622 29.125 49.5304 28.7898 49.5957 28.4176C49.661 28.0455 49.7051 27.6676 49.7278 27.2841L49.7576 26.8068H51.1213ZM55.6884 28.1193C55.2708 28.1165 54.8631 28.044 54.4654 27.902C54.0676 27.7571 53.7097 27.5227 53.3915 27.1989C53.0733 26.8722 52.8205 26.4389 52.633 25.8991C52.4455 25.3565 52.3532 24.6847 52.356 23.8835C52.356 23.1364 52.4355 22.4702 52.5946 21.8849C52.7537 21.2997 52.9824 20.8054 53.2807 20.402C53.579 19.9957 53.9384 19.6861 54.3588 19.473C54.7821 19.2599 55.2551 19.1534 55.7779 19.1534C56.3262 19.1534 56.812 19.2614 57.2353 19.4773C57.6614 19.6932 58.0051 19.9886 58.2665 20.3636C58.5279 20.7358 58.6898 21.1562 58.7523 21.625H57.1969C57.1174 21.2898 56.954 21.0227 56.7069 20.8239C56.4625 20.6222 56.1529 20.5213 55.7779 20.5213C55.1728 20.5213 54.7069 20.7841 54.3801 21.3097C54.0563 21.8352 53.8929 22.5568 53.8901 23.4744H53.9498C54.089 23.2244 54.2694 23.0099 54.4909 22.831C54.7125 22.652 54.9625 22.5142 55.2409 22.4176C55.5222 22.3182 55.8191 22.2685 56.1316 22.2685C56.6429 22.2685 57.1017 22.3906 57.508 22.6349C57.9171 22.8793 58.2409 23.2159 58.4796 23.6449C58.7182 24.071 58.8361 24.5597 58.8333 25.1108C58.8361 25.6847 58.7054 26.2003 58.4412 26.6577C58.177 27.1122 57.8091 27.4702 57.3375 27.7315C56.8659 27.9929 56.3162 28.1222 55.6884 28.1193ZM55.6799 26.8409C55.9895 26.8409 56.2665 26.7656 56.5108 26.6151C56.7551 26.4645 56.9483 26.2614 57.0904 26.0057C57.2324 25.75 57.302 25.4631 57.2992 25.1449C57.302 24.8324 57.2338 24.5497 57.0946 24.2969C56.9583 24.044 56.7694 23.8437 56.5279 23.696C56.2864 23.5483 56.0108 23.4744 55.7012 23.4744C55.4711 23.4744 55.2566 23.5185 55.0577 23.6065C54.8588 23.6946 54.6855 23.8168 54.5378 23.973C54.3901 24.1264 54.2736 24.3054 54.1884 24.5099C54.106 24.7116 54.0634 24.9276 54.0605 25.1577C54.0634 25.4616 54.1344 25.7415 54.2736 25.9972C54.4128 26.2528 54.6046 26.4574 54.8489 26.6108C55.0932 26.7642 55.3702 26.8409 55.6799 26.8409ZM60.2736 28V26.858L63.3034 23.8878C63.5932 23.5952 63.8347 23.3352 64.0279 23.108C64.2211 22.8807 64.3659 22.6605 64.4625 22.4474C64.5591 22.2344 64.6074 22.0071 64.6074 21.7656C64.6074 21.4901 64.5449 21.2543 64.4199 21.0582C64.2949 20.8594 64.123 20.706 63.9043 20.598C63.6855 20.4901 63.437 20.4361 63.1586 20.4361C62.8716 20.4361 62.6202 20.4957 62.4043 20.6151C62.1884 20.7315 62.0208 20.8977 61.9015 21.1136C61.785 21.3295 61.7267 21.5866 61.7267 21.8849H60.2225C60.2225 21.331 60.3489 20.8494 60.6017 20.4403C60.8546 20.0312 61.2026 19.7145 61.6458 19.4901C62.0918 19.2656 62.6032 19.1534 63.1799 19.1534C63.7651 19.1534 64.2793 19.2628 64.7225 19.4815C65.1657 19.7003 65.5094 20 65.7537 20.3807C66.0009 20.7614 66.1245 21.196 66.1245 21.6847C66.1245 22.0114 66.062 22.3324 65.937 22.6477C65.812 22.9631 65.5918 23.3125 65.2765 23.696C64.964 24.0795 64.525 24.544 63.9597 25.0895L62.4554 26.6193V26.679H66.2566V28H60.2736ZM70.8588 28.1193C70.2452 28.1193 69.6998 28.0142 69.2225 27.804C68.748 27.5937 68.373 27.3011 68.0975 26.9261C67.8219 26.5511 67.6756 26.1179 67.6586 25.6264H69.2608C69.275 25.8622 69.3532 26.0682 69.4952 26.2443C69.6373 26.4176 69.8262 26.5526 70.062 26.6491C70.2978 26.7457 70.562 26.794 70.8546 26.794C71.1671 26.794 71.4441 26.7401 71.6855 26.6321C71.927 26.5213 72.1159 26.3679 72.2523 26.1719C72.3887 25.9759 72.4554 25.75 72.4526 25.4943C72.4554 25.2301 72.3873 24.9972 72.248 24.7955C72.1088 24.5937 71.9071 24.4361 71.6429 24.3224C71.3816 24.2088 71.0662 24.152 70.6969 24.152H69.9256V22.9332H70.6969C71.0009 22.9332 71.2665 22.8807 71.4938 22.7756C71.7239 22.6705 71.9043 22.5227 72.035 22.3324C72.1657 22.1392 72.2296 21.9162 72.2267 21.6634C72.2296 21.4162 72.1742 21.2017 72.0605 21.0199C71.9498 20.8352 71.7921 20.6918 71.5875 20.5895C71.3858 20.4872 71.1486 20.4361 70.8759 20.4361C70.6088 20.4361 70.3617 20.4844 70.1344 20.581C69.9071 20.6776 69.7239 20.8153 69.5847 20.9943C69.4455 21.1705 69.3716 21.3807 69.3631 21.625H67.8418C67.8532 21.1364 67.9938 20.7074 68.2637 20.3381C68.5364 19.9659 68.9 19.6761 69.3546 19.4688C69.8091 19.2585 70.3191 19.1534 70.8844 19.1534C71.4668 19.1534 71.9725 19.2628 72.4015 19.4815C72.8333 19.6974 73.1671 19.9886 73.4029 20.3551C73.6387 20.7216 73.7566 21.1264 73.7566 21.5696C73.7594 22.0611 73.6145 22.473 73.3219 22.8054C73.0321 23.1378 72.6515 23.3551 72.1799 23.4574V23.5256C72.7935 23.6108 73.2637 23.8381 73.5904 24.2074C73.9199 24.5739 74.0833 25.0298 74.0804 25.5753C74.0804 26.0639 73.9412 26.5014 73.6628 26.8878C73.3873 27.2713 73.0066 27.5724 72.5208 27.7912C72.0378 28.0099 71.4838 28.1193 70.8588 28.1193Z" fill="#C4CCBC"/>`
	svg += `<circle cx="87" cy="24" r="3" fill="#F0B400"/>`
	svg += `<path data-testId="silver" d="M98.1825 19.2727V28H97.1257V20.3807H97.0746L94.9439 21.7955V20.7216L97.1257 19.2727H98.1825Z" fill="#F0B400"/>`
	svg += `<circle cx="137" cy="24" r="3" fill="#AB8A5F"/>`
	svg += `<path data-testId="bronze" d="M146.393 19.2727V28H145.337V20.3807H145.286L143.155 21.7955V20.7216L145.337 19.2727H146.393ZM151.972 19.2727V28H150.915V20.3807H150.864L148.733 21.7955V20.7216L150.915 19.2727H151.972Z" fill="#AB8A5F"/>`
	svg += `<circle cx="112" cy="24" r="3" fill="#999C9F"/>`
	svg += `<path d="M120.248 28L124.152 20.2784V20.2102H119.652V19.2727H125.243V20.2614L121.356 28H120.248Z" fill="#999C9F"/>`
	svg += `</svg>`

	return svg
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func getStackStats(userId string) Stats {
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

	return Stats{
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
