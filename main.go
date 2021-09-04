package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	masterChef "github.com/oat9002/auto-compound/contracts"
)

func main() {
	myAddress := common.HexToAddress("0xf07dfFaFA1a9E49345f919040b1cbd6e4dDd5706")
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")

	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")

	if err != nil {
		log.Fatal(err)
	}

	instanceContract, err := masterChef.NewMasterChef(pancakeMainStakingAddress, client)
	cakePool := big.NewInt(int64(0))

	pendingCake, err := instanceContract.PendingCake(cakePool, myAddress)

	fmt.Println(pendingCake)
}
