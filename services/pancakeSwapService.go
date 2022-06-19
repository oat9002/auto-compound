package services

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/oat9002/auto-compound/contracts"
	"github.com/oat9002/auto-compound/utils"
)

type PancakeSwapService struct {
	client                *ethclient.Client
	masterChefContract    *contracts.MasterChef
	smartChefInitilizable *contracts.SmartChefInitializable
	cakePoolContract      *contracts.CakePool
	chainId               uint64
	gasLimit              uint64
	gasPriceThreshold     uint64
}

func NewPancakeSwapService(client *ethclient.Client, chainId uint64, gasLimit uint64, gasPriceThreshold uint64) (*PancakeSwapService, error) {
	masterChefContract, err := getMasterChefContract(client)

	if err != nil {
		return nil, err
	}

	smartChefInitializable, err := getSmartChefInitializableContract(client)

	if err != nil {
		return nil, err
	}

	cakePoolContract, err := getCakePoolContract(client)

	if err != nil {
		return nil, err
	}

	service := &PancakeSwapService{
		client:                client,
		masterChefContract:    masterChefContract,
		smartChefInitilizable: smartChefInitializable,
		cakePoolContract:      cakePoolContract,
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

func getSmartChefInitializableContract(client *ethclient.Client) (*contracts.SmartChefInitializable, error) {
	contractAddress := common.HexToAddress("0x6f660C58723922c6f866a058199FF4881019B4B5")
	contract, err := contracts.NewSmartChefInitializable(contractAddress, client)

	if err != nil {
		return nil, fmt.Errorf("create smartChefInitializable contract failed, %w", err)
	}

	return contract, nil
}

func getCakePoolContract(client *ethclient.Client) (*contracts.CakePool, error) {
	contractAddress := common.HexToAddress("0x45c54210128a065de780C4B0Df3d16664f7f859e")
	contract, err := contracts.NewCakePool(contractAddress, client)

	if err != nil {
		return nil, fmt.Errorf("create cakePool contract failed, %w", err)
	}

	return contract, nil
}

func (p *PancakeSwapService) GetEarningCakeSinceLastActionFromCakePool(address common.Address) (*big.Int, error) {
	userInfo, err := p.cakePoolContract.UserInfo(utils.GetDefaultCallOpts(address), address)

	if err != nil {
		return big.NewInt(int64(0)), fmt.Errorf("get user info from cake pool failed, %w", err)
	}

	pricePerFullShare, err := p.cakePoolContract.GetPricePerFullShare(utils.GetDefaultCallOpts(address))

	if err != nil {
		return big.NewInt(int64(0)), fmt.Errorf("get price per full share from cake pool failed, %w", err)
	}

	stakedCake := &big.Int{}
	earningCake := &big.Int{}
	i, pow := big.NewInt(10), big.NewInt(18)
	i.Exp(i, pow, nil)
	stakedCake.Mul(userInfo.Shares, pricePerFullShare).Div(stakedCake, i).Sub(stakedCake, userInfo.UserBoostedShare)
	earningCake.Sub(stakedCake, userInfo.CakeAtLastUserAction)

	return earningCake, nil
}
