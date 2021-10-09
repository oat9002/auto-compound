package services

import (
	"fmt"
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
	gasLimit uint64
}

func NewPancakeSwapService(client *ethclient.Client, chainId uint64, gasLimit uint64) (*PancakeSwapService, error) {
	contract, err := getMasterChefContract(client)
	service := &PancakeSwapService{client: client, contract: contract, chainId: chainId, gasLimit: gasLimit}

	return service, err
}

func getMasterChefContract(client *ethclient.Client) (*contracts.MasterChef, error) {
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")
	contract, err := contracts.NewMasterChef(pancakeMainStakingAddress, client)

	if err != nil {
		return nil, fmt.Errorf("create masterchef contract failed, %w", err)
	}

	return contract, nil
}

func (p *PancakeSwapService) GetPendingCakeFromSylupPool(address common.Address) (*big.Int, error) {
	cakePool := big.NewInt(int64(0)) // Sylup pool
	pendingCake, err := p.contract.PendingCake(utils.GetDefaultCallOpts(address), cakePool, address)

	if err != nil {
		return big.NewInt(int64(0)), fmt.Errorf("get pending cake failed, %w", err)
	}

	return pendingCake, nil
}

func (p *PancakeSwapService) CompoundEarnCake(privateKey string, amount *big.Int) (*types.Transaction, error) {
	txOpts, err := utils.GetDefautlTransactionOpts(p.client, privateKey, p.chainId, p.gasLimit)

	if err != nil {
		return nil, fmt.Errorf("get default trandaction opts failed, %w", err)
	}

	transaction, err := p.contract.EnterStaking(txOpts, amount)

	if err != nil {
		return nil, fmt.Errorf("enter staking failed, %w", err)
	}

	return transaction, nil
}
