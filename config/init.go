package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserBEP20Address string
}

var config Config

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	config.UserBEP20Address = getEnv("USER_BEP20_ADDRESS")
}

func getEnv(env string) string {
	const prefix = "AUTO_COMPOUND_"

	return os.Getenv(prefix + env)
}

func GetConfig() Config {
	return config
}
