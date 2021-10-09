package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services"
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

	fmt.Println("\nPlease any key to exit.")
	fmt.Scanln()
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

	pancakeSwapService, err := services.NewPancakeSwapService(client, uint64(chainId), conf.GasLimit)

	if err != nil {
		log.Fatal(err)
		return
	}

	schedulerService := services.NewSchedulerService(cron.New())
	lineService := services.NewLineService(&http.Client{}, conf.LineApiKey)
	userService := services.NewUserService(myAddress, conf.UserPrivateKey, lineService, pancakeSwapService)

	if conf.ForceRun {
		userService.ProcessReward(conf.OnlyCheckReward)
	} else {
		schedulerService.AddFunc("0 21 * * *", func() {
			userService.ProcessReward(conf.OnlyCheckReward)
		})

		schedulerService.RunAsync()
	}
}
