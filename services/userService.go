package services

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/utils"
)

type UserService struct {
	address            common.Address
	privateKey         string
	lineService        *LineService
	pancakeSwapService *PancakeSwapService
}

func NewUserService(address common.Address, privateKey string, lineService *LineService, pancakeSwapService *PancakeSwapService) *UserService {
	userService := &UserService{address: address, lineService: lineService, pancakeSwapService: pancakeSwapService}

	return userService
}

func (u *UserService) GetRewardMessage(balance map[string]*big.Int) string {
	var toReturn = fmt.Sprintln("[Your rewards]")

	for key, value := range balance {
		toReturn += fmt.Sprintln(key, ": ", utils.FromWei(value))
	}

	return toReturn
}

func (u *UserService) ProcessReward() {
	var msg string
	isCompound := false
	pendingCake, err := u.pancakeSwapService.GetPendingCakeFromSylupPool(u.address)

	if err != nil {
		log.Fatal(err)
		return
	}

	if utils.FromWei(pendingCake) >= 0.5 {
		u.pancakeSwapService.CompoundEarnCake(u.privateKey, pendingCake)
		isCompound = true
	}

	balance := make(map[string]*big.Int)

	if isCompound {
		balance["cake"] = big.NewInt(0)
		msg = fmt.Sprintln(u.GetRewardMessage(balance)) + fmt.Sprintln() + fmt.Sprintln("CompoundEarnCake", ": ", utils.FromWei(pendingCake))
	} else {
		balance["cake"] = pendingCake
		msg = fmt.Sprintln(u.GetRewardMessage(balance))
	}

	u.lineService.Send(msg)
}
