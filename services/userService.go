package services

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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
	cacheService             *CacheService
	client                   *ethclient.Client
}

type balanceInfo struct {
	amount         *big.Int
	previousAmount *big.Int
	isCompound     bool
	gasFee         float64
	cacheKey       string
}

func NewUserService(address common.Address, privateKey string, lineService *LineService, pancakeSwapService *PancakeSwapService, pancakeCompoundThreshold float64, cacheService *CacheService, client *ethclient.Client) *UserService {
	userService := &UserService{
		address:                  address,
		privateKey:               privateKey,
		lineService:              lineService,
		pancakeSwapService:       pancakeSwapService,
		pancakeCompoundThreshold: pancakeCompoundThreshold,
		cacheService:             cacheService,
		client:                   client,
	}

	return userService
}

func (u *UserService) GetRewardMessage(balance map[string]balanceInfo) string {
	toReturn := fmt.Sprintln("[Your rewards]")

	for key, value := range balance {
		amount := utils.FromWei(value.amount)

		if value.isCompound {
			toReturn += fmt.Sprint(key, ": ", 0, " Compound: ", amount)
			if value.gasFee != 0 {
				toReturn += fmt.Sprint(" Gas Fee: ", value.gasFee, " BNB")
			}

			u.cacheService.Delete(value.cacheKey)
		} else {
			toReturn += fmt.Sprint(key, ": ", amount)
		}

		if value.previousAmount != nil && !value.isCompound && amount > 0 {
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

func (u *UserService) compoundCake(pendingCake *big.Int, isOnlyCheckReward bool) (bool, float64, error) {
	if utils.FromWei(pendingCake) < u.pancakeCompoundThreshold && isOnlyCheckReward {
		return false, 0, nil
	}

	tx, err := u.pancakeSwapService.CompoundEarnCake(u.privateKey, pendingCake)

	if err != nil {
		return false, 0, err
	}

	gasFee, err := utils.GetGasFree(u.client, tx)

	if err != nil {
		return false, 0, err
	}

	return true, gasFee, nil
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

	isCompoundCake, compoundCakeGasFee, err := u.compoundCake(pendingCake, isOnlyCheckReward)

	if err != nil {
		u.handleError(err)
		return
	}

	balance := make(map[string]balanceInfo)

	if previousPendingCake, foundPreviousPendingCake := u.cacheService.Get(previousPendingCakeCacheKey); foundPreviousPendingCake {
		balance["cake"] = balanceInfo{amount: pendingCake, previousAmount: previousPendingCake.(*big.Int), isCompound: isCompoundCake, gasFee: compoundCakeGasFee, cacheKey: previousPendingCakeCacheKey}
	} else {
		balance["cake"] = balanceInfo{amount: pendingCake, previousAmount: nil, isCompound: isCompoundCake, gasFee: compoundCakeGasFee, cacheKey: previousPendingCakeCacheKey}
	}

	if previousPendingBeta, foundPreviousPendingBeta := u.cacheService.Get(previousPendingBetaCacheKey); foundPreviousPendingBeta {
		balance["beta"] = balanceInfo{amount: pendingBeta, previousAmount: previousPendingBeta.(*big.Int), isCompound: false, gasFee: 0, cacheKey: previousPendingBetaCacheKey}
	} else {
		balance["beta"] = balanceInfo{amount: pendingBeta, previousAmount: nil, isCompound: false, gasFee: 0, cacheKey: previousPendingBetaCacheKey}
	}

	msg := u.GetRewardMessage(balance)
	u.lineService.Send(msg)
	u.cacheService.SetWithoutExpiry(previousPendingCakeCacheKey, pendingCake)
	u.cacheService.SetWithoutExpiry(previousPendingBetaCacheKey, pendingBeta)
}
