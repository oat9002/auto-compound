package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/config"
	c "github.com/oat9002/auto-compound/services/cache"
	"github.com/oat9002/auto-compound/services/crypto"
	"github.com/oat9002/auto-compound/services/messaging"
	"github.com/oat9002/auto-compound/services/scheduler"
	"github.com/oat9002/auto-compound/services/user"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("auto-compound started.")
	conf, err := config.GetConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	if conf.IsDevelopment {
		executeTest(conf)
	} else {
		execute(conf)
	}

	createHealthcheckFile()
	fmt.Println("\nPlease Ctrl+C to exit.")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop
	removeHealthcheckfile()
}

func execute(conf config.Config) {
	network, chainId := conf.GetBscNetworkAndChainId()
	myAddress := common.HexToAddress(conf.UserAddress)
	clientService := crypto.NewClientService()
	client, err := clientService.GetClient(network)

	if err != nil {
		log.Fatal(err)
		return
	}

	pancakeSwapService, err := crypto.NewPancakeSwapService(client, uint64(chainId), conf.GasLimit, conf.GasPriceThreshold)

	if err != nil {
		log.Fatal(err)
		return
	}

	schedulerService := scheduler.NewSchedulerService(cron.New())
	lineService := messaging.NewLineService(&http.Client{}, conf.LineApiKey)
	telegramService := messaging.NewTelegramService(&http.Client{}, conf.Telegram.BotToken, conf.Telegram.ChatId)
	inMemCacheService := c.NewInMemCacheService(cache.DefaultExpiration, 10*time.Minute)
	messagingService := messaging.NewMessageService(conf, lineService, telegramService)
	userService := user.NewUserService(myAddress, conf.UserPrivateKey, messagingService, pancakeSwapService, inMemCacheService, client)

	if conf.ForceRun {
		userService.ProcessReward(conf.OnlyCheckReward)

		return
	}
	_, err = schedulerService.AddFunc(conf.QueryCron, func() {
		userService.ProcessReward(true)
	})

	if err != nil {
		log.Fatal(err)
	}

	if !conf.OnlyCheckReward {
		_, err = schedulerService.AddFunc(conf.MutationCron, func() {
			userService.ProcessReward(false)
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	schedulerService.RunAsync()
}

func createHealthcheckFile() {
	alive := []byte("Alive")
	err := os.WriteFile("/tmp/auto-compound-healthcheck.txt", alive, 0644)

	if err != nil {
		panic(err)
	}
}

func removeHealthcheckfile() {
	err := os.Remove("/tmp/auto-compound-healthcheck.txt")

	if err != nil {
		panic(err)
	}
}
