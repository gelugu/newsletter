package main

import (
	"flag"
)

type Arguments struct {
	ConfigPath string
	Debug      bool
}

func parseArgs() Arguments {
	args := Arguments{}

	flag.StringVar(&args.ConfigPath, "config", "/etc/newsletter/config.yaml", "Path to the configuration file")
	flag.BoolVar(&args.Debug, "debug", false, "Enable debug mode")

	flag.Parse()

	return args
}

var args = parseArgs()
