package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/todo-web/api/handler"
	"github.com/rafaelbreno/todo-web/api/storage"
)

// HTTP - manages the HTTP server.
type HTTP struct {
	App *fiber.App
}

// NewHTTP returns a new instance of HTTP.
func NewHTTP(st storage.Storage) *HTTP {
	h := &HTTP{
		App: fiber.New(fiber.Config{
			StrictRouting: true,
		}),
	}

	h.App.Get(handler.HealthCheck(st))
	handler.SetListHandlers(h.App, st)
	handler.SetItemHandlers(h.App, st)

	return h
}
