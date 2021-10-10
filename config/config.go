package config

import (
	"flag"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

const bscNetwork string = "https://bsc-dataseed.binance.org/"
const bscChainId int = 56
const bscTestNetwork string = "https://data-seed-prebsc-1-s1.binance.org:8545/"
const bscTestChainId int = 97

type Config struct {
	IsDevelopment            bool
	UseTestNetwork           bool
	OnlyCheckReward          bool
	ForceRun                 bool
	UserAddress              string
	UserPrivateKey           string
	LineApiKey               string
	GasLimit                 uint64
	PancakeCompoundThreshold float64
}

var once sync.Once
var config *Config

func loadConfig() (*Config, error) {
	isDevelopmentFlag := flag.Bool("dev", false, "Run as development mode.")
	useTestNetWorkFlag := flag.Bool("testnet", false, "Use test network.")
	onlyCheckRewardFlag := flag.Bool("onlycheck", false, "Only check the reward.")
	forceRunFlag := flag.Bool("force", false, "Force run application (One time run, ignore schedule)")
	userAddressFlag := flag.String("address", "", "User public address.")
	userPrivateKeyFlag := flag.String("privatekey", "", "User private key.")
	lineApiKeyFlag := flag.String("lineapikey", "", "Send notification by line notify.")
	gasLimitFlag := flag.Uint64("gaslimit", 3000000, "Gas limit. Default: 3000000")
	pancakeCompoundThresholdFlag := flag.Float64("pancakethreshold", 0.5, "Threshold for amount of pancake to trigger compound. Default: 0.5")

	flag.Parse()
	godotenv.Load()

	isDevelopment := *isDevelopmentFlag || getEnv("MODE") == "development"
	useTestNetWork := *useTestNetWorkFlag || getEnv("USE_TEST_NETWORK") == "true"
	onlyCheckReward := *onlyCheckRewardFlag || getEnv("ONLY_CHECK_REWARD") == "true"
	forceRun := *forceRunFlag || getEnv("FORCE_RUN") == "true"
	userAddress := *userAddressFlag
	if userAddress == "" && getEnv("USER_ADDRESS") != "" {
		userAddress = getEnv("USER_ADDRESS")
	}
	userPrivateKey := *userPrivateKeyFlag
	if userPrivateKey == "" && getEnv("USER_PRIVATE_KEY") != "" {
		userPrivateKey = getEnv("USER_PRIVATE_KEY")
	}
	lineApiKey := *lineApiKeyFlag
	if lineApiKey == "" && getEnv("LINE_API_KEY") != "" {
		lineApiKey = getEnv("LINE_API_KEY")
	}
	gasLimit := *gasLimitFlag
	if gasLimit == 3000000 && getEnv("GAS_LIMIT") != "3000000" {
		gasLimit, _ = strconv.ParseUint(getEnv("GAS_LIMIT"), 10, 64)
	}
	pancakeCompoundThreshold := *pancakeCompoundThresholdFlag
	if pancakeCompoundThreshold == 0.5 && getEnv("PANCAKE_COMPOUND_THRESHOLD") != "0.5" {
		pancakeCompoundThreshold, _ = strconv.ParseFloat(getEnv("PANCAKE_COMPOUND_THRESHOLD"), 64)
	}

	config = &Config{
		IsDevelopment:            isDevelopment,
		UseTestNetwork:           useTestNetWork,
		OnlyCheckReward:          onlyCheckReward,
		ForceRun:                 forceRun,
		UserAddress:              userAddress,
		UserPrivateKey:           userPrivateKey,
		LineApiKey:               lineApiKey,
		GasLimit:                 gasLimit,
		PancakeCompoundThreshold: pancakeCompoundThreshold,
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
		return bscTestNetwork, bscTestChainId
	}

	return bscNetwork, bscChainId
}
