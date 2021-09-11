package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services"
	"github.com/oat9002/auto-compound/utils"
)

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	myAddress := common.HexToAddress(conf.UserAddress)
	clientService := services.NewClientService()
	client, err := clientService.GetClient(conf.NetworkUrl)

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

	fmt.Println(utils.FromWei(pendingCake))
}
