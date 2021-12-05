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
	client                *ethclient.Client
	masterChefContract    *contracts.MasterChef
	smartChefInitilizable *contracts.SmartChefInitializable
	chainId               uint64
	gasLimit              uint64
	gasPriceThreshold     uint64
}

func NewPancakeSwapService(client *ethclient.Client, chainId uint64, gasLimit uint64, gasPriceThreshold uint64) (*PancakeSwapService, error) {
	masterChefContract, err := getMasterChefContract(client)

	if err != nil {
		return nil, err
	}

	smartChefInitializable, err := getSmartChefInitializable(client)

	if err != nil {
		return nil, err
	}

	service := &PancakeSwapService{
		client:                client,
		masterChefContract:    masterChefContract,
		smartChefInitilizable: smartChefInitializable,
		chainId:               chainId,
		gasLimit:              gasLimit,
		gasPriceThreshold:     gasPriceThreshold,
	}

	return service, nil
}

func getMasterChefContract(client *ethclient.Client) (*contracts.MasterChef, error) {
	pancakeMainStakingAddress := common.HexToAddress("0x73feaa1eE314F8c655E354234017bE2193C9E24E")
	contract, err := contracts.NewMasterChef(pancakeMainStakingAddress, client)

	if err != nil {
		return nil, fmt.Errorf("create masterchef contract failed, %w", err)
	}

	return contract, nil
}

func getSmartChefInitializable(client *ethclient.Client) (*contracts.SmartChefInitializable, error) {
	contractAddress := common.HexToAddress("0x6f660C58723922c6f866a058199FF4881019B4B5")
	contract, err := contracts.NewSmartChefInitializable(contractAddress, client)

	if err != nil {
		return nil, fmt.Errorf("create smartChefInitializable contract failed, %w", err)
	}

	return contract, nil
}

func (p *PancakeSwapService) GetPendingCakeFromSylupPool(address common.Address) (*big.Int, error) {
	cakePool := big.NewInt(int64(0)) // Sylup pool
	pendingCake, err := p.masterChefContract.PendingCake(utils.GetDefaultCallOpts(address), cakePool, address)

	if err != nil {
		return big.NewInt(int64(0)), fmt.Errorf("get pending cake failed, %w", err)
	}

	return pendingCake, nil
}

func (p *PancakeSwapService) GetPendingBetaFromSylupPool(address common.Address) (*big.Int, error) {
	pendingBeta, err := p.smartChefInitilizable.PendingReward(utils.GetDefaultCallOpts(address), address)

	if err != nil {
		return big.NewInt(int64(0)), fmt.Errorf("get pending beta failed, %w", err)
	}

	return pendingBeta, nil
}

func (p *PancakeSwapService) CompoundEarnCake(privateKey string, amount *big.Int) (*types.Transaction, error) {
	txOpts, err := utils.GetDefautlTransactionOpts(p.client, privateKey, p.chainId, p.gasLimit, p.gasPriceThreshold)

	if err != nil {
		return nil, fmt.Errorf("get default trandaction opts failed, %w", err)
	}

	transaction, err := p.masterChefContract.EnterStaking(txOpts, amount)

	if err != nil {
		return nil, fmt.Errorf("enter staking failed, %w", err)
	}

	return transaction, nil
}

func (p *PancakeSwapService) HarvestEarnBeta(privateKey string) (*types.Transaction, error) {
	txOpts, err := utils.GetDefautlTransactionOpts(p.client, privateKey, p.chainId, p.gasLimit, p.gasPriceThreshold)

	if err != nil {
		return nil, fmt.Errorf("get default trandaction opts failed, %w", err)
	}

	transaction, err := p.smartChefInitilizable.Deposit(txOpts, big.NewInt(0))

	if err != nil {
		return nil, fmt.Errorf("HarvestEarnBeta failed, %w", err)
	}

	return transaction, nil
}
