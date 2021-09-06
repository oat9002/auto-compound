package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	masterChef "github.com/oat9002/auto-compound/contracts"
	"github.com/oat9002/auto-compound/utils"
)

func GetMasterChefContract(client *ethclient.Client) (*masterChef.MasterChef, error) {
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")

	return masterChef.NewMasterChef(pancakeMainStakingAddress, client)
}

func GetPendingCakeFromSylupPool(contract *masterChef.MasterChef, address common.Address) (float64, error) {
	cakePool := big.NewInt(int64(0)) // Sylup pool
	pendingCake, err := contract.PendingCake(utils.GetDefaultCallOpts(address), cakePool, address)

	if err != nil {
		return 0, err
	}

	return utils.FromWei(pendingCake), nil
}
