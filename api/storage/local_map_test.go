package storage

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/stretchr/testify/require"
)

func TestLocalMapStorage(t *testing.T) {
	t.Run("CheckInterface", func(_ *testing.T) {
		var _ Storage = NewLocalMap()
	})

	t.Run("NewLocalMap", func(t *testing.T) {
		require := require.New(t)

		expected := LocalMap{
			lists: make(map[string]models.List),
			items: make(map[string]models.Item),
		}

		actual := NewLocalMap()

		require.Equal(expected, *actual)
	})

	t.Run("Create", func(t *testing.T) {
		l := NewLocalMap()

		t.Run("Unsupported Type", func(t *testing.T) {
			require := require.New(t)

			expected := fmt.Errorf("type '%T' not supported", "foo")

			actual := l.Create("bar")

			require.Equal(expected, actual)
		})

		t.Run("List", func(t *testing.T) {
			require := require.New(t)

			actual := l.Create(models.List{
				ID: uuid.New(),
			})

			require.Nil(actual)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		l := NewLocalMap()

		t.Run("Unsupported Type", func(t *testing.T) {
			require := require.New(t)

			expected := fmt.Errorf("type '%T' not supported", "foo")

			actual := l.Delete("bar")

			require.Equal(expected, actual)
		})

		t.Run("Unexistent List", func(t *testing.T) {
			require := require.New(t)

			actual := l.Delete(models.List{
				ID: uuid.New(),
			})

			require.NotNil(actual)
		})

		t.Run("List", func(t *testing.T) {
			require := require.New(t)

			m := models.List{
				ID: uuid.New(),
			}

			require.Nil(l.Create(m))

			actual := l.Delete(m)

			if actual != nil {
				t.Fatal(actual)
			}

			require.Nil(actual)
		})
	})

	t.Run("Read", func(t *testing.T) {
		l := NewLocalMap()

		t.Run("Unsupported Type", func(t *testing.T) {
			require := require.New(t)

			expected := fmt.Errorf("type '%T' not supported", "foo")

			actual := l.Read("bar")

			require.Equal(expected, actual)
		})

		t.Run("Unexistent List", func(t *testing.T) {
			require := require.New(t)

			actual := l.Read(models.List{
				ID: uuid.New(),
			})

			require.NotNil(actual)
		})

		t.Run("List", func(t *testing.T) {
			require := require.New(t)

			m1 := models.List{
				ID:   uuid.New(),
				Name: "foo",
			}

			require.Nil(l.Create(m1))

			m2 := models.List{
				ID: m1.ID,
			}

			require.Nil(l.Read(&m2))

			require.Equal(m1, m2)
		})
	})

	t.Run("ReadAll", func(t *testing.T) {
		l := NewLocalMap()

		t.Run("Unsupported Type", func(t *testing.T) {
			require := require.New(t)

			expected := fmt.Errorf("type '%T' not supported", "foo")

			actual := l.ReadAll("bar")

			require.Equal(expected, actual)
		})

		t.Run("List", func(t *testing.T) {
			require := require.New(t)

			m1 := models.List{
				ID:   uuid.New(),
				Name: "foo",
			}

			require.Nil(l.Create(m1))

			listArr := []models.List{}

			require.Nil(l.ReadAll(&listArr))

			require.Equal(len(listArr), 1)
		})

		t.Run("List", func(t *testing.T) {
			require := require.New(t)

			m1 := models.List{
				ID:   uuid.New(),
				Name: "foo",
			}

			require.Nil(l.Create(m1))

			listArr := []models.List{}

			require.Nil(l.ReadAll(&listArr))

			require.Equal(len(listArr), 2)
		})
	})

	t.Run("Update", func(t *testing.T) {
		l := NewLocalMap()

		t.Run("Unsupported Type", func(t *testing.T) {
			require := require.New(t)

			expected := fmt.Errorf("type '%T' not supported", "foo")

			actual := l.Update("bar")

			require.Equal(expected, actual)
		})

		t.Run("Unexistent List", func(t *testing.T) {
			require := require.New(t)

			actual := l.Update(models.List{
				ID: uuid.New(),
			})

			require.NotNil(actual)
		})

		t.Run("List", func(t *testing.T) {
			require := require.New(t)

			m := models.List{
				ID: uuid.New(),
			}

			require.Nil(l.Create(m))

			actual := l.Update(m)

			require.Nil(actual)
		})
	})
}
