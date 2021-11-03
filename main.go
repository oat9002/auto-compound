package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services"
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

	fmt.Println("\nPlease Ctrl+C to exit.")
	select {}
}

func execute(conf config.Config) {
	network, chainId := conf.GetBscNetworkAndChainId()
	myAddress := common.HexToAddress(conf.UserAddress)
	clientService := services.NewClientService()
	client, err := clientService.GetClient(network)

	if err != nil {
		log.Fatal(err)
		return
	}

	pancakeSwapService, err := services.NewPancakeSwapService(client, uint64(chainId), conf.GasLimit, conf.GasPriceThreshold)

	if err != nil {
		log.Fatal(err)
		return
	}

	schedulerService := services.NewSchedulerService(cron.New())
	lineService := services.NewLineService(&http.Client{}, conf.LineApiKey)
	cacheService := services.NewCacheService(cache.DefaultExpiration, 10*time.Minute)
	userService := services.NewUserService(myAddress, conf.UserPrivateKey, lineService, pancakeSwapService, conf.PancakeCompoundThreshold, cacheService, client)

	if conf.ForceRun {
		userService.ProcessReward(conf.OnlyCheckReward)
	} else {
		_, err := schedulerService.AddFunc(conf.Cron, func() {
			userService.ProcessReward(conf.OnlyCheckReward)
		})

		if err != nil {
			log.Fatal(err)
		}

		schedulerService.RunAsync()
	}
}
