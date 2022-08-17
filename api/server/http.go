package server

import "github.com/gofiber/fiber/v2"

// HTTP - manages the HTTP server.
type HTTP struct {
	App *fiber.App
}

// NewHTTP returns a new instance of HTTP.
func NewHTTP() *HTTP {
	return &HTTP{
		App: fiber.New(),
	}
}
