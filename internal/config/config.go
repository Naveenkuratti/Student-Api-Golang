package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"addr" env:"ADDR" env-default:":8080"`
}

type Config struct {
	ENV         string      `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string      `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	// Try to get path from environment variable
	configPath = os.Getenv("CONFIG_PATH")

	// If not found, check command-line flag
	if configPath == "" {
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	// Load config
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read config: %s", err.Error())
	}

	return &cfg
}
