package config

import (
	"os"
)

type Config struct {
	AppPort      string
	Environment  string
	DataBaseUser string
	DataBasePass string
	DataBaseHost string
	DataBaseName string
}

func (conf *Config) init() {
	conf.AppPort = os.Getenv("APP_PORT")
	conf.Environment = os.Getenv("ENVIRONMENT")
	conf.DataBaseUser = os.Getenv("DB_USER")
	conf.DataBasePass = os.Getenv("DB_PASS")
	conf.DataBaseHost = os.Getenv("DB_URL")
	conf.DataBaseName = os.Getenv("DB_NAME")
}

func (conf *Config) DataBaseConnectionUrl() string {
	//TODO: investigate connection url
	return ""
}

func NewConfig() *Config {
	config := new(Config)
	config.init()
	return config
}
