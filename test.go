package main

import (
	"fmt"
	"log"

	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services"
)

func executeTest(conf config.Config) {
	network, chainId := conf.GetBscNetworkAndChainId()
	clientService := services.NewClientService()
	client, err := clientService.GetClient(network)

	if err != nil {
		log.Fatal(err)
		return
	}

	testContract, err := services.NewTestContractService(client, uint64(chainId))

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
