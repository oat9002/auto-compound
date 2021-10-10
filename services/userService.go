package services

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/utils"
)

type UserService struct {
	address                  common.Address
	privateKey               string
	lineService              *LineService
	pancakeSwapService       *PancakeSwapService
	pancakeCompoundThreshold float64
}

func NewUserService(address common.Address, privateKey string, lineService *LineService, pancakeSwapService *PancakeSwapService, pancakeCompoundThreshold float64) *UserService {
	userService := &UserService{
		address:                  address,
		privateKey:               privateKey,
		lineService:              lineService,
		pancakeSwapService:       pancakeSwapService,
		pancakeCompoundThreshold: pancakeCompoundThreshold,
	}

	return userService
}

func (u *UserService) GetRewardMessage(balance map[string]*big.Int) string {
	var toReturn = fmt.Sprintln("[Your rewards]")

	for key, value := range balance {
		toReturn += fmt.Sprintln(key, ": ", utils.FromWei(value))
	}

	return toReturn
}

func (u *UserService) ProcessReward(isOnlyCheckReward bool) {
	var msg string
	isCompound := false
	pendingCake, err := u.pancakeSwapService.GetPendingCakeFromSylupPool(u.address)

	if err != nil {
		log.Fatal(err)
		return
	}

	if utils.FromWei(pendingCake) >= u.pancakeCompoundThreshold && !isOnlyCheckReward {
		_, err := u.pancakeSwapService.CompoundEarnCake(u.privateKey, pendingCake)

		if err != nil {
			log.Fatal(err)
			return
		}

		isCompound = true
	}

	balance := make(map[string]*big.Int)

	if isCompound {
		balance["cake"] = big.NewInt(0)
		msg = fmt.Sprintln(u.GetRewardMessage(balance)) + fmt.Sprintln("CompoundEarnCake", ": ", utils.FromWei(pendingCake))
	} else {
		balance["cake"] = pendingCake
		msg = fmt.Sprintln(u.GetRewardMessage(balance))
	}

	u.lineService.Send(msg)
}
