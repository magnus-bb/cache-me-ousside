package main

import (
	"github.com/magnus-bb/cache-me-ousside/cache"
	commandline "github.com/magnus-bb/cache-me-ousside/internal/cli"
	"github.com/magnus-bb/cache-me-ousside/internal/logger"
	"github.com/magnus-bb/cache-me-ousside/internal/router"
)

func main() {
	// Initialize logger in terminal mode to log any startup errors to stdout before a potential log file is provided
	logger.Initialize("") // we want all startup errors etc to be logged to terminal, then we will log to file later if one is provided

	// Get configuration struct from CLI (which might read a config file, if provided)
	conf, err := commandline.CreateConfFromCli()
	if err != nil {
		logger.Fatal(err)
	}

	// Create the actual cache to hold entries
	dataCache, err := cache.New(conf.Capacity, conf.CapacityUnit)
	if err != nil {
		logger.Fatal(err)
	}

	// Setup the router
	app := router.New(conf, dataCache)

	// Say hello in terminal
	logger.HiMom(conf.String(), conf.Address())

	// Set logger to use log file if any is provided
	if conf.LogFilePath != "" {
		logFile := logger.Initialize(conf.LogFilePath)
		if logFile != nil {
			defer logFile.Close()
		}
	}

	// Start the server
	logger.Panic(app.Listen(conf.Address()))
}
