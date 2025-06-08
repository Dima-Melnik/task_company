package config

import (
	"os"
	"tasks_company/app/utils/logger"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Port    int    `yaml:"port"`
		User    string `yaml:"user"`
		Host    string `yaml:"host"`
		DBName  string `yaml:"dbname"`
		SSLMode string `yaml:"sslmode"`
	} `yaml:"database"`
}

var Cfg Config

func LoadConfig() {
	file, err := os.ReadFile("config/conf.yaml")
	if err != nil {
		logger.LogWithFieldsFatalf("config", "reading", err)
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		logger.LogWithFieldsFatalf("yaml", "unmarshalling", err)
	}
}
