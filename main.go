package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	masterChef "github.com/oat9002/auto-compound/contracts"
)

func main() {
	myAddress := common.HexToAddress("0xf07dfFaFA1a9E49345f919040b1cbd6e4dDd5706")
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")
	callOpts := &bind.CallOpts{Pending: false, From: myAddress, BlockNumber: nil, Context: nil}
	cakePool := big.NewInt(int64(0)) // Sylup pool

	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")

	if err != nil {
		log.Fatal(err)
	}

	instanceContract, err := masterChef.NewMasterChef(pancakeMainStakingAddress, client)

	pendingCake, err := instanceContract.PendingCake(callOpts, cakePool, myAddress)

	fmt.Println(fromWei(pendingCake))
}

func fromWei(data *big.Int) float64 {
	toReturn := float64(data.Int64()) / math.Pow10(18)

	return math.Round(toReturn*math.Pow10(6)) / math.Pow10(6)
}
