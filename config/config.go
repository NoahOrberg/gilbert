package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GistToken string `default:""`
	GistURL   string `default:""`
}

func GetConfig() Config {
	var config Config
	err := envconfig.Process("gilbert", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
