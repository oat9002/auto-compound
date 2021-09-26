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

	execute(conf)

	fmt.Println("\nPlease any key to exit.")
	fmt.Scanln()
}

func execute(conf config.Config) {
	if conf.IsTest {
		executeTest(conf)
		return
	}

	myAddress := common.HexToAddress(conf.UserAddress)
	clientService := services.NewClientService()
	client, err := clientService.GetClient(conf.NetworkUrl)

	if err != nil {
		log.Fatal(err)
		return
	}

	pancakeSwapService, err := services.NewPancakeSwapService(client, conf.ChainId)

	if err != nil {
		log.Fatal(err)
		return
	}

	schedulerService := services.NewSchedulerService(cron.New())
	lineService := services.NewLineService(&http.Client{}, conf.LineApiKey)
	userService := services.NewUserService(myAddress, conf.UserPrivateKey, lineService, pancakeSwapService)

	schedulerService.AddFunc("0 21 * * *", func() {
		userService.ProcessReward()
	})

	schedulerService.RunAsync()
}

func executeTest(conf config.Config) {
	clientService := services.NewClientService()
	client, err := clientService.GetClient(conf.NetworkUrl)

	if err != nil {
		log.Fatal(err)
		return
	}

	testContract, err := services.NewTestContractService(client, conf.ChainId)

	if err != nil {
		log.Fatal(err)
		return
	}

	transaction, err := testContract.Increase(conf.UserPrivateKey)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(transaction.Value())
	fmt.Println(transaction.Cost())
	fmt.Println(transaction.Gas())
	fmt.Println(transaction.GasPrice())
}
