package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

const bscNetwork string = "https://bsc-dataseed.binance.org/"
const bscChainId int = 56
const bscTestNetwork string = "https://data-seed-prebsc-1-s1.binance.org:8545/"
const bscTestChainId int = 97

type MessagingProvider int

const (
	None MessagingProvider = iota
	Line
	Telegram
)

type TelegramConfig struct {
	BotToken string
	ChatId   string
}

type Config struct {
	IsDevelopment     bool
	UseTestNetwork    bool
	OnlyCheckReward   bool
	ForceRun          bool
	UserAddress       string
	UserPrivateKey    string
	LineApiKey        string
	GasLimit          uint64
	GasPriceThreshold uint64
	QueryCron         string
	MutationCron      string
	Telegram          *TelegramConfig
	MessagingProvider MessagingProvider
}

var once sync.Once
var config *Config

const prefixEnv = "AUTO_COMPOUND_"
const defaultGasPriceThreshold = 10000000000
const defaultGasLimit = 3000000
const defaultQueryCron = "0 9,21 * * *"
const defaultMutationCron = "0 9 */7 * *"

func loadConfig() (*Config, error) {
	isDevelopmentFlag := flag.Bool("dev", false, "Run as development mode.")
	useTestNetWorkFlag := flag.Bool("testnet", false, "Use test network.")
	onlyCheckRewardFlag := flag.Bool("onlycheck", false, "Only check the reward.")
	forceRunFlag := flag.Bool("force", false, "Force run application (One time run, ignore schedule)")
	userAddressFlag := flag.String("address", "", "User public address.")
	userPrivateKeyFlag := flag.String("privatekey", "", "User private key. (Required if onlycheck is false)")
	lineApiKeyFlag := flag.String("lineapikey", "", "Send notification by line notify.")
	gasLimitFlag := flag.Uint64("gaslimit", defaultGasLimit, "Gas limit.")
	gasPriceThresholdFlag := flag.Uint64("gaspricethreshold", defaultGasPriceThreshold, "Threshld for gas price in Wei.")
	queryCronFlag := flag.String("querycron", defaultQueryCron, "Schedule for query reward e.g. 0 9,21 * * *")
	mutationCronFlag := flag.String("mutatationcron", defaultMutationCron, "Schedule for compound or harvest e.g. 0 9 */7 * *")
	telegramFlag := flag.String("telegram", "", "Telegram bot token and chat id, e.g. <bot_token>,<chat_id>")

	flag.Parse()
	godotenv.Load()

	isDevelopment := get("MODE", *isDevelopmentFlag, func(s string) bool { return s == "development" })
	useTestNetWork := get("USE_TEST_NETWORK", *useTestNetWorkFlag, func(s string) bool { return s == "true" })
	onlyCheckReward := get("ONLY_CHECK_REWARD", *onlyCheckRewardFlag, func(s string) bool { return s == "true" })
	forceRun := get("FORCE_RUN", *forceRunFlag, func(s string) bool { return s == "true" })
	userAddress := get("USER_ADDRESS", *userAddressFlag, func(s string) string { return s })
	userPrivateKey := get("USER_PRIVATE_KEY", *userPrivateKeyFlag, func(s string) string { return s })
	lineApiKey := get("LINE_API_KEY", *lineApiKeyFlag, func(s string) string { return s })
	gasLimit := get("GAS_LIMIT", *gasLimitFlag, func(s string) uint64 {
		limit, err := strconv.ParseUint(s, 10, 64)

		if err != nil {
			panic(fmt.Sprintf("Parse gasLimit config failed, %s", err))
		}

		return limit
	})
	gasPriceThreashold := get("GAS_PRICE_THRESHOLD", *gasPriceThresholdFlag, func(s string) uint64 {
		threshold, err := strconv.ParseUint(s, 10, 64)

		if err != nil {
			panic(fmt.Sprintf("Parse gasPriceThreashold config failed, %s", err))
		}

		return threshold
	})
	queryCron := get("QUERY_CRON", *queryCronFlag, func(s string) string { return s })
	mutationCron := get("MUTATION_CRON", *mutationCronFlag, func(s string) string { return s })
	telegramConfigString := get("TELEGRAM", *telegramFlag, func(s string) string { return s })
	var telegramConfig *TelegramConfig
	if configSplit := strings.Split(telegramConfigString, ","); len(configSplit) != 2 {
		telegramConfig = nil
	} else {
		telegramConfig = &TelegramConfig{
			BotToken: configSplit[0],
			ChatId:   configSplit[1],
		}
	}

	var messagingProvider MessagingProvider
	if telegramConfig != nil {
		messagingProvider = Telegram
	} else if lineApiKey != "" {
		messagingProvider = Line
	} else {
		messagingProvider = None
	}

	config = &Config{
		IsDevelopment:     isDevelopment,
		UseTestNetwork:    useTestNetWork,
		OnlyCheckReward:   onlyCheckReward,
		ForceRun:          forceRun,
		UserAddress:       userAddress,
		UserPrivateKey:    userPrivateKey,
		LineApiKey:        lineApiKey,
		GasLimit:          gasLimit,
		GasPriceThreshold: gasPriceThreashold,
		QueryCron:         queryCron,
		MutationCron:      mutationCron,
		Telegram:          telegramConfig,
		MessagingProvider: messagingProvider,
	}

	return config, nil
}

func getEnv(env string) (string, bool) {
	return os.LookupEnv(prefixEnv + env)
}

func get[T any](evnName string, defaultValue T, parseFunc func(string) T) T {
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
