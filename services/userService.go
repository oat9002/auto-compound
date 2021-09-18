package services

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oat9002/auto-compound/utils"
)

type userService struct {
	address            common.Address
	lineService        *lineService
	pancakeSwapService *pancakeSwapService
	schedulerService   *schedulerService
}

func NewUserService(address common.Address, lineService *lineService, pancakeSwapService *pancakeSwapService, schedulerService *schedulerService) *userService {
	userService := &userService{address: address, lineService: lineService, pancakeSwapService: pancakeSwapService, schedulerService: schedulerService}

	return userService
}

func (u *userService) GetRewardMessage(balance map[string]*big.Int) string {
	var toReturn = fmt.Sprintln("---Your rewards---")

	for key, value := range balance {
		toReturn += fmt.Sprintln(key, ": ", utils.FromWei(value))
	}

	return toReturn
}

func (u *userService) NotifyReward() {
	pendingCake, err := u.pancakeSwapService.GetPendingCakeFromSylupPool(u.address)

	if err != nil {
		log.Fatal(err)
		return
	}

	balance := make(map[string]*big.Int)
	balance["cake"] = pendingCake

	msg := u.GetRewardMessage(balance)

	u.schedulerService.AddFunc("* * * * *", func() {
		u.lineService.Send(msg)
	})

	u.schedulerService.RunAsync()
}
