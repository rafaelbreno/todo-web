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

func TestCreateList(t *testing.T) {
	app := fiber.New()
	SetListHandlers(app, storage.NewLocalMap())

	require := require.New(t)

	list := models.List{
		Name: "name",
	}

	listBytes, err := json.Marshal(list)
	require.Nil(err)

	body := bytes.NewReader(listBytes)

	resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/list", body))
	require.Nil(err)
	require.Equal(resp.StatusCode, fiber.StatusOK)
}

func TestDelete(t *testing.T) {
	app := fiber.New()
	SetListHandlers(app, storage.NewLocalMap())

	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodDelete, "/list/222", nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		randomID := uuid.NewString()

		resp, err := app.Test(helpers.NewTestRequest(http.MethodDelete, "/list/"+randomID, nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("Valid Request", func(t *testing.T) {
		returnedList := new(models.List)
		t.Run("Create", func(t *testing.T) {
			require := require.New(t)

			list := models.List{
				Name: "name",
			}

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
			require.Nil(json.Unmarshal(bodyBytes, returnedList))
		})
		require := require.New(t)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodDelete, "/list/"+returnedList.ID.String(), nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})
}

func TestReadList(t *testing.T) {
	app := fiber.New()
	SetListHandlers(app, storage.NewLocalMap())

	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/list/222", nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("Bad Request", func(t *testing.T) {
		require := require.New(t)

		randomID := uuid.NewString()

		resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/list/"+randomID, nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("Valid Request", func(t *testing.T) {
		returnedList := new(models.List)
		t.Run("Create", func(t *testing.T) {
			require := require.New(t)

			list := models.List{
				Name: "name",
			}

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
			require.Nil(json.Unmarshal(bodyBytes, returnedList))
		})
		require := require.New(t)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/list/"+returnedList.ID.String(), nil))
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})
}

func TestReadAllList(t *testing.T) {
	app := fiber.New()
	SetListHandlers(app, storage.NewLocalMap())

	t.Run("Create", func(t *testing.T) {
		require := require.New(t)

		list := models.List{
			Name: "name",
		}

		listBytes, err := json.Marshal(list)
		require.Nil(err)

		body := bytes.NewReader(listBytes)

		resp, err := app.Test(helpers.NewTestRequest(http.MethodPost, "/list", body))
		defer func() {
			require.Nil(resp.Body.Close())
		}()
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})

	resp, err := app.Test(helpers.NewTestRequest(http.MethodGet, "/list", nil))

	t.Run("Check Response Values", func(_ *testing.T) {
		require := require.New(t)
		require.Nil(err)
		require.Equal(resp.StatusCode, fiber.StatusOK)
	})

	t.Run("Check Response body", func(_ *testing.T) {
		require := require.New(t)
		bodyBytes, err := io.ReadAll(resp.Body)
		require.Nil(err)

		ms := new([]models.List)
		json.Unmarshal(bodyBytes, ms)
		require.Equal(1, len(*ms))
	})
}

func TestUpdateList(t *testing.T) {
	app := fiber.New()
	SetListHandlers(app, storage.NewLocalMap())

	t.Run("Bad Request", func(t *testing.T) {
		t.Run("Method Put", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPut, "/list/222", nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})
		t.Run("Method Patch", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPatch, "/list/222", nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})

	})

	t.Run("Bad Request", func(t *testing.T) {
		randomID := uuid.NewString()
		t.Run("Method Put", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPut, "/list/"+randomID, nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})
		t.Run("Method Patch", func(t *testing.T) {
			require := require.New(t)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPatch, "/list/"+randomID, nil))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusBadRequest)
		})
	})

	t.Run("Valid Request", func(t *testing.T) {
		returnedList := new(models.List)
		t.Run("Create", func(t *testing.T) {
			require := require.New(t)

			list := models.List{
				Name: "name",
			}

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
			require.Nil(json.Unmarshal(bodyBytes, returnedList))
		})

		newList := models.List{
			Name: "bar",
		}
		listBytes, err := json.Marshal(newList)
		require.Nil(t, err)

		t.Run("Method Put", func(t *testing.T) {
			require := require.New(t)
			body := bytes.NewReader(listBytes)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPut, "/list/"+returnedList.ID.String(), body))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusOK)
		})

		t.Run("Method Patch", func(t *testing.T) {
			require := require.New(t)
			body := bytes.NewReader(listBytes)
			resp, err := app.Test(helpers.NewTestRequest(http.MethodPatch, "/list/"+returnedList.ID.String(), body))
			require.Nil(err)
			require.Equal(resp.StatusCode, fiber.StatusOK)
		})
	})
}
