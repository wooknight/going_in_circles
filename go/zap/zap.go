package main

import (
	"go.uber.org/zap"
)

func main() {
	// Create a sugared logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// Add a key-value pair
	sugar = sugar.With("kanna", "chuka")

	// Write out a few sentences
	sugar.Info("This is an info message.")

	sugar.Infof("Kanna and Chuka are logged together.")
	sugar.Debugw("Debugging values", "additionalKey", 42)
	sugar = sugar.With("kanna", "amulu")
	sugar.Error("This is an error message.")
}
