package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func getLogger() *logrus.Logger {
	var logger = logrus.New()
	logger.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		fmt.Printf("Invalid log level: %s\n", config.LogLevel)
		os.Exit(1)
	}

	if args.Debug {
		logLevel = logrus.DebugLevel
	}

	logger.SetLevel(logLevel)

	return logger
}

var log = getLogger()
