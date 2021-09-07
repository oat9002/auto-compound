package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services"
)

func main() {
	conf := config.GetConfig()
	myAddress := common.HexToAddress(conf.UserBEP20Address)
	clientService := services.NewClientService()
	client, err := clientService.GetClient("https://bsc-dataseed.binance.org/")

	if err != nil {
		log.Fatal(err)
		return
	}

	pancakeSwapService, err := services.NewPancakeSwapService(client)

	if err != nil {
		log.Fatal(err)
		return
	}

	pendingCake, err := pancakeSwapService.GetPendingCakeFromSylupPool(myAddress)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(pendingCake)
}
