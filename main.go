package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services"
)

func main() {
	conf := config.GetConfig()
	myAddress := common.HexToAddress(conf.UserBEP20Address)

	client, err := GetClient("https://bsc-dataseed.binance.org/")

	if err != nil {
		log.Fatal(err)
		return
	}

	pancakeSwapService, err := services.NewPancakeSwapService(client)

	pendingCake, err := pancakeSwapService.GetPendingCakeFromSylupPool(myAddress)

	fmt.Println(pendingCake)
}

func GetClient(networkUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(networkUrl)
}
