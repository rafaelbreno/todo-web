package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/todo-web/api/storage"
)

// HealthCheck returns a Fiber Handler to check the
// 	health of the APP.
// TODO:
// 	- Implement Storage Health Check
func HealthCheck(st storage.Storage) (string, func(c *fiber.Ctx) error) {
	return "/health", func(c *fiber.Ctx) error {
		if err := st.HealthCheck(); err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"message": err.Error(),
				})
		}
		return c.JSON(fiber.Map{
			"message": "healthy",
		})
	}
}
