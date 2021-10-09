package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const bscNetwork string = "https://bsc-dataseed.binance.org/"
const bscChainId int = 56
const bscTestNetwork string = "https://data-seed-prebsc-1-s1.binance.org:8545/"
const bacTestChainId int = 97

type Config struct {
	IsDevelopment   bool
	UseTestNetwork  bool
	OnlyCheckReward bool
	UserAddress     string
	UserPrivateKey  string
	LineApiKey      string
}

var once sync.Once
var config *Config

func loadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)

		return nil, err
	}

	config = &Config{
		IsDevelopment:   getEnv("MODE") == "development",
		UseTestNetwork:  getEnv("USE_TEST_NETWORK") == "true",
		OnlyCheckReward: getEnv("ONLY_CHECK_REWARD") == "true",
		UserAddress:     getEnv("USER_ADDRESS"),
		UserPrivateKey:  getEnv("USER_PRIVATE_KEY"),
		LineApiKey:      getEnv("LINE_API_KEY"),
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

func (c Config) GetBscNetworkAndChainId() (string, int) {
	if c.UseTestNetwork {
		return bscNetwork, bscChainId
	}

	return bscTestNetwork, bacTestChainId
}
