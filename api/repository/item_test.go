package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/rafaelbreno/todo-web/api/storage"
	"github.com/stretchr/testify/require"
)

func TestItemRepo(t *testing.T) {
	st := storage.NewLocalMap()
	list := &models.List{
		Name: "List",
	}
	_ = NewListRepo(st).Create(list)

	t.Run("NewItemRepo", func(t *testing.T) {
		require := require.New(t)

		actual := NewItemRepo(nil)
		expected := ItemRepo{
			st: nil,
		}

		require.Equal(expected, *(actual.(*ItemRepo)))
	})

	t.Run("Create", func(t *testing.T) {
		require := require.New(t)
		repo := NewItemRepo(st)

		err := repo.Create(&models.Item{
			Text:   "Foo",
			ListID: list.ID,
		})

		require.Nil(err)
	})

	t.Run("Delete", func(t *testing.T) {
		repo := NewItemRepo(st)
		t.Run("Not Found", func(t *testing.T) {
			require := require.New(t)

			require.NotNil(repo.Delete(&models.Item{
				ID:     uuid.New(),
				ListID: list.ID,
			}))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.Item{
				Text:   "Foo",
				ListID: list.ID,
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Delete(&m))

		})
	})

	t.Run("Read", func(t *testing.T) {
		repo := NewItemRepo(st)
		t.Run("Not Found", func(t *testing.T) {
			require := require.New(t)

			require.NotNil(repo.Read(&models.Item{
				ID:     uuid.New(),
				ListID: list.ID,
			}))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.Item{
				Text:   "Foo",
				ListID: list.ID,
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Read(&m))
		})
	})

	t.Run("ReadAll", func(t *testing.T) {
		repo := NewItemRepo(st)
		t.Run("Empty", func(t *testing.T) {
			require := require.New(t)

			ms := []models.Item{}

			require.Nil(repo.ReadAll(&ms))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.Item{
				Text:   "Foo",
				ListID: list.ID,
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Create(&m))

			ms := []models.Item{}
			require.Nil(repo.ReadAll(&ms))
			require.Equal(len(ms), 4)
		})
	})

	t.Run("Update", func(t *testing.T) {
		repo := NewItemRepo(st)
		t.Run("Not Found", func(t *testing.T) {
			require := require.New(t)

			require.NotNil(repo.Update(&models.Item{
				ID:     uuid.New(),
				ListID: list.ID,
			}))
		})

		t.Run("Sucessful", func(t *testing.T) {
			require := require.New(t)

			m := models.Item{
				Text:   "Foo",
				ListID: list.ID,
			}
			require.Nil(repo.Create(&m))
			require.Nil(repo.Update(&m))
		})
	})
}
