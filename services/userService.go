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
	lineService        *LineService
	pancakeSwapService *PancakeSwapService
	schedulerService   *SchedulerService
}

func NewUserService(address common.Address, lineService *LineService, pancakeSwapService *PancakeSwapService, schedulerService *SchedulerService) *UserService {
	userService := &UserService{address: address, lineService: lineService, pancakeSwapService: pancakeSwapService, schedulerService: schedulerService}

	return userService
}

func (u *UserService) GetRewardMessage(balance map[string]*big.Int) string {
	var toReturn = fmt.Sprintln("[Your rewards]")

	for key, value := range balance {
		toReturn += fmt.Sprintln(key, ": ", utils.FromWei(value))
	}

	return toReturn
}

func (u *UserService) NotifyReward() {
	pendingCake, err := u.pancakeSwapService.GetPendingCakeFromSylupPool(u.address)

	if err != nil {
		log.Fatal(err)
		return
	}

	balance := make(map[string]*big.Int)
	balance["cake"] = pendingCake

	msg := u.GetRewardMessage(balance)

	u.schedulerService.AddFunc("0 21 * * *", func() {
		u.lineService.Send(msg)
	})

	u.schedulerService.RunAsync()
}
