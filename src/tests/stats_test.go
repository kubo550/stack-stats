package routes

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"net/http/httptest"
	"stats/src/middleware"
	"stats/src/routes"
	"stats/src/structs"
	"stats/src/tests/builders"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStatsRoute(t *testing.T) {

	t.Cleanup(func() {
		gock.OffAll()
		gock.EnableNetworking()
	})

	gock.DisableNetworking()

	app := fiber.New()
	app.Use(middleware.StatsMiddleware())
	app.Get("/stats", routes.StatsHandler)

	t.Run("should return 400 when id is missing", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/stats", nil)

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("should return svg content type", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().Build())
		req := httptest.NewRequest("GET", "/stats?id=1", nil)

		resp, _ := app.Test(req)

		assert.Equal(t, "image/svg+xml; charset=utf-8", resp.Header.Get("Content-Type"))
	})

	t.Run("should svg be correctly coded", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().Build())
		req := httptest.NewRequest("GET", "/stats?id=-1", nil)

		resp, _ := app.Test(req)

		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(body), "<svg data-testUserId=\"-1\" width=\"158\" height=\"47\" viewBox=\"0 0 158 47\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">")
	})

	t.Run("should response body contain userId", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().Build())
		userId := "14513625"

		req := httptest.NewRequest("GET", "/stats?id="+userId, nil)
		resp, _ := app.Test(req)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), "data-testUserId=\""+userId+"\"")
	})

	t.Run("should response body contain reputation", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().WithReputation(100).Build())

		req := httptest.NewRequest("GET", "/stats?id=1", nil)
		resp, _ := app.Test(req)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), "data-testReputation=\"100\"")
	})

	t.Run("should response body contain badge", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().WithBadgeCounts(structs.BadgeCounts{
			Gold:   2,
			Silver: 4,
			Bronze: 6,
		}).Build())

		req := httptest.NewRequest("GET", "/stats?id=1", nil)
		resp, _ := app.Test(req)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), "data-testBadgeGold=\"2\"")
		assert.Contains(t, string(body), "data-testBadgeSilver=\"4\"")
		assert.Contains(t, string(body), "data-testBadgeBronze=\"6\"")
	})

	t.Run("should response body contain image URL", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().WithImageUrl("https://www.gravatar.com/avatar/123").Build())

		req := httptest.NewRequest("GET", "/stats?id=1", nil)
		resp, _ := app.Test(req)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), "data-testImageUrl=\"https://www.gravatar.com/avatar/123\"")
	})

	t.Run("should format reputation with comma when it is a more than 1000", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().WithReputation(1500).Build())

		req := httptest.NewRequest("GET", "/stats?id=1", nil)
		resp, _ := app.Test(req)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), "data-testReputation=\"1,500\"")
	})

	t.Run("should format reputation when it is a more than 10 000", func(t *testing.T) {
		builders.StackExchangeWillRespondWith(fiber.StatusOK, builders.NewStackResponseBuilder().WithReputation(26500).Build())

		req := httptest.NewRequest("GET", "/stats?id=1", nil)
		resp, _ := app.Test(req)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), "data-testReputation=\"26.5k\"")
	})

}
