package internal

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"log"
)

type Config struct {
	Listen struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"listen"`
}

var cfg Config

func GetConfig(logger *logrus.Logger) *Config {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalf("can't read config file! Err: %s", err)
	}

	logger.Info("reading configuration file was successful.")

	return &cfg
}
