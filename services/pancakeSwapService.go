package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/oat9002/auto-compound/contracts"
	"github.com/oat9002/auto-compound/utils"
)

type PancakeSwapService struct {
	client   *ethclient.Client
	contract *contracts.MasterChef
	chainId  uint64
}

func NewPancakeSwapService(client *ethclient.Client, chainId uint64) (*PancakeSwapService, error) {
	contract, err := getMasterChefContract(client)
	service := &PancakeSwapService{client: client, contract: contract, chainId: chainId}

	return service, err
}

func getMasterChefContract(client *ethclient.Client) (*contracts.MasterChef, error) {
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")

	return contracts.NewMasterChef(pancakeMainStakingAddress, client)
}

func (p *PancakeSwapService) GetPendingCakeFromSylupPool(address common.Address) (*big.Int, error) {
	cakePool := big.NewInt(int64(0)) // Sylup pool
	pendingCake, err := p.contract.PendingCake(utils.GetDefaultCallOpts(address), cakePool, address)

	if err != nil {
		return big.NewInt(int64(0)), err
	}

	return pendingCake, nil
}

func (p *PancakeSwapService) CompoundEarnCake(privateKey string, amount *big.Int) (*types.Transaction, error) {
	txOpts, err := utils.GetDefautlTransactionOpts(p.client, privateKey, p.chainId)

	if err != nil {
		return nil, err
	}

	transaction, err := p.contract.EnterStaking(txOpts, amount)

	return transaction, err
}
