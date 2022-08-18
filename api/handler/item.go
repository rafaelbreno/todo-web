package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/rafaelbreno/todo-web/api/repository"
	"github.com/rafaelbreno/todo-web/api/storage"
)

// ItemHandler - Stores all Handlers related to Item model.
type ItemHandler struct {
	repo     repository.Repository[models.Item]
	listRepo repository.Repository[models.List]
	route    fiber.Router
}

// SetItemHandlers receives an instance of fiber.App and sets all handlers
// 	related to item model.
func SetItemHandlers(app *fiber.App, st storage.Storage) {
	l := &ItemHandler{
		repo:     repository.NewItemRepo(st),
		listRepo: repository.NewListRepo(st),
		route:    app.Group("/item"),
	}

	l.route.Post(l.Create())
	l.route.Delete(l.Delete())
	l.route.Get(l.ReadAll())
	l.route.Get(l.Read())
	l.route.Put(l.Update())
	l.route.Patch(l.Update())
}

// Create - Handler to create a Item.
func (l *ItemHandler) Create() (string, func(*fiber.Ctx) error) {
	return "", func(c *fiber.Ctx) error {
		m := models.Item{}
		if err := c.BodyParser(&m); err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

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
				"item": m,
			})
	}
}

// Delete - Handler to delete a Item.
func (l *ItemHandler) Delete() (string, func(*fiber.Ctx) error) {
	return "/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		m := models.Item{
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
				"item": m,
			})
	}
}

// Read - Handler to retrieve a Item with given ID.
func (l *ItemHandler) Read() (string, func(*fiber.Ctx) error) {
	return "/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		m := models.Item{
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
				"item": m,
			})
	}
}

// ReadAll - Handler to retrieve all items.
func (l *ItemHandler) ReadAll() (string, func(*fiber.Ctx) error) {
	return "", func(c *fiber.Ctx) error {
		m := []models.Item{}

		b := map[string]string{}

		if err := c.BodyParser(b); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		if listID, ok := b["list_id"]; ok {
			id, err := uuid.Parse(listID)
			if err != nil {
				return c.
					Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{
						"error": err.Error(),
					})
			}
			m = append(m, models.Item{
				ListID: id,
			})
		}

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
				"items": m,
			})
	}
}

// Update - Handler to update a Item with given ID and body.
func (l *ItemHandler) Update() (string, func(*fiber.Ctx) error) {
	return "/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		m := models.Item{}

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
				"item": m,
			})
	}
}
