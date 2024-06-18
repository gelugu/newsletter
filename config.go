package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" default:"info"`

	TelegramToken     string `yaml:"telegram_token" env:"TELEGRAM_TOKEN"`
	TelegramChannelID int64  `yaml:"telegram_channel_id" env:"TELEGRAM_CHANNEL_ID"`

	GitLabToken   string `yaml:"gitlab_token"`
	GitLabGroupID int    `yaml:"gitlab_group"`
}

func readConfig() Config {
	path := args.ConfigPath

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s", err.Error())
		os.Exit(1)
	}

	cfg := Config{}
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		fmt.Printf("Error unmarshalling yaml: %s", err.Error())
		os.Exit(1)
	}

	return cfg
}

var config = readConfig()
