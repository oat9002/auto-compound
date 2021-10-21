package services

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/utils"
)

const previousPendingCakeCacheKey = "pendingCakeSylupPool"

type UserService struct {
	address                  common.Address
	privateKey               string
	lineService              *LineService
	pancakeSwapService       *PancakeSwapService
	pancakeCompoundThreshold float64
	cacheService             *CacheService
}

type balanceInfo struct {
	amount         *big.Int
	previousAmount *big.Int
	isCompound     bool
}

func NewUserService(address common.Address, privateKey string, lineService *LineService, pancakeSwapService *PancakeSwapService, pancakeCompoundThreshold float64, cacheService *CacheService) *UserService {
	userService := &UserService{
		address:                  address,
		privateKey:               privateKey,
		lineService:              lineService,
		pancakeSwapService:       pancakeSwapService,
		pancakeCompoundThreshold: pancakeCompoundThreshold,
		cacheService:             cacheService,
	}

	return userService
}

func (u *UserService) GetRewardMessage(balance map[string]balanceInfo) string {
	toReturn := fmt.Sprintln("[Your rewards]")

	for key, value := range balance {
		amount := utils.FromWei(value.amount)

		if value.isCompound {
			toReturn += fmt.Sprint(key, ": ", 0, " Compound: ", amount)
		} else {
			toReturn += fmt.Sprint(key, ": ", amount)
		}

		if value.previousAmount != nil {
			previousAmount := utils.FromWei(value.previousAmount)
			increasePercent := math.Round(((amount-previousAmount)*100/amount)*math.Pow10(2)) / math.Pow10(2)

			toReturn += fmt.Sprintln(" ↑ ", increasePercent, "%")
		} else {
			toReturn += fmt.Sprintln()
		}
	}

	return strings.TrimSuffix(toReturn, "\n")
}

func (u *UserService) ProcessReward(isOnlyCheckReward bool) {
	isCompoundCake := false
	pendingCake, err := u.pancakeSwapService.GetPendingCakeFromSylupPool(u.address)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if utils.FromWei(pendingCake) >= u.pancakeCompoundThreshold && !isOnlyCheckReward {
		_, err := u.pancakeSwapService.CompoundEarnCake(u.privateKey, pendingCake)

		if err != nil {
			fmt.Println(err.Error())
			u.lineService.Send(err.Error())

			return
		}∏

		isCompoundCake = true
	}

	balance := make(map[string]balanceInfo)

	if previousPendingCake, foundPreviousPendingCake := u.cacheService.Get(previousPendingCakeCacheKey); foundPreviousPendingCake {
		balance["cake"] = balanceInfo{amount: pendingCake, previousAmount: previousPendingCake.(*big.Int), isCompound: isCompoundCake}
	} else {
		balance["cake"] = balanceInfo{amount: pendingCake, previousAmount: nil, isCompound: isCompoundCake}
	}
	msg := u.GetRewardMessage(balance)
	u.lineService.Send(msg)
	u.cacheService.SetWithoutExpiry(previousPendingCakeCacheKey, pendingCake)
}
