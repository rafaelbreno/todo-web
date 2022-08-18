package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/todo-web/api/storage"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	require := require.New(t)

	app := fiber.New()
	app.Get(HealthCheck(storage.NewLocalMap()))

	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/health", nil))

	require.Nil(err)
	require.Equal(resp.StatusCode, fiber.StatusOK)
}
