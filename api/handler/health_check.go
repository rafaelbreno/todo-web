package handler

import "github.com/gofiber/fiber/v2"

// HealthCheck returns a Fiber Handler to check the
// 	health of the APP.
// TODO:
// 	- Implement Storage Health Check
func HealthCheck() (string, func(c *fiber.Ctx) error) {
	return "/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "healthy",
		})
	}
}