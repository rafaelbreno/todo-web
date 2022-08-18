package models

import (
	"testing"

	"github.com/rafaelbreno/todo-web/api/enum"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("SetDefault", func(t *testing.T) {
		require := require.New(t)

		actual := List{}
		actual.SetDefault()

		expected := List{
			Status: enum.ListNotStarted,
		}

		require.Equal(expected, actual)
	})

	t.Run("Validate", func(t *testing.T) {
		t.Run("Empty Name", func(t *testing.T) {
			require := require.New(t)

			list := List{}

			require.NotNil(list.Validate())
		})

		t.Run("Valid", func(t *testing.T) {
			require := require.New(t)

			list := List{
				Name: "foo",
			}

			require.Nil(list.Validate())
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Name", func(t *testing.T) {
			require := require.New(t)

			list := List{}
			newList := List{
				Name: "foo",
			}

			list.Update(newList)

			require.Equal(list, newList)
		})
		t.Run("Status", func(t *testing.T) {
			require := require.New(t)

			list := List{}
			newList := List{
				Name:   "foo",
				Status: enum.ListCompleted,
			}

			list.Update(newList)

			require.Equal(list, newList)
		})
	})
}
