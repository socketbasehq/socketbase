package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		return
	}
}

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func NewConfig() *Config {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	f, err := os.Open(cwd + "/config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	return &cfg
}
