package main

import "go.uber.org/zap"

func main() {
	logger, err := zap.NewProduction()
	// Err MUSTN'T be nil.
	if err != nil {
		panic(err)
	}

	undo := zap.ReplaceGlobals(logger)

	defer undo()
}
