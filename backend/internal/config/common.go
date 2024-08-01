package config

import "github.com/caarlos0/env/v6"

type ServerConfig struct {
	Port int `env:"PORT" envDefault:"8080"`
}

func LoadServerConfig() ServerConfig {
	var serverConfig ServerConfig
	env.Parse(&serverConfig)
	return serverConfig
}
