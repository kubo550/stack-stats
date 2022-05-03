package routes

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"stats/src/routes"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthRoute(t *testing.T) {

	tests := []struct {
		name         string
		expectedCode int
	}{
		{
			name:         "Should return 200",
			expectedCode: 200,
		},
	}

	app := fiber.New()
	app.Get("/health", routes.HealthHandler)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/health", nil)

			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.name)
		})
	}
}
