package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	IsTest         bool
	UserAddress    string
	UserPrivateKey string
	NetworkUrl     string
	ChainId        uint64
	LineApiKey     string
}

var once sync.Once
var config *Config

func loadConfig() (*Config, error) {
	isTest := false
	mode, err := godotenv.Read("mode.conf")

	if err != nil {
		log.Fatal("Error reading .env file", err)

		return nil, err
	}

	if mode["mode"] != "production" {
		err = godotenv.Load()
		isTest = true
	}

	if err != nil {
		log.Fatal("Error loading .env file", err)

		return nil, err
	}

	chainId, err := strconv.ParseUint(getEnv("CHAIN_ID"), 10, 64)

	if err != nil {
		log.Fatal("Error parsing chain id", err)

		return nil, err
	}

	config = &Config{
		IsTest:         isTest,
		UserAddress:    getEnv("USER_ADDRESS"),
		UserPrivateKey: getEnv("USER_PRIVATE_KEY"),
		NetworkUrl:     getEnv("NETWORK_URL"),
		ChainId:        chainId,
		LineApiKey:     getEnv("LINE_API_KEY"),
	}

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
