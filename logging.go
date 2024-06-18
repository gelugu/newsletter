package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func getLogger() *logrus.Logger {
	var log = logrus.New()
	log.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		fmt.Printf("Invalid log level: %s\n", config.LogLevel)
		os.Exit(1)
	}

	if args.Debug {
		logLevel = logrus.DebugLevel
	}

	log.SetLevel(logLevel)

	return log
}

var log = getLogger()
