package user

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	c "github.com/oat9002/auto-compound/services/cache"
	"github.com/oat9002/auto-compound/services/crypto"
	"github.com/oat9002/auto-compound/services/messaging"
	"github.com/oat9002/auto-compound/utils"
)

const previousPendingCakeCacheKey = "pendingCakeSylupPool"

type UserService struct {
	address            common.Address
	privateKey         string
	messagingService   messaging.MessagingService
	pancakeSwapService *crypto.PancakeSwapService
	cacheService       c.CacheService
	client             *ethclient.Client
}

type balanceInfo struct {
	amount         *big.Int
	previousAmount *big.Int
	isCompound     bool
	isHarvest      bool
	gasFee         float64
}

func NewUserService(address common.Address, privateKey string, messagingService messaging.MessagingService, pancakeSwapService *crypto.PancakeSwapService, cacheService c.CacheService, client *ethclient.Client) *UserService {
	userService := &UserService{
		address:            address,
		privateKey:         privateKey,
		messagingService:   messagingService,
		pancakeSwapService: pancakeSwapService,
		cacheService:       cacheService,
		client:             client,
	}

	return userService
}

func (u *UserService) GetRewardMessage(balance map[string]balanceInfo) string {
	toReturn := fmt.Sprintln("[Your rewards]")

	for key, value := range balance {
		amount := utils.FromWei(value.amount)

		if value.isCompound || value.isHarvest {
			compoundOrHarvest := "Compound"
			if value.isHarvest {
				compoundOrHarvest = "Harvest"
			}
			toReturn += fmt.Sprintf("%s: 0 %s: %.6f ", key, compoundOrHarvest, amount)
			if value.gasFee != 0 {
				toReturn += fmt.Sprintf("fee: %.6f bnb", value.gasFee)
			}

		} else {
			toReturn += fmt.Sprintf("%s: %.6f", key, amount)
		}

		if value.previousAmount != nil && !value.isCompound && !value.isHarvest && amount > 0 {
			previousAmount := utils.FromWei(value.previousAmount)
			increasePercent := math.Round(((amount-previousAmount)*100/amount)*math.Pow10(2)) / math.Pow10(2)

			if increasePercent > 0 {
				toReturn += fmt.Sprintln(" ??? ", increasePercent, "%")
			}
		} else {
			toReturn += fmt.Sprintln()
		}
	}

	return strings.TrimSuffix(toReturn, "\n")
}

// func (u *UserService) compoundOrHarvest(isOnlyCheckReward bool, prevCacheKey string, execute func() (*types.Transaction, error)) (bool, float64, error) {
// 	if isOnlyCheckReward {
// 		return false, 0, nil
// 	}

// 	tx, err := execute()

// 	if err != nil {
// 		return false, 0, err
// 	}

// 	gasFee, err := utils.GetGasFee(u.client, tx.Hash())

// 	if err != nil {
// 		return true, 0, err
// 	}

// 	u.cacheService.Delete(prevCacheKey)

// 	return true, gasFee, nil
// }

func (u *UserService) getPreviousToken(cacheKey string) *big.Int {
	if previousToken, foundPreviousToken := u.cacheService.Get(cacheKey); foundPreviousToken {
		return previousToken.(*big.Int)
	}

	return nil
}

func (u *UserService) ProcessReward(isOnlyCheckReward bool) {
	earningCake, err := u.pancakeSwapService.GetEarningCakeSinceLastActionFromCakePool(u.address)
	defer u.cacheService.SetWithoutExpiry(previousPendingCakeCacheKey, earningCake)

	if err != nil {
		fmt.Printf("GetPendingCakeFromSylupPool failed, %s\n", err)
		return
	}

	balance := make(map[string]balanceInfo)
	balance["cake"] = balanceInfo{
		amount:         earningCake,
		previousAmount: u.getPreviousToken(previousPendingCakeCacheKey),
		isCompound:     false,
		isHarvest:      false,
		gasFee:         0,
	}

	msg := u.GetRewardMessage(balance)
	err = u.messagingService.SendMessage(msg)
	if err != nil {
		log.Printf("send message failed, %s", err)
	}
}
