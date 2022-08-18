package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/enum"
	"github.com/stretchr/testify/require"
)

func TestItem(t *testing.T) {
	t.Run("SetDefault", func(t *testing.T) {
		require := require.New(t)

		actual := Item{}
		actual.SetDefault()

		expected := Item{
			Status: enum.ItemNotStarted,
		}

		require.Equal(expected, actual)
	})

	t.Run("Validate", func(t *testing.T) {
		t.Run("Empty ListID", func(t *testing.T) {
			require := require.New(t)

			item := Item{
				Text: "foo",
			}

			require.NotNil(item.Validate())
		})

		t.Run("Empty Text", func(t *testing.T) {
			require := require.New(t)

			item := Item{
				ListID: uuid.New(),
			}

			require.NotNil(item.Validate())
		})

		t.Run("Valid", func(t *testing.T) {
			require := require.New(t)

			item := Item{
				ListID: uuid.New(),
				Text:   "foo",
			}

			require.Nil(item.Validate())
		})
	})
}
