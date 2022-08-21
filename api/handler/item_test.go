package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/rafaelbreno/todo-web/api/pkg/helpers"
	"github.com/rafaelbreno/todo-web/api/storage"
	"github.com/stretchr/testify/require"
)

func TestCreateItem(t *testing.T) {
	app := fiber.New()
	st := storage.NewLocalMap()
	SetItemHandlers(app, st)
	SetListHandlers(app, st)
	list := new(models.List)

	t.Run("Create List", func(t *testing.T) {
		require := require.New(t)

		list.Name = "name"

		listBytes, err := json.Marshal(list)
		require.Nil(err)

		body := bytes.NewReader(listBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/list", body))
		defer func() {
			require.Nil(resp.Body.Close())
		}()

		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
		bodyBytes, err := io.ReadAll(resp.Body)

		require.Nil(err)
		require.Nil(json.Unmarshal(bodyBytes, list))
	})

	t.Run("Create Item with ListID", func(t *testing.T) {
		require := require.New(t)

		item := &models.Item{
			Text:   "name",
			ListID: list.ID,
		}

		itemBytes, err := json.Marshal(item)
		require.Nil(err)

		body := bytes.NewReader(itemBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/item", body))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})
}

func TestDeleteItem(t *testing.T) {
	app := fiber.New()
	st := storage.NewLocalMap()
	SetItemHandlers(app, st)
	SetListHandlers(app, st)
	list := new(models.List)

	t.Run("Create List", func(t *testing.T) {
		require := require.New(t)

		list.Name = "name"

		listBytes, err := json.Marshal(list)
		require.Nil(err)

		body := bytes.NewReader(listBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/list", body))
		defer func() {
			require.Nil(resp.Body.Close())
		}()

		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
		bodyBytes, err := io.ReadAll(resp.Body)

		require.Nil(err)
		require.Nil(json.Unmarshal(bodyBytes, list))
	})

	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodDelete, "/item/222", nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		randomID := uuid.NewString()

		resp, err := app.Test(helpers.NewTestRequest(http.MethodDelete, "/item/"+randomID, nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("Valid Request", func(t *testing.T) {
		returnedItem := new(models.Item)
		t.Run("Create", func(t *testing.T) {
			require := require.New(t)

			item := models.Item{
				Text:   "name",
				ListID: list.ID,
			}

			itemBytes, err := json.Marshal(item)
			require.Nil(err)

			body := bytes.NewReader(itemBytes)

			resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/item", body))
			defer func() {
				require.Nil(resp.Body.Close())
			}()
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusOK)

			bodyBytes, err := io.ReadAll(resp.Body)
			require.Nil(err)
			require.Nil(json.Unmarshal(bodyBytes, returnedItem))
		})
		require := require.New(t)
		resp, err := app.Test(helpers.NewTestRequest(http.MethodDelete, "/item/"+returnedItem.ID.String(), nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})
}

func TestReadItem(t *testing.T) {
	app := fiber.New()
	st := storage.NewLocalMap()
	SetItemHandlers(app, st)
	SetListHandlers(app, st)
	list := new(models.List)

	t.Run("Create List", func(t *testing.T) {
		require := require.New(t)

		list.Name = "name"

		listBytes, err := json.Marshal(list)
		require.Nil(err)

		body := bytes.NewReader(listBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/list", body))
		defer func() {
			require.Nil(resp.Body.Close())
		}()

		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
		bodyBytes, err := io.ReadAll(resp.Body)

		require.Nil(err)
		require.Nil(json.Unmarshal(bodyBytes, list))
	})

	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/item/222", nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		randomID := uuid.NewString()

		resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/item/"+randomID, nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})

	t.Run("Valid Request", func(t *testing.T) {
		returnedItem := new(models.Item)
		t.Run("Create", func(t *testing.T) {
			require := require.New(t)

			item := models.Item{
				Text:   "name",
				ListID: list.ID,
			}

			itemBytes, err := json.Marshal(item)
			require.Nil(err)

			body := bytes.NewReader(itemBytes)

			resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/item", body))
			defer func() {
				require.Nil(resp.Body.Close())
			}()
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusOK)

			bodyBytes, err := io.ReadAll(resp.Body)
			require.Nil(err)
			require.Nil(json.Unmarshal(bodyBytes, returnedItem))
		})
		require := require.New(t)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/item/"+returnedItem.ID.String(), nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})
}

func TestReadAllItem(t *testing.T) {
	app := fiber.New()
	st := storage.NewLocalMap()
	SetItemHandlers(app, st)
	SetListHandlers(app, st)
	list := new(models.List)

	t.Run("Create List", func(t *testing.T) {
		require := require.New(t)

		list.Name = "name"

		listBytes, err := json.Marshal(list)
		require.Nil(err)

		body := bytes.NewReader(listBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/list", body))
		defer func() {
			require.Nil(resp.Body.Close())
		}()

		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
		bodyBytes, err := io.ReadAll(resp.Body)

		require.Nil(err)
		require.Nil(json.Unmarshal(bodyBytes, list))
	})

	t.Run("Create", func(t *testing.T) {
		require := require.New(t)

		item := models.Item{
			Text:   "name",
			ListID: list.ID,
		}

		itemBytes, err := json.Marshal(item)
		require.Nil(err)

		body := bytes.NewReader(itemBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/item", body))
		defer func() {
			require.Nil(resp.Body.Close())
		}()
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})

	resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/item", nil))

	t.Run("Check Response Values", func(_ *testing.T) {
		require := require.New(t)
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})

	t.Run("Check Response body", func(_ *testing.T) {
		require := require.New(t)
		bodyBytes, err := io.ReadAll(resp.Body)
		require.Nil(err)

		ms := new([]models.Item)
		json.Unmarshal(bodyBytes, ms)
		require.Equal(1, len(*ms))
	})
}

func TestUpdateItem(t *testing.T) {
	app := fiber.New()
	st := storage.NewLocalMap()
	SetItemHandlers(app, st)
	SetListHandlers(app, st)
	list := new(models.List)

	t.Run("Create List", func(t *testing.T) {
		require := require.New(t)

		list.Name = "name"

		listBytes, err := json.Marshal(list)
		require.Nil(err)

		body := bytes.NewReader(listBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/list", body))
		defer func() {
			require.Nil(resp.Body.Close())
		}()

		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
		bodyBytes, err := io.ReadAll(resp.Body)

		require.Nil(err)
		require.Nil(json.Unmarshal(bodyBytes, list))
	})

	t.Run("Bad Request", func(t *testing.T) {
		t.Run("Method Put", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPut, "/item/222", nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})
		t.Run("Method Patch", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPatch, "/item/222", nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})

	})

	t.Run("Bad Request", func(t *testing.T) {
		randomID := uuid.NewString()
		t.Run("Method Put", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPut, "/item/"+randomID, nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})
		t.Run("Method Patch", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPatch, "/item/"+randomID, nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})
	})

	t.Run("Valid Request", func(t *testing.T) {
		returnedItem := new(models.Item)
		t.Run("Create", func(t *testing.T) {
			require := require.New(t)

			item := models.Item{
				Text:   "name",
				ListID: list.ID,
			}

			itemBytes, err := json.Marshal(item)
			require.Nil(err)

			body := bytes.NewReader(itemBytes)

			resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/item", body))
			defer func() {
				require.Nil(resp.Body.Close())
			}()
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusOK)

			bodyBytes, err := io.ReadAll(resp.Body)
			require.Nil(err)
			require.Nil(json.Unmarshal(bodyBytes, returnedItem))
		})

		newItem := models.Item{
			Text: "bar",
		}
		itemBytes, err := json.Marshal(newItem)
		require.Nil(t, err)

		t.Run("Method Put", func(t *testing.T) {
			require := require.New(t)
			body := bytes.NewReader(itemBytes)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPut, "/item/"+returnedItem.ID.String(), body))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusOK)
		})

		t.Run("Method Patch", func(t *testing.T) {
			require := require.New(t)
			body := bytes.NewReader(itemBytes)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPatch, "/item/"+returnedItem.ID.String(), body))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusOK)
		})
	})
}
