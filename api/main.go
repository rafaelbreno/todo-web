package main

import (
	"os"
	"os/signal"

	"github.com/rafaelbreno/todo-web/api/server"
	"github.com/rafaelbreno/todo-web/api/storage"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	// Err MUSTN'T be nil.
	if err != nil {
		panic(err)
	}

	undo := zap.ReplaceGlobals(logger)

	defer undo()

	st := storage.NewLocalMap()

	srv := server.NewHTTP(st)

	go srv.App.Listen(":5001")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	zap.L().Info("gracefully shutdown...")
	if err := srv.App.Shutdown(); err != nil {
		zap.Error(err)
	}
}
