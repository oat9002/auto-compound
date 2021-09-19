package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	UserAddress   string
	BscNetworkUrl string
	LineApiKey    string
}

var once sync.Once
var config *Config

func loadConfig() (*Config, error) {
	mode, err := godotenv.Read("mode.conf")

	if err != nil {
		log.Fatal("Error reading .env file", err)

		return config, err
	}

	if mode["mode"] != "production" {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error loading .env file", err)

		return config, err
	}

	config = &Config{UserAddress: getEnv("USER_ADDRESS"), BscNetworkUrl: getEnv("BSC_NETWORK_URL"), LineApiKey: getEnv("LINE_API_KEY")}

	return config, nil
}

func getEnv(env string) string {
	const prefix = "AUTO_COMPOUND_"

	return os.Getenv(prefix + env)
}

func GetConfig() (Config, error) {
	if config != nil {
		return *config, nil
	}

	var err error

	once.Do(func() {
		config, err = loadConfig()
	})

	return *config, err
}
