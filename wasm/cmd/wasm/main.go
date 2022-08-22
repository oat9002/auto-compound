package main

import (
	"fmt"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/config"
	"github.com/oat9002/auto-compound/services/crypto"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("wasmGetUserInfoFromCakePool", js.FuncOf(getUserInfoFromCakePool))
	<-done
}

func getUserInfoFromCakePool(this js.Value, args []js.Value) interface{} {
	conf, _ := config.GetConfig()
	network, chainId := conf.GetBscNetworkAndChainId()
	clientService := crypto.NewClientService()
	client, err := clientService.GetClient(network)
	if err != nil {
		return fmt.Sprint(err)
	}
	pancakeSwapService, err := crypto.NewPancakeSwapService(client, uint64(chainId), 0, 0)
	if err != nil {
		return fmt.Sprint(err)
	}
	userInfo, err := pancakeSwapService.GetUserInfoFromCakePool(common.HexToAddress(args[0].String()))
	if err != nil {
		return fmt.Sprint(err)
	}

	toReturn := fmt.Sprintf("Lock amount: %s", userInfo.LockedAmount.Text(10))
	return toReturn
}
