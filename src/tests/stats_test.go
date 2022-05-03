package routes

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"stats/src/middleware"
	"stats/src/routes"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStatsRoute(t *testing.T) {

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
		req := httptest.NewRequest("GET", "/stats?id=1", nil)

		resp, _ := app.Test(req)

		assert.Equal(t, "image/svg+xml; charset=utf-8", resp.Header.Get("Content-Type"))
	})

	t.Run("should svg be correctly coded", func(t *testing.T) {
		userId := "14513625"
		req := httptest.NewRequest("GET", "/stats?id="+userId, nil)

		resp, _ := app.Test(req)

		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(body), "<svg width=\"158\" height=\"47\" viewBox=\"0 0 158 47\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">")
	})

	t.Run("should response body contain userId", func(t *testing.T) {
		userId := "14513625"
		req := httptest.NewRequest("GET", "/stats?id="+userId, nil)

		resp, _ := app.Test(req)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Contains(t, string(body), "data-testUserId=\""+userId+"\"")
	})

	//	todo: mock stackoverflow api
}
