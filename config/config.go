package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GistToken string `default:""`
	GistURL   string `default:"https://api.github.com/gists"`
	Workspace string `default:".gilbert"`
}

func GetConfig() Config {
	var config Config
	err := envconfig.Process("gilbert", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	config.Workspace = os.Getenv("HOME") + "/" + config.Workspace
	return config
}
