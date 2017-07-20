package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token string
}

func GetConfig() Config {
	var config Config
	err := envconfig.Process("gist", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
