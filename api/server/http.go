package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/todo-web/api/handler"
)

// HTTP - manages the HTTP server.
type HTTP struct {
	App *fiber.App
}

// NewHTTP returns a new instance of HTTP.
func NewHTTP() *HTTP {
	h := &HTTP{
		App: fiber.New(),
	}

	h.setHandlers()

	return h
}

func (h *HTTP) setHandlers() {
	h.App.Get("/health", handler.HealthCheck())
}
