package config

import (
	"flag"
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
	ForceRun        bool
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

	isDevelopmentFlag := *flag.Bool("dev", false, "Run as development mode.")
	useTestNetWorkFlag := *flag.Bool("testnet", false, "Use test network.")
	onlyCheckRewardFlag := *flag.Bool("onlycheck", false, "Only check the reward.")
	forceRunFlag := *flag.Bool("force", false, "Force run application (One time run, ignore schedule)")
	userAddressFlag := *flag.String("address", "", "User public address.")
	userPrivateKeyFlag := *flag.String("privatekey", "", "User private key.")
	lineApiKeyFlag := *flag.String("lineapikey", "", "Send notification by line notify.")

	isDevelopment := isDevelopmentFlag || getEnv("MODE") == "development"
	useTestNetWork := useTestNetWorkFlag || getEnv("USE_TEST_NETWORK") == "true"
	onlyCheckReward := onlyCheckRewardFlag || getEnv("ONLY_CHECK_REWARD") == "true"
	forceRun := forceRunFlag || getEnv("FORCE_RUN") == "true"
	userAddress := userAddressFlag
	if userAddress == "" {
		userAddress = getEnv("USER_ADDRESS")
	}
	userPrivateKey := userPrivateKeyFlag
	if userPrivateKey == "" {
		userPrivateKey = getEnv("USER_PRIVATE_KEY")
	}
	lineApiKey := lineApiKeyFlag
	if lineApiKey == "" {
		lineApiKey = getEnv("LINE_API_KEY")
	}

	config = &Config{
		IsDevelopment:   isDevelopment,
		UseTestNetwork:  useTestNetWork,
		OnlyCheckReward: onlyCheckReward,
		ForceRun:        forceRun,
		UserAddress:     userAddress,
		UserPrivateKey:  userPrivateKey,
		LineApiKey:      lineApiKey,
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
