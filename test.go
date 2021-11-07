package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services"
	"github.com/oat9002/auto-compound/utils"
)

func executeTest(conf config.Config) {
	network, chainId := conf.GetBscNetworkAndChainId()
	clientService := services.NewClientService()
	client, err := clientService.GetClient(network)

	if err != nil {
		log.Fatal(err)
		return
	}

	testContract, err := services.NewTestContractService(client, uint64(chainId), conf.GasLimit, conf.GasPriceThreshold)

	if err != nil {
		log.Fatal(err)
		return
	}

	transaction, err := testContract.Increase(conf.UserPrivateKey)

	if err != nil {
		log.Fatal(err)
		return
	}


	receipt, err := client.TransactionReceipt(context.Background(), transaction.Hash())

	if err != nil {
		log.Fatal(err)
		return
	}
	gasFee := math.Round(float64(receipt.GasUsed)*utils.FromWei(transaction.GasPrice())*math.Pow10(6)) / math.Pow10(6)

	fmt.Println(transaction.Value())
	fmt.Println(transaction.Cost())
	fmt.Println(transaction.Gas())
	fmt.Println("gas price: " + transaction.GasPrice().Text(10))
	fmt.Println("gas use: " + strconv.FormatUint(receipt.GasUsed, 10))
	fmt.Println("gas fee: " + strconv.FormatFloat(gasFee, 'd', 6, 64))
}
