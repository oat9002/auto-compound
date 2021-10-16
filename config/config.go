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
	GasPriceThreshold        uint64
	Cron                     string
}

var once sync.Once
var config *Config

const prefixEnv = "AUTO_COMPOUND_"
const defaultGasPriceThreshold = 10
const defaultPancakeCoumpoundThreshold = 0.5
const defaultGasLimit = 3000000
const defaultCron = "0 21 * * *"

func loadConfig() (*Config, error) {
	isDevelopmentFlag := flag.Bool("dev", false, "Run as development mode.")
	useTestNetWorkFlag := flag.Bool("testnet", false, "Use test network.")
	onlyCheckRewardFlag := flag.Bool("onlycheck", false, "Only check the reward.")
	forceRunFlag := flag.Bool("force", false, "Force run application (One time run, ignore schedule)")
	userAddressFlag := flag.String("address", "", "User public address.")
	userPrivateKeyFlag := flag.String("privatekey", "", "User private key. (Required if onlycheck is false)")
	lineApiKeyFlag := flag.String("lineapikey", "", "Send notification by line notify.")
	gasLimitFlag := flag.Uint64("gaslimit", defaultGasLimit, "Gas limit.")
	pancakeCompoundThresholdFlag := flag.Float64("pancakethreshold", defaultPancakeCoumpoundThreshold, "Threshold for amount of pancake to trigger compound.")
	gasPriceThresholdFlag := flag.Uint64("gaspricethreshold", defaultGasPriceThreshold, "Threshld for gas price in Wei.")
	cronFlag := flag.String("cron", defaultCron, "Schedule for running app e.g. 0 21 * * *")

	flag.Parse()
	godotenv.Load()

	isDevelopment := get("MODE", *isDevelopmentFlag, func(s string) interface{} { return s == "development" }).(bool)
	useTestNetWork := get("USE_TEST_NETWORK", *useTestNetWorkFlag, func(s string) interface{} { return s == "true" }).(bool)
	onlyCheckReward := get("ONLY_CHECK_REWARD", *onlyCheckRewardFlag, func(s string) interface{} { return s == "true" }).(bool)
	forceRun := get("FORCE_RUN", *forceRunFlag, func(s string) interface{} { return s == "true" }).(bool)
	userAddress := get("USER_ADDRESS", *userAddressFlag, func(s string) interface{} { return s }).(string)
	userPrivateKey := get("USER_PRIVATE_KEY", *userPrivateKeyFlag, func(s string) interface{} { return s }).(string)
	lineApiKey := get("LINE_API_KEY", *lineApiKeyFlag, func(s string) interface{} { return s }).(string)
	gasLimit := get("GAS_LIMIT", *gasLimitFlag, func(s string) interface{} {
		limit, _ := strconv.ParseUint(s, 10, 64)
		return limit
	}).(uint64)
	pancakeCompoundThreshold := get("PANCAKE_COMPOUND_THRESHOLD", *pancakeCompoundThresholdFlag, func(s string) interface{} {
		threshold, _ := strconv.ParseFloat(s, 64)
		return threshold
	}).(float64)
	gasPriceThreashold := get("GAS_PRICE_THRESHOLD", *gasPriceThresholdFlag, func(s string) interface{} {
		threshold, _ := strconv.ParseUint(s, 10, 64)
		return threshold
	}).(uint64)
	cron := get("CRON", *cronFlag, func(s string) interface{} { return s }).(string)

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
		GasPriceThreshold:        gasPriceThreashold,
		Cron:                     cron,
	}

	return config, nil
}

func getEnv(env string) (string, bool) {
	return os.LookupEnv(prefixEnv + env)
}

func get(evnName string, defaultValue interface{}, parseFunc func(string) interface{}) interface{} {
	if env, isExist := getEnv(evnName); isExist {
		return parseFunc(env)
	}

	return defaultValue
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
