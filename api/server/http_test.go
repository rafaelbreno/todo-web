package server

import (
	"testing"

	"github.com/rafaelbreno/todo-web/api/storage"
	"github.com/stretchr/testify/require"
)

func TestHTTP(t *testing.T) {
	srv := NewHTTP(storage.NewLocalMap())

	t.Run("Check Config", func(t *testing.T) {
		require := require.New(t)

		require.True(srv.App.Config().StrictRouting)
	})

	t.Run("Check Stack", func(t *testing.T) {
		require := require.New(t)

		srvStack := srv.App.Stack()

		require.Greater(len(srvStack), 0)
	})
}
