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
	conf, err := config.GetConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	myAddress := common.HexToAddress(conf.UserAddress)
	clientService := services.NewClientService()
	client, err := clientService.GetClient(conf.BscNetworkUrl)

	if err != nil {
		log.Fatal(err)
		return
	}

	pancakeSwapService, err := services.NewPancakeSwapService(client)

	if err != nil {
		log.Fatal(err)
		return
	}

	schedulerService := services.NewSchedulerService(cron.New())
	lineService := services.NewLineService(&http.Client{}, conf.LineApiKey)
	userService := services.NewUserService(myAddress, lineService, pancakeSwapService, schedulerService)

	userService.NotifyReward()

	fmt.Println("Please any key to exit.")
	fmt.Scanln()
}
