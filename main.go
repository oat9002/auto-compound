package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	pancakeMainStakingAddress := "0x73feaa1eE314F8c655E354234017bE2193C9E24E"
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")

	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(int64(1))

	block, err := client.BlockByNumber(context.Background(), blockNumber)

	fmt.Println(pancakeMainStakingAddress)
	fmt.Println(block.Body())
}
