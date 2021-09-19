package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/oat9002/auto-compound/contracts"
	"github.com/oat9002/auto-compound/utils"
)

type pancakeSwapService struct {
	contract *contracts.MasterChef
}

func NewPancakeSwapService(client *ethclient.Client) (*pancakeSwapService, error) {
	contract, err := getMasterChefContract(client)
	service := &pancakeSwapService{contract: contract}

	return service, err
}

func getMasterChefContract(client *ethclient.Client) (*contracts.MasterChef, error) {
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")

	return contracts.NewMasterChef(pancakeMainStakingAddress, client)
}

func (p *pancakeSwapService) GetPendingCakeFromSylupPool(address common.Address) (*big.Int, error) {
	cakePool := big.NewInt(int64(0)) // Sylup pool
	pendingCake, err := p.contract.PendingCake(utils.GetDefaultCallOpts(address), cakePool, address)

	if err != nil {
		return big.NewInt(int64(0)), err
	}

	return pendingCake, nil
}
