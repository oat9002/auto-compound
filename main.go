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

	client, err := GetClient("https://bsc-dataseed.binance.org/")

	if err != nil {
		log.Fatal(err)
		return
	}

	contract, err := GetMasterChefContract(client)

	pendingCake, err := GetPendingCakeFromSylupPool(contract, myAddress)

	fmt.Println(pendingCake)
}

func GetClient(networkUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(networkUrl)
}

func GetMasterChefContract(client *ethclient.Client) (*masterChef.MasterChef, error) {
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")

	return masterChef.NewMasterChef(pancakeMainStakingAddress, client)
}

func GetPendingCakeFromSylupPool(contract *masterChef.MasterChef, address common.Address) (float64, error) {
	cakePool := big.NewInt(int64(0)) // Sylup pool
	pendingCake, err := contract.PendingCake(GetDefaultCallOpts(address), cakePool, address)

	if err != nil {
		return 0, err
	}

	return FromWei(pendingCake), nil
}

func GetDefaultCallOpts(address common.Address) *bind.CallOpts {
	return &bind.CallOpts{Pending: false, From: address, BlockNumber: nil, Context: nil}
}

func FromWei(data *big.Int) float64 {
	toReturn := float64(data.Int64()) / math.Pow10(18)

	return math.Round(toReturn*math.Pow10(6)) / math.Pow10(6)
}
