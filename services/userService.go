package services

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oat9002/auto-compound/utils"
)

const previousPendingCakeCacheKey = "pendingCakeSylupPool"
const previousPendingBetaCacheKey = "pendingBetaSylupPool"

type UserService struct {
	address                  common.Address
	privateKey               string
	lineService              *LineService
	pancakeSwapService       *PancakeSwapService
	pancakeCompoundThreshold float64
	betaHarvestThreshold     float64
	cacheService             *CacheService
	client                   *ethclient.Client
}

type balanceInfo struct {
	amount              *big.Int
	previousAmount      *big.Int
	isCompoundOrHarvest bool
	gasFee              float64
	cacheKey            string
}

func NewUserService(address common.Address, privateKey string, lineService *LineService, pancakeSwapService *PancakeSwapService, pancakeCompoundThreshold float64, cacheService *CacheService, client *ethclient.Client, betaHarvestThreshold float64) *UserService {
	userService := &UserService{
		address:                  address,
		privateKey:               privateKey,
		lineService:              lineService,
		pancakeSwapService:       pancakeSwapService,
		pancakeCompoundThreshold: pancakeCompoundThreshold,
		betaHarvestThreshold:     betaHarvestThreshold,
		cacheService:             cacheService,
		client:                   client,
	}

	return userService
}

func (u *UserService) GetRewardMessage(balance map[string]balanceInfo) string {
	toReturn := fmt.Sprintln("[Your rewards]")

	for key, value := range balance {
		amount := utils.FromWei(value.amount)

		if value.isCompoundOrHarvest {
			toReturn += fmt.Sprint(key, ": ", 0, " Compound: ", amount)
			if value.gasFee != 0 {
				toReturn += fmt.Sprint(" Gas Fee: ", value.gasFee, " BNB")
			}

			u.cacheService.Delete(value.cacheKey)
		} else {
			toReturn += fmt.Sprint(key, ": ", amount)
		}

		if value.previousAmount != nil && !value.isCompoundOrHarvest && amount > 0 {
			previousAmount := utils.FromWei(value.previousAmount)
			increasePercent := math.Round(((amount-previousAmount)*100/amount)*math.Pow10(2)) / math.Pow10(2)

			toReturn += fmt.Sprintln(" ↑ ", increasePercent, "%")
		} else {
			toReturn += fmt.Sprintln()
		}
	}

	return strings.TrimSuffix(toReturn, "\n")
}

func (u *UserService) handleError(err error) {
	fmt.Println(err.Error())
	u.lineService.Send(err.Error())
}

func (u *UserService) compoundOrHarvest(pendingToken *big.Int, threshold float64, isOnlyCheckReward bool, execute func() (*types.Transaction, error)) (bool, float64, error) {
	if utils.FromWei(pendingToken) < threshold || isOnlyCheckReward {
		return false, 0, nil
	}

	tx, err := execute()

	if err != nil {
		return false, 0, err
	}

	gasFee, err := utils.GetGasFree(u.client, tx)

	if err != nil {
		fmt.Println(err.Error())
	}

	return true, gasFee, nil
}

func (u *UserService) getPreviousToken(cacheKey string) *big.Int {
	if previousToken, foundPreviousToken := u.cacheService.Get(cacheKey); foundPreviousToken {
		return previousToken.(*big.Int)
	}

	return nil
}

func (u *UserService) ProcessReward(isOnlyCheckReward bool) {
	pendingCake, err := u.pancakeSwapService.GetPendingCakeFromSylupPool(u.address)

	if err != nil {
		u.handleError(err)
		return
	}

	pendingBeta, err := u.pancakeSwapService.GetPendingBetaFromSylupPool(u.address)

	if err != nil {
		u.handleError(err)
		return
	}

	isCompoundCake, compoundCakeGasFee, err := u.compoundOrHarvest(pendingCake, u.pancakeCompoundThreshold, isOnlyCheckReward, func() (*types.Transaction, error) {
		return u.pancakeSwapService.CompoundEarnCake(u.privateKey, pendingCake)
	})

	if err != nil {
		u.handleError(err)
		return
	}

	isHarvestBeta, harvestBetaGasFee, err := u.compoundOrHarvest(pendingBeta, u.betaHarvestThreshold, isOnlyCheckReward, func() (*types.Transaction, error) {
		return u.pancakeSwapService.HarvestEarnBeta(u.privateKey)
	})

	if err != nil {
		u.handleError(err)
		return
	}

	balance := make(map[string]balanceInfo)
	balance["cake"] = balanceInfo{
		amount:              pendingCake,
		previousAmount:      u.getPreviousToken(previousPendingCakeCacheKey),
		isCompoundOrHarvest: isCompoundCake,
		gasFee:              compoundCakeGasFee,
		cacheKey:            previousPendingCakeCacheKey,
	}
	balance["beta"] = balanceInfo{
		amount:              pendingBeta,
		previousAmount:      u.getPreviousToken(previousPendingBetaCacheKey),
		isCompoundOrHarvest: isHarvestBeta,
		gasFee:              harvestBetaGasFee,
		cacheKey:            previousPendingBetaCacheKey,
	}

	msg := u.GetRewardMessage(balance)
	u.lineService.Send(msg)
	u.cacheService.SetWithoutExpiry(previousPendingCakeCacheKey, pendingCake)
	u.cacheService.SetWithoutExpiry(previousPendingBetaCacheKey, pendingBeta)
}
