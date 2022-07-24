package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services/crypto"
	"github.com/oat9002/auto-compound/utils"
)

func executeTest(conf config.Config) {
	network, chainId := conf.GetBscNetworkAndChainId()
	clientService := crypto.NewClientService()
	client, err := clientService.GetClient(network)

	if err != nil {
		log.Fatal(err)
		return
	}

	testContract, err := crypto.NewTestContractService(client, uint64(chainId), conf.GasLimit, conf.GasPriceThreshold)

	if err != nil {
		log.Fatal(err)
		return
	}

	transaction, err := testContract.Increase(conf.UserPrivateKey)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Get GasFee")

	gasFee, err := utils.GetGasFee(client, transaction.Hash())

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("gas fee: " + strconv.FormatFloat(gasFee, 'e', 6, 64))
}
