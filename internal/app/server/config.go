package server

import "tast_for_AvitoAdvertising/store"

type Config struct {
	BindAddr    string        `toml:"bind_addr"`
	LogLevel    string        `toml:"log_level"`
	StoreConfig *store.Config `toml:"store"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		StoreConfig: store.NewConfig(),
		LogLevel:    "debug",
	}
}
