package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/rafaelbreno/todo-web/api/storage"
	"github.com/stretchr/testify/require"
)

func TestListRepo(t *testing.T) {
	t.Run("NewListRepo", func(t *testing.T) {
		require := require.New(t)

		actual := NewListRepo(nil)
		expected := ListRepo{
			st: nil,
		}

		require.Equal(expected, *actual)
	})

	t.Run("Create", func(t *testing.T) {
		require := require.New(t)
		repo := NewListRepo(storage.NewLocalMap())

		err := repo.Create(&models.List{
			Name: "Foo",
		})

		require.Nil(err)
	})

	t.Run("Delete", func(t *testing.T) {
		repo := NewListRepo(storage.NewLocalMap())
		t.Run("Not Found", func(t *testing.T) {
			require := require.New(t)

			require.NotNil(repo.Delete(&models.List{
				ID: uuid.New(),
			}))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.List{
				Name: "Foo",
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Delete(&m))

		})
	})

	t.Run("Read", func(t *testing.T) {
		repo := NewListRepo(storage.NewLocalMap())
		t.Run("Not Found", func(t *testing.T) {
			require := require.New(t)

			require.NotNil(repo.Read(&models.List{
				ID: uuid.New(),
			}))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.List{
				Name: "Foo",
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Read(&m))
		})
	})

	t.Run("ReadAll", func(t *testing.T) {
		repo := NewListRepo(storage.NewLocalMap())
		t.Run("Empty", func(t *testing.T) {
			require := require.New(t)

			ms := []models.List{}

			require.Nil(repo.ReadAll(&ms))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.List{
				Name: "Foo",
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Create(&m))

			ms := []models.List{}
			require.Nil(repo.ReadAll(&ms))
			require.Equal(len(ms), 2)
		})
	})

	t.Run("Update", func(t *testing.T) {
		repo := NewListRepo(storage.NewLocalMap())
		t.Run("Not Found", func(t *testing.T) {
			require := require.New(t)

			require.NotNil(repo.Update(&models.List{
				ID: uuid.New(),
			}))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.List{
				Name: "Foo",
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Update(&m))
		})
	})
}
