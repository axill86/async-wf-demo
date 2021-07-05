package config

import "os"

type Config struct {
	EventBusName string
}

func ReadConfig() *Config {
	return &Config{
		EventBusName: os.Getenv("MINICOMPUTE_EVENTBUS_NAME"),
	}
}
