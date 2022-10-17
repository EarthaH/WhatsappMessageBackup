package config

import (
	"github.com/tkanos/gonfig"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func GetConfig() Config {
	config := Config{}

	filename := "./dev_config.json"
	gonfig.GetConf(filename, &config)

	return config
}
