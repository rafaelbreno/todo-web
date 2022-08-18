package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/rafaelbreno/todo-web/api/repository"
	"github.com/rafaelbreno/todo-web/api/storage"
)

// ListHandler - Stores all Handlers related to List model.
type ListHandler struct {
	repo  repository.Repository[models.List]
	route fiber.Router
}

// SetListHandlers receives an instance of fiber.App and sets all handlers
// 	related to list model.
func SetListHandlers(app *fiber.App, st storage.Storage) {
	l := &ListHandler{
		repo:  repository.NewListRepo(st),
		route: app.Group("/list"),
	}

	l.route.Post(l.Create())
	l.route.Delete(l.Delete())
	l.route.Get(l.ReadAll())
	l.route.Get(l.Read())
	l.route.Put(l.Update())
	l.route.Patch(l.Update())
}

// Create - Handler to create a List.
func (l *ListHandler) Create() (string, func(*fiber.Ctx) error) {
	return "", func(c *fiber.Ctx) error {
		m := models.List{}
		if err := c.BodyParser(&m); err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		// TODO:
		// 	- Validate `m`

		if err := l.repo.Create(&m); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		return c.
			Status(fiber.StatusOK).
			JSON(fiber.Map{
				"list": m,
			})
	}
}

// Delete - Handler to delete a List.
func (l *ListHandler) Delete() (string, func(*fiber.Ctx) error) {
	return "/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		m := models.List{
			ID: id,
		}

		if err := l.repo.Delete(&m); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		return c.
			Status(fiber.StatusOK).
			JSON(fiber.Map{
				"list": m,
			})
	}
}

// Read - Handler to retrieve a List with given ID.
func (l *ListHandler) Read() (string, func(*fiber.Ctx) error) {
	return "/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		m := models.List{
			ID: id,
		}

		if err := l.repo.Read(&m); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		return c.
			Status(fiber.StatusOK).
			JSON(fiber.Map{
				"list": m,
			})
	}
}

// ReadAll - Handler to retrieve all lists.
func (l *ListHandler) ReadAll() (string, func(*fiber.Ctx) error) {
	return "", func(c *fiber.Ctx) error {
		m := []models.List{}

		if err := l.repo.ReadAll(&m); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		return c.
			Status(fiber.StatusOK).
			JSON(fiber.Map{
				"lists": m,
			})
	}
}

// Update - Handler to update a List with given ID and body.
func (l *ListHandler) Update() (string, func(*fiber.Ctx) error) {
	return "/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		m := models.List{}

		if err := c.BodyParser(&m); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		m.ID = id

		if err := l.repo.Update(&m); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		return c.
			Status(fiber.StatusOK).
			JSON(fiber.Map{
				"list": m,
			})
	}
}
