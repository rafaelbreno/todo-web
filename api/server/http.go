package server

import (
	"encoding/json"
	"fmt"

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

	h.
		setDefaultHandlers()

	handler.SetListHandlers(h.App, st)

	data, _ := json.MarshalIndent(h.App.Stack(), "", "  ")

	fmt.Println(string(data))

	return h
}

func (h *HTTP) setDefaultHandlers() { h.App.Get(handler.HealthCheck()) }
