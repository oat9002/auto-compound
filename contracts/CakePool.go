// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CakePoolMetaData contains all meta data concerning the CakePool contract.
var CakePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractIMasterChefV2\",\"name\":\"_masterchefV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastDepositedTime\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"free\",\"type\":\"bool\"}],\"name\":\"FreeFeeUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"boostContract\",\"type\":\"address\"}],\"name\":\"NewBoostContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostWeight\",\"type\":\"uint256\"}],\"name\":\"NewBoostWeight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationFactor\",\"type\":\"uint256\"}],\"name\":\"NewDurationFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationFactorOverdue\",\"type\":\"uint256\"}],\"name\":\"NewDurationFactorOverdue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLockDuration\",\"type\":\"uint256\"}],\"name\":\"NewMaxLockDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"overdueFee\",\"type\":\"uint256\"}],\"name\":\"NewOverdueFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"NewPerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFeeContract\",\"type\":\"uint256\"}],\"name\":\"NewPerformanceFeeContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"NewTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockFreeDuration\",\"type\":\"uint256\"}],\"name\":\"NewUnlockFreeDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"VCake\",\"type\":\"address\"}],\"name\":\"NewVCakeContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawFeeContract\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFeeContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawFeePeriod\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFeePeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOST_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_WEIGHT_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_FACTOR_OVERDUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LOCK_DURATION_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OVERDUE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PERFORMANCE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_FEE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAW_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNLOCK_FREE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VCake\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakePoolPID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"calculateOverdueFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"calculatePerformanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateTotalPendingCakeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"calculateWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockDuration\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeOverdueFeeUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freePerformanceFeeUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeWithdrawFeeUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricePerFullShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"inCaseTokensGetStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"dummyToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterchefV2\",\"outputs\":[{\"internalType\":\"contractIMasterChefV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overdueFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_boostContract\",\"type\":\"address\"}],\"name\":\"setBoostContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_boostWeight\",\"type\":\"uint256\"}],\"name\":\"setBoostWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_durationFactor\",\"type\":\"uint256\"}],\"name\":\"setDurationFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_durationFactorOverdue\",\"type\":\"uint256\"}],\"name\":\"setDurationFactorOverdue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_free\",\"type\":\"bool\"}],\"name\":\"setFreePerformanceFeeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLockDuration\",\"type\":\"uint256\"}],\"name\":\"setMaxLockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_overdueFee\",\"type\":\"uint256\"}],\"name\":\"setOverdueFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_free\",\"type\":\"bool\"}],\"name\":\"setOverdueFeeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_performanceFee\",\"type\":\"uint256\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_performanceFeeContract\",\"type\":\"uint256\"}],\"name\":\"setPerformanceFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockFreeDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockFreeDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_VCake\",\"type\":\"address\"}],\"name\":\"setVCakeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFeeContract\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFeePeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_free\",\"type\":\"bool\"}],\"name\":\"setWithdrawFeeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBoostDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cakeAtLastUserAction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUserActionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBoostedShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawByAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405262093a80600e556301e13380600f556301e1338060105562ed4e0060115564e8d4a5100060125560c860135560c8601455600a601555600a60165564e8d4a510006017556203f4806018553480156200005c57600080fd5b5060405162005369380380620053698339810160408190526200007f9162000157565b620000936200008d62000103565b62000107565b6000805460ff60a01b191690556001600160601b0319606096871b81166080529490951b90931660a052600880546001600160a01b039384166001600160a01b0319918216179091556009805492841692821692909217909155600a8054929093169116179055600b55620001f7565b3390565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008060008060008060c0878903121562000170578182fd5b86516200017d81620001de565b60208801519096506200019081620001de565b6040880151909550620001a381620001de565b6060880151909450620001b681620001de565b6080880151909350620001c981620001de565b8092505060a087015190509295509295509295565b6001600160a01b0381168114620001f457600080fd5b50565b60805160601c60a05160601c6150e46200028560003960008181610ccb01528181610d880152818161176e015281816125d901528181613b960152613c7001526000818161151701528181611af50152818161276801528181612d760152818161312c0152818161317e015281816134820152818161355d01528181613ec301526140fb01526150e46000f3fe608060405234801561001057600080fd5b50600436106105145760003560e01c806387788782116102a1578063cb528b521161016b578063e464c623116100e3578063f2fde38b11610097578063f851a4401161007c578063f851a440146108fb578063fc0c546a14610903578063fd253b641461090b57610514565b8063f2fde38b146108d5578063f786b958146108e857610514565b8063e73008bc116100c8578063e73008bc146108b2578063e941fa78146108ba578063f0f44260146108c257610514565b8063e464c623146108a2578063e4b37ef5146108aa57610514565b8063def68a9c1161013a578063df10b4e61161011f578063df10b4e61461087f578063dfcedeee14610887578063e2bbb1581461088f57610514565b8063def68a9c14610859578063def7869d1461086c57610514565b8063cb528b5214610836578063ccd34cd51461076e578063d4b0de2f1461083e578063d826ed061461084657610514565b8063acaf88cd11610219578063bc75f4b8116101cd578063beba0fa0116101b2578063beba0fa0146107fd578063c54d349c14610810578063c6ed51be1461082357610514565b8063bc75f4b8146107ed578063bdca9165146107f557610514565b8063b6857844116101fe578063b6857844146105a2578063b6ac642a146107c7578063bb9f408d146107da57610514565b8063acaf88cd146107ac578063b3ab15fb146107b457610514565b8063948a03f211610270578063a3639b3911610255578063a3639b3914610789578063a5834e061461079c578063aaada5da146107a457610514565b8063948a03f21461076e57806395dc14e11461077657610514565b8063877887821461073857806387d4bda9146107405780638da5cb5b1461075357806393c99e6a1461075b57610514565b806348a0d754116103e25780636d4710b91161035a578063731ff24a1161030e57806378b4330f116102f357806378b4330f146105d05780638456cb5914610728578063853828b61461073057610514565b8063731ff24a1461071857806377c7b8fc1461072057610514565b806370897b231161033f57806370897b23146106f5578063715018a614610708578063722713f71461071057610514565b80636d4710b9146106da578063704b6c02146106e257610514565b8063570ca735116103b15780635c975abb116103965780635c975abb146106b757806361d027b3146106bf578063668679ba146106c757610514565b8063570ca735146106a757806358ebceb6146106af57610514565b806348a0d754146106715780634e4de1e9146106795780634f1bfc9e1461068c5780635521e9bf1461069457610514565b80632cfc5f01116104905780633a98ef39116104445780633f4ba83a116104295780633f4ba83a146106365780633fec4e321461063e578063423b93ed1461065e57610514565b80633a98ef39146106265780633eb788741461062e57610514565b80632e1a7d4d116104755780632e1a7d4d146105ed5780632f6c493c14610600578063359819211461061357610514565b80632cfc5f01146105d05780632d19b982146105d857610514565b80631959a002116104e75780631ea30fef116104cc5780631ea30fef146105a25780631efac1b8146105aa57806329a5cfd6146105bd57610514565b80631959a0021461056757806319ab453c1461058f57610514565b806301e813261461051957806305a9f274146105375780630c59696b1461053f57806314ff303914610554575b600080fd5b610521610913565b60405161052e9190614ee7565b60405180910390f35b61052161091b565b61055261054d3660046144b9565b610921565b005b6105526105623660046144b9565b6109fa565b61057a61057536600461441e565b610abc565b60405161052e99989796959493929190614f3d565b61055261059d36600461441e565b610b0e565b610521610e1d565b6105526105b83660046144b9565b610e27565b6105216105cb366004614472565b610eea565b610521611039565b6105e0611040565b60405161052e9190614526565b6105526105fb3660046144b9565b61105c565b61055261060e36600461441e565b6110e3565b6105526106213660046144b9565b611217565b6105216112d7565b6105216112dd565b6105526112e3565b61065161064c36600461441e565b6113a5565b60405161052e9190614607565b61055261066c36600461443a565b6113ba565b6105216114d7565b61055261068736600461443a565b6115a1565b6105216116be565b6105526106a23660046144b9565b6116c4565b6105e061174d565b610521611769565b61065161181f565b6105e0611840565b6106516106d536600461441e565b61185c565b610521611871565b6105526106f036600461441e565b611877565b6105526107033660046144b9565b6119ad565b610552611a6f565b610521611aee565b610521611ba6565b610521611bb6565b610552611c0d565b610552611cd0565b610521611ce9565b61065161074e36600461441e565b611cef565b6105e0611d04565b6105526107693660046144b9565b611d20565b610521611de6565b61052161078436600461441e565b611def565b6105526107973660046144b9565b611f60565b610521612020565b610521612026565b61052161202c565b6105526107c236600461441e565b612032565b6105526107d53660046144b9565b612168565b6105526107e83660046144b9565b61222a565b6105216122ec565b6105216122f2565b61055261080b36600461443a565b6122f8565b61055261081e3660046144b9565b612415565b61052161083136600461441e565b6124d5565b6105e06125d7565b6105216125fb565b61055261085436600461441e565b612601565b61055261086736600461441e565b612715565b61055261087a36600461441e565b6128b8565b6105216129cc565b6105e06129d2565b61055261089d3660046144e9565b6129ee565b610521612a7d565b610521612a83565b610521612a89565b610521612a8f565b6105526108d036600461441e565b612a95565b6105526108e336600461441e565b612bcb565b6105526108f63660046144b9565b612c94565b6105e0612d58565b6105e0612d74565b610521612d98565b6305265c0081565b600d5481565b60085473ffffffffffffffffffffffffffffffffffffffff16331461097b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b60405180910390fd5b64e8d4a510008111156109ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614aee565b60178190556040517ff4bd1c5978320077e792afbb3911e8cab1325ce28a6b3e67f9067a1d80692961906109ef908390614ee7565b60405180910390a150565b60085473ffffffffffffffffffffffffffffffffffffffff163314610a4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b6101f4811115610a87576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614df6565b60168190556040517fcab352e118188b8a2f20a2e8c4ce1241ce2c1740aac4f17c5b0831e65824d8ef906109ef908390614ee7565b6003602081905260009182526040909120805460018201546002830154938301546004840154600585015460068601546007870154600890970154959794969495939492939192909160ff9091169089565b610b16612da2565b73ffffffffffffffffffffffffffffffffffffffff16610b34611d04565b73ffffffffffffffffffffffffffffffffffffffff1614610b81576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c05565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190610bd6903390600401614526565b60206040518083038186803b158015610bee57600080fd5b505afa158015610c02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2691906144d1565b905080610c5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614830565b610c8173ffffffffffffffffffffffffffffffffffffffff8316333084612da6565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063095ea7b390610cf5907f0000000000000000000000000000000000000000000000000000000000000000908590600401614578565b602060405180830381600087803b158015610d0f57600080fd5b505af1158015610d23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d47919061449d565b50600b546040517fe2bbb15800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163e2bbb15891610dbe91908590600401614f14565b600060405180830381600087803b158015610dd857600080fd5b505af1158015610dec573d6000803e3d6000fd5b50506040517f57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14925060009150a15050565b6509184e72a00081565b60085473ffffffffffffffffffffffffffffffffffffffff163314610e78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b62093a80811115610eb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614d3c565b60188190556040517fb89ddaddb7435be26824cb48d2d0186c9525a2e1ec057abcb502704cdc0686cc906109ef908390614ee7565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604081208054831115610f1d57805492505b3360009081526005602052604090205460ff16158015610f4d57506018548160010154610f4a9190614f7e565b42105b1561102d576000610f5c611769565b610f64611aee565b610f6e9190614f7e565b8254909150600090610f8564e8d4a5100087614fcf565b610f8f9190614f96565b90506000610f9c87612e4f565b60068501546007548654610fb09087614fcf565b610fba9190614f96565b610fc4919061500c565b610fce919061500c565b9050600064e8d4a51000610fe28484614fcf565b610fec9190614f96565b601554909150610ffb33612e6d565b1561100557506016545b60006127106110148385614fcf565b61101e9190614f96565b97506110339650505050505050565b60009150505b92915050565b62093a8081565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b61106461181f565b1561109b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a23565b600081116110d5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610972906146d1565b6110e0816000612e73565b50565b803373ffffffffffffffffffffffffffffffffffffffff8216148061111f5750600a5473ffffffffffffffffffffffffffffffffffffffff1633145b611155576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c71565b61115d61181f565b15611194576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a23565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020600781015460ff1680156111d05750428160050154105b611206576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614eb0565b61121260008085613246565b505050565b60085473ffffffffffffffffffffffffffffffffffffffff163314611268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b600081116112a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a91565b60118190556040517f18b6d179114082d7eda9837e15a39eb30032d5f3df00487a67541398f48fabfe906109ef908390614ee7565b60075481565b60145481565b60085473ffffffffffffffffffffffffffffffffffffffff163314611334576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b61133c61181f565b611372576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109729061469a565b61137a6138bc565b6040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b60046020526000908152604090205460ff1681565b60085473ffffffffffffffffffffffffffffffffffffffff16331461140b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b73ffffffffffffffffffffffffffffffffffffffff8216611458576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b73ffffffffffffffffffffffffffffffffffffffff821660008181526004602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917f3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d4491a35050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a082319061154c903090600401614526565b60206040518083038186803b15801561156457600080fd5b505afa158015611578573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159c91906144d1565b905090565b60085473ffffffffffffffffffffffffffffffffffffffff1633146115f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b73ffffffffffffffffffffffffffffffffffffffff821661163f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b73ffffffffffffffffffffffffffffffffffffffff821660008181526006602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917f3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d4491a35050565b600f5481565b6116cc61181f565b15611703576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a23565b6509184e72a0008111611742576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614ca8565b6110e0600082612e73565b600a5473ffffffffffffffffffffffffffffffffffffffff1681565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631175a1dd600b54306040518363ffffffff1660e01b81526004016117c9929190614ef0565b60206040518083038186803b1580156117e157600080fd5b505afa1580156117f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061181991906144d1565b91505090565b60005474010000000000000000000000000000000000000000900460ff1690565b60095473ffffffffffffffffffffffffffffffffffffffff1681565b60066020526000908152604090205460ff1681565b600b5481565b61187f612da2565b73ffffffffffffffffffffffffffffffffffffffff1661189d611d04565b73ffffffffffffffffffffffffffffffffffffffff16146118ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c05565b73ffffffffffffffffffffffffffffffffffffffff8116611937576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117918290556040517f71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c926109ef921690614526565b60085473ffffffffffffffffffffffffffffffffffffffff1633146119fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b6107d0811115611a3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610972906149c6565b60138190556040517fefeafcf03e479a9566d7ef321b4816de0ba19cfa3cd0fae2f8c5f4a0afb342c4906109ef908390614ee7565b611a77612da2565b73ffffffffffffffffffffffffffffffffffffffff16611a95611d04565b73ffffffffffffffffffffffffffffffffffffffff1614611ae2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c05565b611aec6000613962565b565b6000600c547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401611b4c9190614526565b60206040518083038186803b158015611b6457600080fd5b505afa158015611b78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b9c91906144d1565b61159c9190614f7e565b6b204fce5e3e2502611000000081565b6000600754600014611c0057600754611bcd611769565b611bd5611aee565b611bdf9190614f7e565b611bf190670de0b6b3a7640000614fcf565b611bfb9190614f96565b61159c565b50670de0b6b3a764000090565b60085473ffffffffffffffffffffffffffffffffffffffff163314611c5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b611c6661181f565b15611c9d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a23565b611ca56139d7565b6040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b33600090815260036020526040902054611aec9061105c565b60135481565b60056020526000908152604090205460ff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60085473ffffffffffffffffffffffffffffffffffffffff163314611d71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b652d79883d2000811115611db1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610972906147d3565b60128190556040517f7666dfff8c3377938e522b4eed3aff079973a976f95969db60a406d49f40da4e906109ef908390614ee7565b64e8d4a5100081565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600360205260408120805415801590611e285750600781015460ff165b8015611e5a575073ffffffffffffffffffffffffffffffffffffffff831660009081526006602052604090205460ff16155b8015611e76575042600e548260050154611e749190614f7e565b105b15611f55576000611e85611769565b611e8d611aee565b611e979190614f7e565b60068301546007548454929350600092611eb19085614fcf565b611ebb9190614f96565b611ec5919061500c565b90506000836008015482611ed9919061500c565b90506000600e54856005015442611ef0919061500c565b611efa919061500c565b9050601154811115611f0b57506011545b600060115460175483611f1e9190614fcf565b611f289190614f96565b9050600064e8d4a51000611f3c8386614fcf565b611f469190614f96565b9750611f5b9650505050505050565b60009150505b919050565b60085473ffffffffffffffffffffffffffffffffffffffff163314611fb1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b60008111611feb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614765565b60108190556040517f9478eb023aac0a7d58a4e935377056bf27cf5b72a2300725f831817a8f62fbde906109ef908390614ee7565b60175481565b600e5481565b60115481565b61203a612da2565b73ffffffffffffffffffffffffffffffffffffffff16612058611d04565b73ffffffffffffffffffffffffffffffffffffffff16146120a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c05565b73ffffffffffffffffffffffffffffffffffffffff81166120f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b600a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117918290556040517fda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a926109ef921690614526565b60085473ffffffffffffffffffffffffffffffffffffffff1633146121b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b6101f48111156121f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614df6565b60158190556040517fd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d906109ef908390614ee7565b60085473ffffffffffffffffffffffffffffffffffffffff16331461227b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b6107d08111156122b7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610972906149c6565b60148190556040517fc5d25457b67b87678c987375af13f6e50beb3ad7bfd009da26766ae986eaa20d906109ef908390614ee7565b60125481565b6107d081565b60085473ffffffffffffffffffffffffffffffffffffffff163314612349576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b73ffffffffffffffffffffffffffffffffffffffff8216612396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b73ffffffffffffffffffffffffffffffffffffffff821660008181526005602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917f3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d4491a35050565b60085473ffffffffffffffffffffffffffffffffffffffff163314612466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b600081116124a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614e53565b600e8190556040517ff84bf2b901cfc02956d4e69556d7448cef4ea13587e7714dba7c6d697091e7ad906109ef908390614ee7565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260036020526040812080541580159061250f5750600781015460ff16155b8015612541575073ffffffffffffffffffffffffffffffffffffffff831660009081526004602052604090205460ff16155b15611f55576000612550611769565b612558611aee565b6125629190614f7e565b905060006007548284600001546125799190614fcf565b6125839190614f96565b90506000836002015482612597919061500c565b6013549091506125a687612e6d565b156125b057506014545b60006127106125bf8385614fcf565b6125c99190614f96565b9650611f5b95505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6101f481565b60085473ffffffffffffffffffffffffffffffffffffffff163314612652576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b73ffffffffffffffffffffffffffffffffffffffff811661269f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117918290556040517f5352e27b0414343d9438a1c6e9d04c65c7cb4d91f44920afee588f91717893f1926109ef921690614526565b60085473ffffffffffffffffffffffffffffffffffffffff163314612766576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156127ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614867565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190612841903090600401614526565b60206040518083038186803b15801561285957600080fd5b505afa15801561286d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061289191906144d1565b90506128b473ffffffffffffffffffffffffffffffffffffffff83163383613a7e565b5050565b60085473ffffffffffffffffffffffffffffffffffffffff163314612909576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b73ffffffffffffffffffffffffffffffffffffffff8116612956576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117918290556040517f8f49a182922022d9119a1a6aeeca151b4a5665e86bd61c1ff32e152d459558b2926109ef921690614526565b60185481565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6129f661181f565b15612a2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a23565b6000821180612a3c5750600081115b612a72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614663565b6128b4828233613246565b60105481565b60165481565b600c5481565b60155481565b612a9d612da2565b73ffffffffffffffffffffffffffffffffffffffff16612abb611d04565b73ffffffffffffffffffffffffffffffffffffffff1614612b08576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c05565b73ffffffffffffffffffffffffffffffffffffffff8116612b55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a5a565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117918290556040517fafa147634b29e2c7bd53ce194256b9f41cfb9ba3036f2b822fdd1d965beea086926109ef921690614526565b612bd3612da2565b73ffffffffffffffffffffffffffffffffffffffff16612bf1611d04565b73ffffffffffffffffffffffffffffffffffffffff1614612c3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c05565b73ffffffffffffffffffffffffffffffffffffffff8116612c8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614708565b6110e081613962565b60085473ffffffffffffffffffffffffffffffffffffffff163314612ce5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614c3a565b6305265c00811115612d23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614ba8565b600f8190556040517fcab2f3455b51b6ca5377e84fccd0f890b6f6ca36c02e18b6d36cb34f469fe4fe906109ef908390614ee7565b60085473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b652d79883d200081565b3390565b612e49846323b872dd60e01b858585604051602401612dc793929190614547565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613a9d565b50505050565b6000612e5a82611def565b612e63836124d5565b6110339190614f7e565b3b151590565b3360009081526003602052604090208054831115612ebd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610972906148c4565b42816005015410612efa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109729061479c565b60025473ffffffffffffffffffffffffffffffffffffffff1615612fa1576002546040517f51cff8d900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906351cff8d990612f6e903390600401614526565b600060405180830381600087803b158015612f8857600080fd5b505af1158015612f9c573d6000803e3d6000fd5b505050505b80548390600090612fbe6b204fce5e3e2502611000000084614fcf565b612fc89190614f96565b9050612fd2613b53565b612fdb33613d2a565b84158015612fe95750600084115b1561302a576000612ff8611aee565b905080600754866130099190614fcf565b6130139190614f96565b845490935083111561302457835492505b50613051565b82546b204fce5e3e25026110000000906130449083614fcf565b61304e9190614f96565b91505b60006007548361305f611aee565b6130699190614fcf565b6130739190614f96565b905082846000016000828254613089919061500c565b9250508190555082600760008282546130a2919061500c565b90915550503360009081526005602052604090205460ff161580156130d7575060185484600101546130d49190614f7e565b42105b15613164576015546130e833612e6d565b156130f257506016545b60006127106131018385614fcf565b61310b9190614f96565b6009549091506131559073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911683613a7e565b61315f818461500c565b925050505b6131a573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163383613a7e565b8354156131d7576007546131b7611aee565b85546131c39190614fcf565b6131cd9190614f96565b60028501556131df565b600060028501555b4260038501556131ee3361419b565b3373ffffffffffffffffffffffffffffffffffffffff167ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5688285604051613236929190614f14565b60405180910390a2505050505050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600360205260409020805415806132795750600084115b156132bd576509184e72a00084116132bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614b4b565b60058101548390421161331b5784156132fa574260048301556008820154600d80546000906132ed90849061500c565b9091555050600060088301555b8160040154826005015461330e919061500c565b6133189082614f7e565b90505b83158061332b575062093a808110155b613361576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614932565b600f5481111561339d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610972906148fb565b60025473ffffffffffffffffffffffffffffffffffffffff1615613448576002546040517f0efe6a8b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690630efe6a8b906134159086908990899060040161459e565b600060405180830381600087803b15801561342f57600080fd5b505af1158015613443573d6000803e3d6000fd5b505050505b613450613b53565b6007546134ad5760006134616114d7565b6009549091506134ab9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911683613a7e565b505b6134b683613d2a565b831561352d5742826005015410156134e55742600483018190556134db908590614f7e565b60058301556134ff565b838260050160008282546134f99190614f7e565b90915550505b6007820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b60008060008061353b611aee565b905088156135895761358573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001688308c612da6565b8892505b85541580159061359d5750600786015460ff165b156136165760075486546135b19083614fcf565b6135bb9190614f96565b91506135c78284614f7e565b92508560000154600760008282546135df919061500c565b90915550506000865560048601544214156136165760088601829055600d8054839190600090613610908490614f7e565b90915550505b6007541561364657613628828261500c565b6007546136359085614fcf565b61363f9190614f96565b935061364a565b8293505b8560040154866005015411156137cd57600060105460125488600401548960050154613676919061500c565b6136809190614fcf565b61368a9190614f96565b9050600064e8d4a5100061369e8784614fcf565b6136a89190614f96565b90506136b48187614f7e565b9550858860000160008282546136ca9190614f7e565b909155506000905064e8d4a510006136e28785614fcf565b6136ec9190614f96565b9050808960060160008282546137029190614f7e565b9250508190555080600c600082825461371b9190614f7e565b925050819055508b8960080160008282546137369190614f7e565b925050819055508b600d600082825461374f9190614f7e565b925050819055508973ffffffffffffffffffffffffffffffffffffffff167f2b943276e5d747f6f7dd46d3b880d8874cb8d6b9b88ca1903990a2738e7dc7a18a600801548b600001548c600401548d600501546137ac919061500c565b426040516137bd9493929190614f22565b60405180910390a25050506137e7565b838660000160008282546137e19190614f7e565b90915550505b60008911806137f65750600088115b15613802574260018701555b83600760008282546138149190614f7e565b90915550506006860154600754613829611aee565b88546138359190614fcf565b61383f9190614f96565b613849919061500c565b600287015542600387015561385d8761419b565b8673ffffffffffffffffffffffffffffffffffffffff167f7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f98a868b426040516138a99493929190614f22565b60405180910390a2505050505050505050565b6138c461181f565b6138fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109729061469a565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61394b612da2565b6040516139589190614526565b60405180910390a1565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6139df61181f565b15613a16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614a23565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861394b612da2565b6112128363a9059cbb60e01b8484604051602401612dc7929190614578565b6000613aff826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166142959092919063ffffffff16565b8051909150156112125780806020019051810190613b1d919061449d565b611212576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614d99565b600b546040517f1175a1dd00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691631175a1dd91613bcb913090600401614ef0565b60206040518083038186803b158015613be357600080fd5b505afa158015613bf7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c1b91906144d1565b905080156110e0576000613c2d6114d7565b600b546040517f441a3e7000000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163441a3e7091613ca691600090600401614f14565b600060405180830381600087803b158015613cc057600080fd5b505af1158015613cd4573d6000803e3d6000fd5b505050506000613ce26114d7565b9050337fc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba613d10848461500c565b604051613d1d9190614ee7565b60405180910390a2505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526003602052604090208054156128b457600781015460ff161561402257600081600601546007548360000154613d7a611aee565b613d849190614fcf565b613d8e9190614f96565b613d98919061500c565b90508160060154600c6000828254613db0919061500c565b909155505060006006830181905582546007805491929091613dd390849061500c565b909155505073ffffffffffffffffffffffffffffffffffffffff831660009081526006602052604090205460ff16158015613e1e575042600e548360050154613e1c9190614f7e565b105b15613efd576000826008015482613e35919061500c565b90506000600e54846005015442613e4c919061500c565b613e56919061500c565b9050601154811115613e6757506011545b600060115460175483613e7a9190614fcf565b613e849190614f96565b9050600064e8d4a51000613e988386614fcf565b613ea29190614f96565b600954909150613eec9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911683613a7e565b613ef6818661500c565b9450505050505b6000613f07611aee565b90506000600754600014613f3d57613f1f838361500c565b600754613f2c9085614fcf565b613f369190614f96565b9050613f40565b50815b80845560078054829190600090613f58908490614f7e565b9091555050600584015442111561401a576007840180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055600060048501819055600585018190556008850154600d805491929091613fbb90849061500c565b90915550506000600885015560405173ffffffffffffffffffffffffffffffffffffffff8616907ff7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004906140119086904290614f14565b60405180910390a25b5050506128b4565b73ffffffffffffffffffffffffffffffffffffffff821660009081526004602052604090205460ff166128b457600060075461405c611aee565b83546140689190614fcf565b6140729190614f96565b905081600001546007600082825461408a919061500c565b9091555050600080835560028301546140a3908361500c565b6013549091506140b285612e6d565b156140bc57506014545b60006127106140cb8385614fcf565b6140d59190614f96565b90508015614131576009546141249073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911683613a7e565b61412e818561500c565b93505b600061413b611aee565b9050600060075460001461417157614153868361500c565b6007546141609088614fcf565b61416a9190614f96565b9050614174565b50845b8087556007805482919060009061418c908490614f7e565b90915550505050505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff16156110e05773ffffffffffffffffffffffffffffffffffffffff81166000908152600360205260408120600481015460058201549192916141f5919061500c565b6001546008840154600d546010546040517fe874fdaf00000000000000000000000000000000000000000000000000000000815294955073ffffffffffffffffffffffffffffffffffffffff9093169363e874fdaf9361425e93899390928892906004016145cc565b600060405180830381600087803b15801561427857600080fd5b505af115801561428c573d6000803e3d6000fd5b50505050505050565b60606142a484846000856142ae565b90505b9392505050565b6060824710156142ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614969565b6142f3856143af565b614329576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290614d05565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051614352919061450a565b60006040518083038185875af1925050503d806000811461438f576040519150601f19603f3d011682016040523d82523d6000602084013e614394565b606091505b50915091506143a48282866143cb565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b606083156143da5750816142a7565b8251156143ea5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109729190614612565b60006020828403121561442f578081fd5b81356142a78161507e565b6000806040838503121561444c578081fd5b82356144578161507e565b91506020830135614467816150a0565b809150509250929050565b60008060408385031215614484578182fd5b823561448f8161507e565b946020939093013593505050565b6000602082840312156144ae578081fd5b81516142a7816150a0565b6000602082840312156144ca578081fd5b5035919050565b6000602082840312156144e2578081fd5b5051919050565b600080604083850312156144fb578182fd5b50508035926020909101359150565b6000825161451c818460208701615023565b9190910192915050565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b73ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152604081019190915260600190565b73ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b73ffffffffffffffffffffffffffffffffffffffff9390931683526020830191909152604082015260600190565b73ffffffffffffffffffffffffffffffffffffffff959095168552602085019390935260408401919091526060830152608082015260a00190565b901515815260200190565b6000602082528251806020840152614631816040850160208701615023565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60208082526012908201527f4e6f7468696e6720746f206465706f7369740000000000000000000000000000604082015260600190565b60208082526014908201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604082015260600190565b60208082526013908201527f4e6f7468696e6720746f20776974686472617700000000000000000000000000604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252601e908201527f4455524154494f4e5f464143544f522063616e6e6f74206265207a65726f0000604082015260600190565b6020808252600d908201527f5374696c6c20696e206c6f636b00000000000000000000000000000000000000604082015260600190565b60208082526033908201527f424f4f53545f5745494748542063616e6e6f74206265206d6f7265207468616e60408201527f20424f4f53545f5745494748545f4c494d495400000000000000000000000000606082015260800190565b60208082526015908201527f42616c616e6365206d7573742065786365656420300000000000000000000000604082015260600190565b60208082526025908201527f546f6b656e2063616e6e6f742062652073616d65206173206465706f7369742060408201527f746f6b656e000000000000000000000000000000000000000000000000000000606082015260800190565b6020808252601f908201527f576974686472617720616d6f756e7420657863656564732062616c616e636500604082015260600190565b6020808252601c908201527f4d6178696d756d206c6f636b20706572696f6420657863656564656400000000604082015260600190565b6020808252601f908201527f4d696e696d756d206c6f636b20706572696f64206973206f6e65207765656b00604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60408201527f722063616c6c0000000000000000000000000000000000000000000000000000606082015260800190565b60208082526036908201527f706572666f726d616e63654665652063616e6e6f74206265206d6f726520746860408201527f616e204d41585f504552464f524d414e43455f46454500000000000000000000606082015260800190565b60208082526010908201527f5061757361626c653a2070617573656400000000000000000000000000000000604082015260600190565b60208082526016908201527f43616e6e6f74206265207a65726f206164647265737300000000000000000000604082015260600190565b60208082526026908201527f4455524154494f4e5f464143544f525f4f5645524455452063616e6e6f74206260408201527f65207a65726f0000000000000000000000000000000000000000000000000000606082015260800190565b6020808252602e908201527f6f7665726475654665652063616e6e6f74206265206d6f7265207468616e204d60408201527f41585f4f5645524455455f464545000000000000000000000000000000000000606082015260800190565b60208082526036908201527f4465706f73697420616d6f756e74206d7573742062652067726561746572207460408201527f68616e204d494e5f4445504f5349545f414d4f554e5400000000000000000000606082015260800190565b6020808252603d908201527f4d41585f4c4f434b5f4455524154494f4e2063616e6e6f74206265206d6f726560408201527f207468616e204d41585f4c4f434b5f4455524154494f4e5f4c494d4954000000606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252600b908201527f61646d696e3a207775743f000000000000000000000000000000000000000000604082015260600190565b6020808252601a908201527f4e6f74206f70657261746f72206f722063616b65206f776e6572000000000000604082015260600190565b60208082526038908201527f576974686472617720616d6f756e74206d75737420626520677265617465722060408201527f7468616e204d494e5f57495448445241575f414d4f554e540000000000000000606082015260800190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252603d908201527f7769746864726177466565506572696f642063616e6e6f74206265206d6f726560408201527f207468616e204d41585f57495448445241575f4645455f504552494f44000000606082015260800190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60408201527f6f74207375636365656400000000000000000000000000000000000000000000606082015260800190565b60208082526030908201527f77697468647261774665652063616e6e6f74206265206d6f7265207468616e2060408201527f4d41585f57495448445241575f46454500000000000000000000000000000000606082015260800190565b60208082526023908201527f554e4c4f434b5f465245455f4455524154494f4e2063616e6e6f74206265207a60408201527f65726f0000000000000000000000000000000000000000000000000000000000606082015260800190565b60208082526011908201527f43616e6e6f7420756e6c6f636b20796574000000000000000000000000000000604082015260600190565b90815260200190565b91825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b918252602082015260400190565b93845260208401929092526040830152606082015260800190565b988952602089019790975260408801959095526060870193909352608086019190915260a085015260c0840152151560e08301526101008201526101200190565b60008219821115614f9157614f9161504f565b500190565b600082614fca577f4e487b710000000000000000000000000000000000000000000000000000000081526012600452602481fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156150075761500761504f565b500290565b60008282101561501e5761501e61504f565b500390565b60005b8381101561503e578181015183820152602001615026565b83811115612e495750506000910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff811681146110e057600080fd5b80151581146110e057600080fdfea2646970667358221220161b312fa06b5510bce35751461f7d7626cd70db69df524969ca0f6c75d7dc4364736f6c634300080100330000000000000000000000000e09fabb73bd3ade0a17ecc321fd13a19e81ce82000000000000000000000000a5f8c5dbd5f286960b9d90548680ae5ebff076520000000000000000000000007a2c5c265bdc9724dace715c5ff60eea40e07f47000000000000000000000000ecc90d54b10add1ab746abe7e83abe178b72aa9e000000000000000000000000ecc90d54b10add1ab746abe7e83abe178b72aa9e0000000000000000000000000000000000000000000000000000000000000000",
}

// CakePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use CakePoolMetaData.ABI instead.
var CakePoolABI = CakePoolMetaData.ABI

// CakePoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CakePoolMetaData.Bin instead.
var CakePoolBin = CakePoolMetaData.Bin

// DeployCakePool deploys a new Ethereum contract, binding an instance of CakePool to it.
func DeployCakePool(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _masterchefV2 common.Address, _admin common.Address, _treasury common.Address, _operator common.Address, _pid *big.Int) (common.Address, *types.Transaction, *CakePool, error) {
	parsed, err := CakePoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CakePoolBin), backend, _token, _masterchefV2, _admin, _treasury, _operator, _pid)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CakePool{CakePoolCaller: CakePoolCaller{contract: contract}, CakePoolTransactor: CakePoolTransactor{contract: contract}, CakePoolFilterer: CakePoolFilterer{contract: contract}}, nil
}

// CakePool is an auto generated Go binding around an Ethereum contract.
type CakePool struct {
	CakePoolCaller     // Read-only binding to the contract
	CakePoolTransactor // Write-only binding to the contract
	CakePoolFilterer   // Log filterer for contract events
}

// CakePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type CakePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CakePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CakePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CakePoolSession struct {
	Contract     *CakePool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CakePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CakePoolCallerSession struct {
	Contract *CakePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CakePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CakePoolTransactorSession struct {
	Contract     *CakePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CakePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type CakePoolRaw struct {
	Contract *CakePool // Generic contract binding to access the raw methods on
}

// CakePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CakePoolCallerRaw struct {
	Contract *CakePoolCaller // Generic read-only contract binding to access the raw methods on
}

// CakePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CakePoolTransactorRaw struct {
	Contract *CakePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCakePool creates a new instance of CakePool, bound to a specific deployed contract.
func NewCakePool(address common.Address, backend bind.ContractBackend) (*CakePool, error) {
	contract, err := bindCakePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CakePool{CakePoolCaller: CakePoolCaller{contract: contract}, CakePoolTransactor: CakePoolTransactor{contract: contract}, CakePoolFilterer: CakePoolFilterer{contract: contract}}, nil
}

// NewCakePoolCaller creates a new read-only instance of CakePool, bound to a specific deployed contract.
func NewCakePoolCaller(address common.Address, caller bind.ContractCaller) (*CakePoolCaller, error) {
	contract, err := bindCakePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CakePoolCaller{contract: contract}, nil
}

// NewCakePoolTransactor creates a new write-only instance of CakePool, bound to a specific deployed contract.
func NewCakePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CakePoolTransactor, error) {
	contract, err := bindCakePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CakePoolTransactor{contract: contract}, nil
}

// NewCakePoolFilterer creates a new log filterer instance of CakePool, bound to a specific deployed contract.
func NewCakePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CakePoolFilterer, error) {
	contract, err := bindCakePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CakePoolFilterer{contract: contract}, nil
}

// bindCakePool binds a generic wrapper to an already deployed contract.
func bindCakePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CakePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CakePool *CakePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CakePool.Contract.CakePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CakePool *CakePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CakePool.Contract.CakePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CakePool *CakePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CakePool.Contract.CakePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CakePool *CakePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CakePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CakePool *CakePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CakePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CakePool *CakePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CakePool.Contract.contract.Transact(opts, method, params...)
}

// BOOSTWEIGHT is a free data retrieval call binding the contract method 0xbc75f4b8.
//
// Solidity: function BOOST_WEIGHT() view returns(uint256)
func (_CakePool *CakePoolCaller) BOOSTWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "BOOST_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTWEIGHT is a free data retrieval call binding the contract method 0xbc75f4b8.
//
// Solidity: function BOOST_WEIGHT() view returns(uint256)
func (_CakePool *CakePoolSession) BOOSTWEIGHT() (*big.Int, error) {
	return _CakePool.Contract.BOOSTWEIGHT(&_CakePool.CallOpts)
}

// BOOSTWEIGHT is a free data retrieval call binding the contract method 0xbc75f4b8.
//
// Solidity: function BOOST_WEIGHT() view returns(uint256)
func (_CakePool *CakePoolCallerSession) BOOSTWEIGHT() (*big.Int, error) {
	return _CakePool.Contract.BOOSTWEIGHT(&_CakePool.CallOpts)
}

// BOOSTWEIGHTLIMIT is a free data retrieval call binding the contract method 0xfd253b64.
//
// Solidity: function BOOST_WEIGHT_LIMIT() view returns(uint256)
func (_CakePool *CakePoolCaller) BOOSTWEIGHTLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "BOOST_WEIGHT_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTWEIGHTLIMIT is a free data retrieval call binding the contract method 0xfd253b64.
//
// Solidity: function BOOST_WEIGHT_LIMIT() view returns(uint256)
func (_CakePool *CakePoolSession) BOOSTWEIGHTLIMIT() (*big.Int, error) {
	return _CakePool.Contract.BOOSTWEIGHTLIMIT(&_CakePool.CallOpts)
}

// BOOSTWEIGHTLIMIT is a free data retrieval call binding the contract method 0xfd253b64.
//
// Solidity: function BOOST_WEIGHT_LIMIT() view returns(uint256)
func (_CakePool *CakePoolCallerSession) BOOSTWEIGHTLIMIT() (*big.Int, error) {
	return _CakePool.Contract.BOOSTWEIGHTLIMIT(&_CakePool.CallOpts)
}

// DURATIONFACTOR is a free data retrieval call binding the contract method 0xe464c623.
//
// Solidity: function DURATION_FACTOR() view returns(uint256)
func (_CakePool *CakePoolCaller) DURATIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "DURATION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATIONFACTOR is a free data retrieval call binding the contract method 0xe464c623.
//
// Solidity: function DURATION_FACTOR() view returns(uint256)
func (_CakePool *CakePoolSession) DURATIONFACTOR() (*big.Int, error) {
	return _CakePool.Contract.DURATIONFACTOR(&_CakePool.CallOpts)
}

// DURATIONFACTOR is a free data retrieval call binding the contract method 0xe464c623.
//
// Solidity: function DURATION_FACTOR() view returns(uint256)
func (_CakePool *CakePoolCallerSession) DURATIONFACTOR() (*big.Int, error) {
	return _CakePool.Contract.DURATIONFACTOR(&_CakePool.CallOpts)
}

// DURATIONFACTOROVERDUE is a free data retrieval call binding the contract method 0xacaf88cd.
//
// Solidity: function DURATION_FACTOR_OVERDUE() view returns(uint256)
func (_CakePool *CakePoolCaller) DURATIONFACTOROVERDUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "DURATION_FACTOR_OVERDUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATIONFACTOROVERDUE is a free data retrieval call binding the contract method 0xacaf88cd.
//
// Solidity: function DURATION_FACTOR_OVERDUE() view returns(uint256)
func (_CakePool *CakePoolSession) DURATIONFACTOROVERDUE() (*big.Int, error) {
	return _CakePool.Contract.DURATIONFACTOROVERDUE(&_CakePool.CallOpts)
}

// DURATIONFACTOROVERDUE is a free data retrieval call binding the contract method 0xacaf88cd.
//
// Solidity: function DURATION_FACTOR_OVERDUE() view returns(uint256)
func (_CakePool *CakePoolCallerSession) DURATIONFACTOROVERDUE() (*big.Int, error) {
	return _CakePool.Contract.DURATIONFACTOROVERDUE(&_CakePool.CallOpts)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_CakePool *CakePoolCaller) MAXLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MAX_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_CakePool *CakePoolSession) MAXLOCKDURATION() (*big.Int, error) {
	return _CakePool.Contract.MAXLOCKDURATION(&_CakePool.CallOpts)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MAXLOCKDURATION() (*big.Int, error) {
	return _CakePool.Contract.MAXLOCKDURATION(&_CakePool.CallOpts)
}

// MAXLOCKDURATIONLIMIT is a free data retrieval call binding the contract method 0x01e81326.
//
// Solidity: function MAX_LOCK_DURATION_LIMIT() view returns(uint256)
func (_CakePool *CakePoolCaller) MAXLOCKDURATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MAX_LOCK_DURATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCKDURATIONLIMIT is a free data retrieval call binding the contract method 0x01e81326.
//
// Solidity: function MAX_LOCK_DURATION_LIMIT() view returns(uint256)
func (_CakePool *CakePoolSession) MAXLOCKDURATIONLIMIT() (*big.Int, error) {
	return _CakePool.Contract.MAXLOCKDURATIONLIMIT(&_CakePool.CallOpts)
}

// MAXLOCKDURATIONLIMIT is a free data retrieval call binding the contract method 0x01e81326.
//
// Solidity: function MAX_LOCK_DURATION_LIMIT() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MAXLOCKDURATIONLIMIT() (*big.Int, error) {
	return _CakePool.Contract.MAXLOCKDURATIONLIMIT(&_CakePool.CallOpts)
}

// MAXOVERDUEFEE is a free data retrieval call binding the contract method 0x948a03f2.
//
// Solidity: function MAX_OVERDUE_FEE() view returns(uint256)
func (_CakePool *CakePoolCaller) MAXOVERDUEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MAX_OVERDUE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOVERDUEFEE is a free data retrieval call binding the contract method 0x948a03f2.
//
// Solidity: function MAX_OVERDUE_FEE() view returns(uint256)
func (_CakePool *CakePoolSession) MAXOVERDUEFEE() (*big.Int, error) {
	return _CakePool.Contract.MAXOVERDUEFEE(&_CakePool.CallOpts)
}

// MAXOVERDUEFEE is a free data retrieval call binding the contract method 0x948a03f2.
//
// Solidity: function MAX_OVERDUE_FEE() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MAXOVERDUEFEE() (*big.Int, error) {
	return _CakePool.Contract.MAXOVERDUEFEE(&_CakePool.CallOpts)
}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_CakePool *CakePoolCaller) MAXPERFORMANCEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MAX_PERFORMANCE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_CakePool *CakePoolSession) MAXPERFORMANCEFEE() (*big.Int, error) {
	return _CakePool.Contract.MAXPERFORMANCEFEE(&_CakePool.CallOpts)
}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MAXPERFORMANCEFEE() (*big.Int, error) {
	return _CakePool.Contract.MAXPERFORMANCEFEE(&_CakePool.CallOpts)
}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_CakePool *CakePoolCaller) MAXWITHDRAWFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MAX_WITHDRAW_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_CakePool *CakePoolSession) MAXWITHDRAWFEE() (*big.Int, error) {
	return _CakePool.Contract.MAXWITHDRAWFEE(&_CakePool.CallOpts)
}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MAXWITHDRAWFEE() (*big.Int, error) {
	return _CakePool.Contract.MAXWITHDRAWFEE(&_CakePool.CallOpts)
}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_CakePool *CakePoolCaller) MAXWITHDRAWFEEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MAX_WITHDRAW_FEE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_CakePool *CakePoolSession) MAXWITHDRAWFEEPERIOD() (*big.Int, error) {
	return _CakePool.Contract.MAXWITHDRAWFEEPERIOD(&_CakePool.CallOpts)
}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MAXWITHDRAWFEEPERIOD() (*big.Int, error) {
	return _CakePool.Contract.MAXWITHDRAWFEEPERIOD(&_CakePool.CallOpts)
}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(uint256)
func (_CakePool *CakePoolCaller) MINDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MIN_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(uint256)
func (_CakePool *CakePoolSession) MINDEPOSITAMOUNT() (*big.Int, error) {
	return _CakePool.Contract.MINDEPOSITAMOUNT(&_CakePool.CallOpts)
}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MINDEPOSITAMOUNT() (*big.Int, error) {
	return _CakePool.Contract.MINDEPOSITAMOUNT(&_CakePool.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_CakePool *CakePoolCaller) MINLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MIN_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_CakePool *CakePoolSession) MINLOCKDURATION() (*big.Int, error) {
	return _CakePool.Contract.MINLOCKDURATION(&_CakePool.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MINLOCKDURATION() (*big.Int, error) {
	return _CakePool.Contract.MINLOCKDURATION(&_CakePool.CallOpts)
}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(uint256)
func (_CakePool *CakePoolCaller) MINWITHDRAWAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "MIN_WITHDRAW_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(uint256)
func (_CakePool *CakePoolSession) MINWITHDRAWAMOUNT() (*big.Int, error) {
	return _CakePool.Contract.MINWITHDRAWAMOUNT(&_CakePool.CallOpts)
}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(uint256)
func (_CakePool *CakePoolCallerSession) MINWITHDRAWAMOUNT() (*big.Int, error) {
	return _CakePool.Contract.MINWITHDRAWAMOUNT(&_CakePool.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_CakePool *CakePoolCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_CakePool *CakePoolSession) PRECISIONFACTOR() (*big.Int, error) {
	return _CakePool.Contract.PRECISIONFACTOR(&_CakePool.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_CakePool *CakePoolCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _CakePool.Contract.PRECISIONFACTOR(&_CakePool.CallOpts)
}

// PRECISIONFACTORSHARE is a free data retrieval call binding the contract method 0x731ff24a.
//
// Solidity: function PRECISION_FACTOR_SHARE() view returns(uint256)
func (_CakePool *CakePoolCaller) PRECISIONFACTORSHARE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "PRECISION_FACTOR_SHARE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTORSHARE is a free data retrieval call binding the contract method 0x731ff24a.
//
// Solidity: function PRECISION_FACTOR_SHARE() view returns(uint256)
func (_CakePool *CakePoolSession) PRECISIONFACTORSHARE() (*big.Int, error) {
	return _CakePool.Contract.PRECISIONFACTORSHARE(&_CakePool.CallOpts)
}

// PRECISIONFACTORSHARE is a free data retrieval call binding the contract method 0x731ff24a.
//
// Solidity: function PRECISION_FACTOR_SHARE() view returns(uint256)
func (_CakePool *CakePoolCallerSession) PRECISIONFACTORSHARE() (*big.Int, error) {
	return _CakePool.Contract.PRECISIONFACTORSHARE(&_CakePool.CallOpts)
}

// UNLOCKFREEDURATION is a free data retrieval call binding the contract method 0xaaada5da.
//
// Solidity: function UNLOCK_FREE_DURATION() view returns(uint256)
func (_CakePool *CakePoolCaller) UNLOCKFREEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "UNLOCK_FREE_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNLOCKFREEDURATION is a free data retrieval call binding the contract method 0xaaada5da.
//
// Solidity: function UNLOCK_FREE_DURATION() view returns(uint256)
func (_CakePool *CakePoolSession) UNLOCKFREEDURATION() (*big.Int, error) {
	return _CakePool.Contract.UNLOCKFREEDURATION(&_CakePool.CallOpts)
}

// UNLOCKFREEDURATION is a free data retrieval call binding the contract method 0xaaada5da.
//
// Solidity: function UNLOCK_FREE_DURATION() view returns(uint256)
func (_CakePool *CakePoolCallerSession) UNLOCKFREEDURATION() (*big.Int, error) {
	return _CakePool.Contract.UNLOCKFREEDURATION(&_CakePool.CallOpts)
}

// VCake is a free data retrieval call binding the contract method 0x2d19b982.
//
// Solidity: function VCake() view returns(address)
func (_CakePool *CakePoolCaller) VCake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "VCake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VCake is a free data retrieval call binding the contract method 0x2d19b982.
//
// Solidity: function VCake() view returns(address)
func (_CakePool *CakePoolSession) VCake() (common.Address, error) {
	return _CakePool.Contract.VCake(&_CakePool.CallOpts)
}

// VCake is a free data retrieval call binding the contract method 0x2d19b982.
//
// Solidity: function VCake() view returns(address)
func (_CakePool *CakePoolCallerSession) VCake() (common.Address, error) {
	return _CakePool.Contract.VCake(&_CakePool.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CakePool *CakePoolCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CakePool *CakePoolSession) Admin() (common.Address, error) {
	return _CakePool.Contract.Admin(&_CakePool.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CakePool *CakePoolCallerSession) Admin() (common.Address, error) {
	return _CakePool.Contract.Admin(&_CakePool.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_CakePool *CakePoolCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_CakePool *CakePoolSession) Available() (*big.Int, error) {
	return _CakePool.Contract.Available(&_CakePool.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_CakePool *CakePoolCallerSession) Available() (*big.Int, error) {
	return _CakePool.Contract.Available(&_CakePool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_CakePool *CakePoolCaller) BalanceOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "balanceOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_CakePool *CakePoolSession) BalanceOf() (*big.Int, error) {
	return _CakePool.Contract.BalanceOf(&_CakePool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_CakePool *CakePoolCallerSession) BalanceOf() (*big.Int, error) {
	return _CakePool.Contract.BalanceOf(&_CakePool.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_CakePool *CakePoolCaller) BoostContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "boostContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_CakePool *CakePoolSession) BoostContract() (common.Address, error) {
	return _CakePool.Contract.BoostContract(&_CakePool.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_CakePool *CakePoolCallerSession) BoostContract() (common.Address, error) {
	return _CakePool.Contract.BoostContract(&_CakePool.CallOpts)
}

// CakePoolPID is a free data retrieval call binding the contract method 0x6d4710b9.
//
// Solidity: function cakePoolPID() view returns(uint256)
func (_CakePool *CakePoolCaller) CakePoolPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "cakePoolPID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePoolPID is a free data retrieval call binding the contract method 0x6d4710b9.
//
// Solidity: function cakePoolPID() view returns(uint256)
func (_CakePool *CakePoolSession) CakePoolPID() (*big.Int, error) {
	return _CakePool.Contract.CakePoolPID(&_CakePool.CallOpts)
}

// CakePoolPID is a free data retrieval call binding the contract method 0x6d4710b9.
//
// Solidity: function cakePoolPID() view returns(uint256)
func (_CakePool *CakePoolCallerSession) CakePoolPID() (*big.Int, error) {
	return _CakePool.Contract.CakePoolPID(&_CakePool.CallOpts)
}

// CalculateOverdueFee is a free data retrieval call binding the contract method 0x95dc14e1.
//
// Solidity: function calculateOverdueFee(address _user) view returns(uint256)
func (_CakePool *CakePoolCaller) CalculateOverdueFee(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "calculateOverdueFee", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateOverdueFee is a free data retrieval call binding the contract method 0x95dc14e1.
//
// Solidity: function calculateOverdueFee(address _user) view returns(uint256)
func (_CakePool *CakePoolSession) CalculateOverdueFee(_user common.Address) (*big.Int, error) {
	return _CakePool.Contract.CalculateOverdueFee(&_CakePool.CallOpts, _user)
}

// CalculateOverdueFee is a free data retrieval call binding the contract method 0x95dc14e1.
//
// Solidity: function calculateOverdueFee(address _user) view returns(uint256)
func (_CakePool *CakePoolCallerSession) CalculateOverdueFee(_user common.Address) (*big.Int, error) {
	return _CakePool.Contract.CalculateOverdueFee(&_CakePool.CallOpts, _user)
}

// CalculatePerformanceFee is a free data retrieval call binding the contract method 0xc6ed51be.
//
// Solidity: function calculatePerformanceFee(address _user) view returns(uint256)
func (_CakePool *CakePoolCaller) CalculatePerformanceFee(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "calculatePerformanceFee", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePerformanceFee is a free data retrieval call binding the contract method 0xc6ed51be.
//
// Solidity: function calculatePerformanceFee(address _user) view returns(uint256)
func (_CakePool *CakePoolSession) CalculatePerformanceFee(_user common.Address) (*big.Int, error) {
	return _CakePool.Contract.CalculatePerformanceFee(&_CakePool.CallOpts, _user)
}

// CalculatePerformanceFee is a free data retrieval call binding the contract method 0xc6ed51be.
//
// Solidity: function calculatePerformanceFee(address _user) view returns(uint256)
func (_CakePool *CakePoolCallerSession) CalculatePerformanceFee(_user common.Address) (*big.Int, error) {
	return _CakePool.Contract.CalculatePerformanceFee(&_CakePool.CallOpts, _user)
}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_CakePool *CakePoolCaller) CalculateTotalPendingCakeRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "calculateTotalPendingCakeRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_CakePool *CakePoolSession) CalculateTotalPendingCakeRewards() (*big.Int, error) {
	return _CakePool.Contract.CalculateTotalPendingCakeRewards(&_CakePool.CallOpts)
}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_CakePool *CakePoolCallerSession) CalculateTotalPendingCakeRewards() (*big.Int, error) {
	return _CakePool.Contract.CalculateTotalPendingCakeRewards(&_CakePool.CallOpts)
}

// CalculateWithdrawFee is a free data retrieval call binding the contract method 0x29a5cfd6.
//
// Solidity: function calculateWithdrawFee(address _user, uint256 _shares) view returns(uint256)
func (_CakePool *CakePoolCaller) CalculateWithdrawFee(opts *bind.CallOpts, _user common.Address, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "calculateWithdrawFee", _user, _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWithdrawFee is a free data retrieval call binding the contract method 0x29a5cfd6.
//
// Solidity: function calculateWithdrawFee(address _user, uint256 _shares) view returns(uint256)
func (_CakePool *CakePoolSession) CalculateWithdrawFee(_user common.Address, _shares *big.Int) (*big.Int, error) {
	return _CakePool.Contract.CalculateWithdrawFee(&_CakePool.CallOpts, _user, _shares)
}

// CalculateWithdrawFee is a free data retrieval call binding the contract method 0x29a5cfd6.
//
// Solidity: function calculateWithdrawFee(address _user, uint256 _shares) view returns(uint256)
func (_CakePool *CakePoolCallerSession) CalculateWithdrawFee(_user common.Address, _shares *big.Int) (*big.Int, error) {
	return _CakePool.Contract.CalculateWithdrawFee(&_CakePool.CallOpts, _user, _shares)
}

// FreeOverdueFeeUsers is a free data retrieval call binding the contract method 0x668679ba.
//
// Solidity: function freeOverdueFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolCaller) FreeOverdueFeeUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "freeOverdueFeeUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FreeOverdueFeeUsers is a free data retrieval call binding the contract method 0x668679ba.
//
// Solidity: function freeOverdueFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolSession) FreeOverdueFeeUsers(arg0 common.Address) (bool, error) {
	return _CakePool.Contract.FreeOverdueFeeUsers(&_CakePool.CallOpts, arg0)
}

// FreeOverdueFeeUsers is a free data retrieval call binding the contract method 0x668679ba.
//
// Solidity: function freeOverdueFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolCallerSession) FreeOverdueFeeUsers(arg0 common.Address) (bool, error) {
	return _CakePool.Contract.FreeOverdueFeeUsers(&_CakePool.CallOpts, arg0)
}

// FreePerformanceFeeUsers is a free data retrieval call binding the contract method 0x3fec4e32.
//
// Solidity: function freePerformanceFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolCaller) FreePerformanceFeeUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "freePerformanceFeeUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FreePerformanceFeeUsers is a free data retrieval call binding the contract method 0x3fec4e32.
//
// Solidity: function freePerformanceFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolSession) FreePerformanceFeeUsers(arg0 common.Address) (bool, error) {
	return _CakePool.Contract.FreePerformanceFeeUsers(&_CakePool.CallOpts, arg0)
}

// FreePerformanceFeeUsers is a free data retrieval call binding the contract method 0x3fec4e32.
//
// Solidity: function freePerformanceFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolCallerSession) FreePerformanceFeeUsers(arg0 common.Address) (bool, error) {
	return _CakePool.Contract.FreePerformanceFeeUsers(&_CakePool.CallOpts, arg0)
}

// FreeWithdrawFeeUsers is a free data retrieval call binding the contract method 0x87d4bda9.
//
// Solidity: function freeWithdrawFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolCaller) FreeWithdrawFeeUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "freeWithdrawFeeUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FreeWithdrawFeeUsers is a free data retrieval call binding the contract method 0x87d4bda9.
//
// Solidity: function freeWithdrawFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolSession) FreeWithdrawFeeUsers(arg0 common.Address) (bool, error) {
	return _CakePool.Contract.FreeWithdrawFeeUsers(&_CakePool.CallOpts, arg0)
}

// FreeWithdrawFeeUsers is a free data retrieval call binding the contract method 0x87d4bda9.
//
// Solidity: function freeWithdrawFeeUsers(address ) view returns(bool)
func (_CakePool *CakePoolCallerSession) FreeWithdrawFeeUsers(arg0 common.Address) (bool, error) {
	return _CakePool.Contract.FreeWithdrawFeeUsers(&_CakePool.CallOpts, arg0)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_CakePool *CakePoolCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "getPricePerFullShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_CakePool *CakePoolSession) GetPricePerFullShare() (*big.Int, error) {
	return _CakePool.Contract.GetPricePerFullShare(&_CakePool.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_CakePool *CakePoolCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _CakePool.Contract.GetPricePerFullShare(&_CakePool.CallOpts)
}

// MasterchefV2 is a free data retrieval call binding the contract method 0xcb528b52.
//
// Solidity: function masterchefV2() view returns(address)
func (_CakePool *CakePoolCaller) MasterchefV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "masterchefV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterchefV2 is a free data retrieval call binding the contract method 0xcb528b52.
//
// Solidity: function masterchefV2() view returns(address)
func (_CakePool *CakePoolSession) MasterchefV2() (common.Address, error) {
	return _CakePool.Contract.MasterchefV2(&_CakePool.CallOpts)
}

// MasterchefV2 is a free data retrieval call binding the contract method 0xcb528b52.
//
// Solidity: function masterchefV2() view returns(address)
func (_CakePool *CakePoolCallerSession) MasterchefV2() (common.Address, error) {
	return _CakePool.Contract.MasterchefV2(&_CakePool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CakePool *CakePoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CakePool *CakePoolSession) Operator() (common.Address, error) {
	return _CakePool.Contract.Operator(&_CakePool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CakePool *CakePoolCallerSession) Operator() (common.Address, error) {
	return _CakePool.Contract.Operator(&_CakePool.CallOpts)
}

// OverdueFee is a free data retrieval call binding the contract method 0xa5834e06.
//
// Solidity: function overdueFee() view returns(uint256)
func (_CakePool *CakePoolCaller) OverdueFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "overdueFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OverdueFee is a free data retrieval call binding the contract method 0xa5834e06.
//
// Solidity: function overdueFee() view returns(uint256)
func (_CakePool *CakePoolSession) OverdueFee() (*big.Int, error) {
	return _CakePool.Contract.OverdueFee(&_CakePool.CallOpts)
}

// OverdueFee is a free data retrieval call binding the contract method 0xa5834e06.
//
// Solidity: function overdueFee() view returns(uint256)
func (_CakePool *CakePoolCallerSession) OverdueFee() (*big.Int, error) {
	return _CakePool.Contract.OverdueFee(&_CakePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CakePool *CakePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CakePool *CakePoolSession) Owner() (common.Address, error) {
	return _CakePool.Contract.Owner(&_CakePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CakePool *CakePoolCallerSession) Owner() (common.Address, error) {
	return _CakePool.Contract.Owner(&_CakePool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CakePool *CakePoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CakePool *CakePoolSession) Paused() (bool, error) {
	return _CakePool.Contract.Paused(&_CakePool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CakePool *CakePoolCallerSession) Paused() (bool, error) {
	return _CakePool.Contract.Paused(&_CakePool.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_CakePool *CakePoolCaller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_CakePool *CakePoolSession) PerformanceFee() (*big.Int, error) {
	return _CakePool.Contract.PerformanceFee(&_CakePool.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_CakePool *CakePoolCallerSession) PerformanceFee() (*big.Int, error) {
	return _CakePool.Contract.PerformanceFee(&_CakePool.CallOpts)
}

// PerformanceFeeContract is a free data retrieval call binding the contract method 0x3eb78874.
//
// Solidity: function performanceFeeContract() view returns(uint256)
func (_CakePool *CakePoolCaller) PerformanceFeeContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "performanceFeeContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFeeContract is a free data retrieval call binding the contract method 0x3eb78874.
//
// Solidity: function performanceFeeContract() view returns(uint256)
func (_CakePool *CakePoolSession) PerformanceFeeContract() (*big.Int, error) {
	return _CakePool.Contract.PerformanceFeeContract(&_CakePool.CallOpts)
}

// PerformanceFeeContract is a free data retrieval call binding the contract method 0x3eb78874.
//
// Solidity: function performanceFeeContract() view returns(uint256)
func (_CakePool *CakePoolCallerSession) PerformanceFeeContract() (*big.Int, error) {
	return _CakePool.Contract.PerformanceFeeContract(&_CakePool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CakePool *CakePoolCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CakePool *CakePoolSession) Token() (common.Address, error) {
	return _CakePool.Contract.Token(&_CakePool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CakePool *CakePoolCallerSession) Token() (common.Address, error) {
	return _CakePool.Contract.Token(&_CakePool.CallOpts)
}

// TotalBoostDebt is a free data retrieval call binding the contract method 0xe73008bc.
//
// Solidity: function totalBoostDebt() view returns(uint256)
func (_CakePool *CakePoolCaller) TotalBoostDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "totalBoostDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBoostDebt is a free data retrieval call binding the contract method 0xe73008bc.
//
// Solidity: function totalBoostDebt() view returns(uint256)
func (_CakePool *CakePoolSession) TotalBoostDebt() (*big.Int, error) {
	return _CakePool.Contract.TotalBoostDebt(&_CakePool.CallOpts)
}

// TotalBoostDebt is a free data retrieval call binding the contract method 0xe73008bc.
//
// Solidity: function totalBoostDebt() view returns(uint256)
func (_CakePool *CakePoolCallerSession) TotalBoostDebt() (*big.Int, error) {
	return _CakePool.Contract.TotalBoostDebt(&_CakePool.CallOpts)
}

// TotalLockedAmount is a free data retrieval call binding the contract method 0x05a9f274.
//
// Solidity: function totalLockedAmount() view returns(uint256)
func (_CakePool *CakePoolCaller) TotalLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "totalLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedAmount is a free data retrieval call binding the contract method 0x05a9f274.
//
// Solidity: function totalLockedAmount() view returns(uint256)
func (_CakePool *CakePoolSession) TotalLockedAmount() (*big.Int, error) {
	return _CakePool.Contract.TotalLockedAmount(&_CakePool.CallOpts)
}

// TotalLockedAmount is a free data retrieval call binding the contract method 0x05a9f274.
//
// Solidity: function totalLockedAmount() view returns(uint256)
func (_CakePool *CakePoolCallerSession) TotalLockedAmount() (*big.Int, error) {
	return _CakePool.Contract.TotalLockedAmount(&_CakePool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_CakePool *CakePoolCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_CakePool *CakePoolSession) TotalShares() (*big.Int, error) {
	return _CakePool.Contract.TotalShares(&_CakePool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_CakePool *CakePoolCallerSession) TotalShares() (*big.Int, error) {
	return _CakePool.Contract.TotalShares(&_CakePool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_CakePool *CakePoolCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_CakePool *CakePoolSession) Treasury() (common.Address, error) {
	return _CakePool.Contract.Treasury(&_CakePool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_CakePool *CakePoolCallerSession) Treasury() (common.Address, error) {
	return _CakePool.Contract.Treasury(&_CakePool.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime, uint256 lockStartTime, uint256 lockEndTime, uint256 userBoostedShare, bool locked, uint256 lockedAmount)
func (_CakePool *CakePoolCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
	LockStartTime        *big.Int
	LockEndTime          *big.Int
	UserBoostedShare     *big.Int
	Locked               bool
	LockedAmount         *big.Int
}, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Shares               *big.Int
		LastDepositedTime    *big.Int
		CakeAtLastUserAction *big.Int
		LastUserActionTime   *big.Int
		LockStartTime        *big.Int
		LockEndTime          *big.Int
		UserBoostedShare     *big.Int
		Locked               bool
		LockedAmount         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastDepositedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CakeAtLastUserAction = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastUserActionTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockStartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LockEndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UserBoostedShare = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Locked = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.LockedAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime, uint256 lockStartTime, uint256 lockEndTime, uint256 userBoostedShare, bool locked, uint256 lockedAmount)
func (_CakePool *CakePoolSession) UserInfo(arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
	LockStartTime        *big.Int
	LockEndTime          *big.Int
	UserBoostedShare     *big.Int
	Locked               bool
	LockedAmount         *big.Int
}, error) {
	return _CakePool.Contract.UserInfo(&_CakePool.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime, uint256 lockStartTime, uint256 lockEndTime, uint256 userBoostedShare, bool locked, uint256 lockedAmount)
func (_CakePool *CakePoolCallerSession) UserInfo(arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
	LockStartTime        *big.Int
	LockEndTime          *big.Int
	UserBoostedShare     *big.Int
	Locked               bool
	LockedAmount         *big.Int
}, error) {
	return _CakePool.Contract.UserInfo(&_CakePool.CallOpts, arg0)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_CakePool *CakePoolCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_CakePool *CakePoolSession) WithdrawFee() (*big.Int, error) {
	return _CakePool.Contract.WithdrawFee(&_CakePool.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_CakePool *CakePoolCallerSession) WithdrawFee() (*big.Int, error) {
	return _CakePool.Contract.WithdrawFee(&_CakePool.CallOpts)
}

// WithdrawFeeContract is a free data retrieval call binding the contract method 0xe4b37ef5.
//
// Solidity: function withdrawFeeContract() view returns(uint256)
func (_CakePool *CakePoolCaller) WithdrawFeeContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "withdrawFeeContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeeContract is a free data retrieval call binding the contract method 0xe4b37ef5.
//
// Solidity: function withdrawFeeContract() view returns(uint256)
func (_CakePool *CakePoolSession) WithdrawFeeContract() (*big.Int, error) {
	return _CakePool.Contract.WithdrawFeeContract(&_CakePool.CallOpts)
}

// WithdrawFeeContract is a free data retrieval call binding the contract method 0xe4b37ef5.
//
// Solidity: function withdrawFeeContract() view returns(uint256)
func (_CakePool *CakePoolCallerSession) WithdrawFeeContract() (*big.Int, error) {
	return _CakePool.Contract.WithdrawFeeContract(&_CakePool.CallOpts)
}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_CakePool *CakePoolCaller) WithdrawFeePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CakePool.contract.Call(opts, &out, "withdrawFeePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_CakePool *CakePoolSession) WithdrawFeePeriod() (*big.Int, error) {
	return _CakePool.Contract.WithdrawFeePeriod(&_CakePool.CallOpts)
}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_CakePool *CakePoolCallerSession) WithdrawFeePeriod() (*big.Int, error) {
	return _CakePool.Contract.WithdrawFeePeriod(&_CakePool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _amount, uint256 _lockDuration) returns()
func (_CakePool *CakePoolTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "deposit", _amount, _lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _amount, uint256 _lockDuration) returns()
func (_CakePool *CakePoolSession) Deposit(_amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.Deposit(&_CakePool.TransactOpts, _amount, _lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _amount, uint256 _lockDuration) returns()
func (_CakePool *CakePoolTransactorSession) Deposit(_amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.Deposit(&_CakePool.TransactOpts, _amount, _lockDuration)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_CakePool *CakePoolTransactor) InCaseTokensGetStuck(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "inCaseTokensGetStuck", _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_CakePool *CakePoolSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.InCaseTokensGetStuck(&_CakePool.TransactOpts, _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_CakePool *CakePoolTransactorSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.InCaseTokensGetStuck(&_CakePool.TransactOpts, _token)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_CakePool *CakePoolTransactor) Init(opts *bind.TransactOpts, dummyToken common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "init", dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_CakePool *CakePoolSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.Init(&_CakePool.TransactOpts, dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_CakePool *CakePoolTransactorSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.Init(&_CakePool.TransactOpts, dummyToken)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CakePool *CakePoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CakePool *CakePoolSession) Pause() (*types.Transaction, error) {
	return _CakePool.Contract.Pause(&_CakePool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CakePool *CakePoolTransactorSession) Pause() (*types.Transaction, error) {
	return _CakePool.Contract.Pause(&_CakePool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CakePool *CakePoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CakePool *CakePoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _CakePool.Contract.RenounceOwnership(&_CakePool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CakePool *CakePoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CakePool.Contract.RenounceOwnership(&_CakePool.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_CakePool *CakePoolTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_CakePool *CakePoolSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetAdmin(&_CakePool.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_CakePool *CakePoolTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetAdmin(&_CakePool.TransactOpts, _admin)
}

// SetBoostContract is a paid mutator transaction binding the contract method 0xdef7869d.
//
// Solidity: function setBoostContract(address _boostContract) returns()
func (_CakePool *CakePoolTransactor) SetBoostContract(opts *bind.TransactOpts, _boostContract common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setBoostContract", _boostContract)
}

// SetBoostContract is a paid mutator transaction binding the contract method 0xdef7869d.
//
// Solidity: function setBoostContract(address _boostContract) returns()
func (_CakePool *CakePoolSession) SetBoostContract(_boostContract common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetBoostContract(&_CakePool.TransactOpts, _boostContract)
}

// SetBoostContract is a paid mutator transaction binding the contract method 0xdef7869d.
//
// Solidity: function setBoostContract(address _boostContract) returns()
func (_CakePool *CakePoolTransactorSession) SetBoostContract(_boostContract common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetBoostContract(&_CakePool.TransactOpts, _boostContract)
}

// SetBoostWeight is a paid mutator transaction binding the contract method 0x93c99e6a.
//
// Solidity: function setBoostWeight(uint256 _boostWeight) returns()
func (_CakePool *CakePoolTransactor) SetBoostWeight(opts *bind.TransactOpts, _boostWeight *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setBoostWeight", _boostWeight)
}

// SetBoostWeight is a paid mutator transaction binding the contract method 0x93c99e6a.
//
// Solidity: function setBoostWeight(uint256 _boostWeight) returns()
func (_CakePool *CakePoolSession) SetBoostWeight(_boostWeight *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetBoostWeight(&_CakePool.TransactOpts, _boostWeight)
}

// SetBoostWeight is a paid mutator transaction binding the contract method 0x93c99e6a.
//
// Solidity: function setBoostWeight(uint256 _boostWeight) returns()
func (_CakePool *CakePoolTransactorSession) SetBoostWeight(_boostWeight *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetBoostWeight(&_CakePool.TransactOpts, _boostWeight)
}

// SetDurationFactor is a paid mutator transaction binding the contract method 0xa3639b39.
//
// Solidity: function setDurationFactor(uint256 _durationFactor) returns()
func (_CakePool *CakePoolTransactor) SetDurationFactor(opts *bind.TransactOpts, _durationFactor *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setDurationFactor", _durationFactor)
}

// SetDurationFactor is a paid mutator transaction binding the contract method 0xa3639b39.
//
// Solidity: function setDurationFactor(uint256 _durationFactor) returns()
func (_CakePool *CakePoolSession) SetDurationFactor(_durationFactor *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetDurationFactor(&_CakePool.TransactOpts, _durationFactor)
}

// SetDurationFactor is a paid mutator transaction binding the contract method 0xa3639b39.
//
// Solidity: function setDurationFactor(uint256 _durationFactor) returns()
func (_CakePool *CakePoolTransactorSession) SetDurationFactor(_durationFactor *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetDurationFactor(&_CakePool.TransactOpts, _durationFactor)
}

// SetDurationFactorOverdue is a paid mutator transaction binding the contract method 0x35981921.
//
// Solidity: function setDurationFactorOverdue(uint256 _durationFactorOverdue) returns()
func (_CakePool *CakePoolTransactor) SetDurationFactorOverdue(opts *bind.TransactOpts, _durationFactorOverdue *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setDurationFactorOverdue", _durationFactorOverdue)
}

// SetDurationFactorOverdue is a paid mutator transaction binding the contract method 0x35981921.
//
// Solidity: function setDurationFactorOverdue(uint256 _durationFactorOverdue) returns()
func (_CakePool *CakePoolSession) SetDurationFactorOverdue(_durationFactorOverdue *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetDurationFactorOverdue(&_CakePool.TransactOpts, _durationFactorOverdue)
}

// SetDurationFactorOverdue is a paid mutator transaction binding the contract method 0x35981921.
//
// Solidity: function setDurationFactorOverdue(uint256 _durationFactorOverdue) returns()
func (_CakePool *CakePoolTransactorSession) SetDurationFactorOverdue(_durationFactorOverdue *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetDurationFactorOverdue(&_CakePool.TransactOpts, _durationFactorOverdue)
}

// SetFreePerformanceFeeUser is a paid mutator transaction binding the contract method 0x423b93ed.
//
// Solidity: function setFreePerformanceFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolTransactor) SetFreePerformanceFeeUser(opts *bind.TransactOpts, _user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setFreePerformanceFeeUser", _user, _free)
}

// SetFreePerformanceFeeUser is a paid mutator transaction binding the contract method 0x423b93ed.
//
// Solidity: function setFreePerformanceFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolSession) SetFreePerformanceFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.Contract.SetFreePerformanceFeeUser(&_CakePool.TransactOpts, _user, _free)
}

// SetFreePerformanceFeeUser is a paid mutator transaction binding the contract method 0x423b93ed.
//
// Solidity: function setFreePerformanceFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolTransactorSession) SetFreePerformanceFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.Contract.SetFreePerformanceFeeUser(&_CakePool.TransactOpts, _user, _free)
}

// SetMaxLockDuration is a paid mutator transaction binding the contract method 0xf786b958.
//
// Solidity: function setMaxLockDuration(uint256 _maxLockDuration) returns()
func (_CakePool *CakePoolTransactor) SetMaxLockDuration(opts *bind.TransactOpts, _maxLockDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setMaxLockDuration", _maxLockDuration)
}

// SetMaxLockDuration is a paid mutator transaction binding the contract method 0xf786b958.
//
// Solidity: function setMaxLockDuration(uint256 _maxLockDuration) returns()
func (_CakePool *CakePoolSession) SetMaxLockDuration(_maxLockDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetMaxLockDuration(&_CakePool.TransactOpts, _maxLockDuration)
}

// SetMaxLockDuration is a paid mutator transaction binding the contract method 0xf786b958.
//
// Solidity: function setMaxLockDuration(uint256 _maxLockDuration) returns()
func (_CakePool *CakePoolTransactorSession) SetMaxLockDuration(_maxLockDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetMaxLockDuration(&_CakePool.TransactOpts, _maxLockDuration)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_CakePool *CakePoolTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_CakePool *CakePoolSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetOperator(&_CakePool.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_CakePool *CakePoolTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetOperator(&_CakePool.TransactOpts, _operator)
}

// SetOverdueFee is a paid mutator transaction binding the contract method 0x0c59696b.
//
// Solidity: function setOverdueFee(uint256 _overdueFee) returns()
func (_CakePool *CakePoolTransactor) SetOverdueFee(opts *bind.TransactOpts, _overdueFee *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setOverdueFee", _overdueFee)
}

// SetOverdueFee is a paid mutator transaction binding the contract method 0x0c59696b.
//
// Solidity: function setOverdueFee(uint256 _overdueFee) returns()
func (_CakePool *CakePoolSession) SetOverdueFee(_overdueFee *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetOverdueFee(&_CakePool.TransactOpts, _overdueFee)
}

// SetOverdueFee is a paid mutator transaction binding the contract method 0x0c59696b.
//
// Solidity: function setOverdueFee(uint256 _overdueFee) returns()
func (_CakePool *CakePoolTransactorSession) SetOverdueFee(_overdueFee *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetOverdueFee(&_CakePool.TransactOpts, _overdueFee)
}

// SetOverdueFeeUser is a paid mutator transaction binding the contract method 0x4e4de1e9.
//
// Solidity: function setOverdueFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolTransactor) SetOverdueFeeUser(opts *bind.TransactOpts, _user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setOverdueFeeUser", _user, _free)
}

// SetOverdueFeeUser is a paid mutator transaction binding the contract method 0x4e4de1e9.
//
// Solidity: function setOverdueFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolSession) SetOverdueFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.Contract.SetOverdueFeeUser(&_CakePool.TransactOpts, _user, _free)
}

// SetOverdueFeeUser is a paid mutator transaction binding the contract method 0x4e4de1e9.
//
// Solidity: function setOverdueFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolTransactorSession) SetOverdueFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.Contract.SetOverdueFeeUser(&_CakePool.TransactOpts, _user, _free)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_CakePool *CakePoolTransactor) SetPerformanceFee(opts *bind.TransactOpts, _performanceFee *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setPerformanceFee", _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_CakePool *CakePoolSession) SetPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetPerformanceFee(&_CakePool.TransactOpts, _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_CakePool *CakePoolTransactorSession) SetPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetPerformanceFee(&_CakePool.TransactOpts, _performanceFee)
}

// SetPerformanceFeeContract is a paid mutator transaction binding the contract method 0xbb9f408d.
//
// Solidity: function setPerformanceFeeContract(uint256 _performanceFeeContract) returns()
func (_CakePool *CakePoolTransactor) SetPerformanceFeeContract(opts *bind.TransactOpts, _performanceFeeContract *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setPerformanceFeeContract", _performanceFeeContract)
}

// SetPerformanceFeeContract is a paid mutator transaction binding the contract method 0xbb9f408d.
//
// Solidity: function setPerformanceFeeContract(uint256 _performanceFeeContract) returns()
func (_CakePool *CakePoolSession) SetPerformanceFeeContract(_performanceFeeContract *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetPerformanceFeeContract(&_CakePool.TransactOpts, _performanceFeeContract)
}

// SetPerformanceFeeContract is a paid mutator transaction binding the contract method 0xbb9f408d.
//
// Solidity: function setPerformanceFeeContract(uint256 _performanceFeeContract) returns()
func (_CakePool *CakePoolTransactorSession) SetPerformanceFeeContract(_performanceFeeContract *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetPerformanceFeeContract(&_CakePool.TransactOpts, _performanceFeeContract)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_CakePool *CakePoolTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_CakePool *CakePoolSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetTreasury(&_CakePool.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_CakePool *CakePoolTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetTreasury(&_CakePool.TransactOpts, _treasury)
}

// SetUnlockFreeDuration is a paid mutator transaction binding the contract method 0xc54d349c.
//
// Solidity: function setUnlockFreeDuration(uint256 _unlockFreeDuration) returns()
func (_CakePool *CakePoolTransactor) SetUnlockFreeDuration(opts *bind.TransactOpts, _unlockFreeDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setUnlockFreeDuration", _unlockFreeDuration)
}

// SetUnlockFreeDuration is a paid mutator transaction binding the contract method 0xc54d349c.
//
// Solidity: function setUnlockFreeDuration(uint256 _unlockFreeDuration) returns()
func (_CakePool *CakePoolSession) SetUnlockFreeDuration(_unlockFreeDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetUnlockFreeDuration(&_CakePool.TransactOpts, _unlockFreeDuration)
}

// SetUnlockFreeDuration is a paid mutator transaction binding the contract method 0xc54d349c.
//
// Solidity: function setUnlockFreeDuration(uint256 _unlockFreeDuration) returns()
func (_CakePool *CakePoolTransactorSession) SetUnlockFreeDuration(_unlockFreeDuration *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetUnlockFreeDuration(&_CakePool.TransactOpts, _unlockFreeDuration)
}

// SetVCakeContract is a paid mutator transaction binding the contract method 0xd826ed06.
//
// Solidity: function setVCakeContract(address _VCake) returns()
func (_CakePool *CakePoolTransactor) SetVCakeContract(opts *bind.TransactOpts, _VCake common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setVCakeContract", _VCake)
}

// SetVCakeContract is a paid mutator transaction binding the contract method 0xd826ed06.
//
// Solidity: function setVCakeContract(address _VCake) returns()
func (_CakePool *CakePoolSession) SetVCakeContract(_VCake common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetVCakeContract(&_CakePool.TransactOpts, _VCake)
}

// SetVCakeContract is a paid mutator transaction binding the contract method 0xd826ed06.
//
// Solidity: function setVCakeContract(address _VCake) returns()
func (_CakePool *CakePoolTransactorSession) SetVCakeContract(_VCake common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.SetVCakeContract(&_CakePool.TransactOpts, _VCake)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_CakePool *CakePoolTransactor) SetWithdrawFee(opts *bind.TransactOpts, _withdrawFee *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setWithdrawFee", _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_CakePool *CakePoolSession) SetWithdrawFee(_withdrawFee *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFee(&_CakePool.TransactOpts, _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_CakePool *CakePoolTransactorSession) SetWithdrawFee(_withdrawFee *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFee(&_CakePool.TransactOpts, _withdrawFee)
}

// SetWithdrawFeeContract is a paid mutator transaction binding the contract method 0x14ff3039.
//
// Solidity: function setWithdrawFeeContract(uint256 _withdrawFeeContract) returns()
func (_CakePool *CakePoolTransactor) SetWithdrawFeeContract(opts *bind.TransactOpts, _withdrawFeeContract *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setWithdrawFeeContract", _withdrawFeeContract)
}

// SetWithdrawFeeContract is a paid mutator transaction binding the contract method 0x14ff3039.
//
// Solidity: function setWithdrawFeeContract(uint256 _withdrawFeeContract) returns()
func (_CakePool *CakePoolSession) SetWithdrawFeeContract(_withdrawFeeContract *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFeeContract(&_CakePool.TransactOpts, _withdrawFeeContract)
}

// SetWithdrawFeeContract is a paid mutator transaction binding the contract method 0x14ff3039.
//
// Solidity: function setWithdrawFeeContract(uint256 _withdrawFeeContract) returns()
func (_CakePool *CakePoolTransactorSession) SetWithdrawFeeContract(_withdrawFeeContract *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFeeContract(&_CakePool.TransactOpts, _withdrawFeeContract)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_CakePool *CakePoolTransactor) SetWithdrawFeePeriod(opts *bind.TransactOpts, _withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setWithdrawFeePeriod", _withdrawFeePeriod)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_CakePool *CakePoolSession) SetWithdrawFeePeriod(_withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFeePeriod(&_CakePool.TransactOpts, _withdrawFeePeriod)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_CakePool *CakePoolTransactorSession) SetWithdrawFeePeriod(_withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFeePeriod(&_CakePool.TransactOpts, _withdrawFeePeriod)
}

// SetWithdrawFeeUser is a paid mutator transaction binding the contract method 0xbeba0fa0.
//
// Solidity: function setWithdrawFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolTransactor) SetWithdrawFeeUser(opts *bind.TransactOpts, _user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "setWithdrawFeeUser", _user, _free)
}

// SetWithdrawFeeUser is a paid mutator transaction binding the contract method 0xbeba0fa0.
//
// Solidity: function setWithdrawFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolSession) SetWithdrawFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFeeUser(&_CakePool.TransactOpts, _user, _free)
}

// SetWithdrawFeeUser is a paid mutator transaction binding the contract method 0xbeba0fa0.
//
// Solidity: function setWithdrawFeeUser(address _user, bool _free) returns()
func (_CakePool *CakePoolTransactorSession) SetWithdrawFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _CakePool.Contract.SetWithdrawFeeUser(&_CakePool.TransactOpts, _user, _free)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CakePool *CakePoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CakePool *CakePoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.TransferOwnership(&_CakePool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CakePool *CakePoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.TransferOwnership(&_CakePool.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _user) returns()
func (_CakePool *CakePoolTransactor) Unlock(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "unlock", _user)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _user) returns()
func (_CakePool *CakePoolSession) Unlock(_user common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.Unlock(&_CakePool.TransactOpts, _user)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _user) returns()
func (_CakePool *CakePoolTransactorSession) Unlock(_user common.Address) (*types.Transaction, error) {
	return _CakePool.Contract.Unlock(&_CakePool.TransactOpts, _user)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CakePool *CakePoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CakePool *CakePoolSession) Unpause() (*types.Transaction, error) {
	return _CakePool.Contract.Unpause(&_CakePool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CakePool *CakePoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _CakePool.Contract.Unpause(&_CakePool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_CakePool *CakePoolTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_CakePool *CakePoolSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.Withdraw(&_CakePool.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_CakePool *CakePoolTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.Withdraw(&_CakePool.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_CakePool *CakePoolTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_CakePool *CakePoolSession) WithdrawAll() (*types.Transaction, error) {
	return _CakePool.Contract.WithdrawAll(&_CakePool.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_CakePool *CakePoolTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _CakePool.Contract.WithdrawAll(&_CakePool.TransactOpts)
}

// WithdrawByAmount is a paid mutator transaction binding the contract method 0x5521e9bf.
//
// Solidity: function withdrawByAmount(uint256 _amount) returns()
func (_CakePool *CakePoolTransactor) WithdrawByAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CakePool.contract.Transact(opts, "withdrawByAmount", _amount)
}

// WithdrawByAmount is a paid mutator transaction binding the contract method 0x5521e9bf.
//
// Solidity: function withdrawByAmount(uint256 _amount) returns()
func (_CakePool *CakePoolSession) WithdrawByAmount(_amount *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.WithdrawByAmount(&_CakePool.TransactOpts, _amount)
}

// WithdrawByAmount is a paid mutator transaction binding the contract method 0x5521e9bf.
//
// Solidity: function withdrawByAmount(uint256 _amount) returns()
func (_CakePool *CakePoolTransactorSession) WithdrawByAmount(_amount *big.Int) (*types.Transaction, error) {
	return _CakePool.Contract.WithdrawByAmount(&_CakePool.TransactOpts, _amount)
}

// CakePoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CakePool contract.
type CakePoolDepositIterator struct {
	Event *CakePoolDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolDeposit represents a Deposit event raised by the CakePool contract.
type CakePoolDeposit struct {
	Sender            common.Address
	Amount            *big.Int
	Shares            *big.Int
	Duration          *big.Int
	LastDepositedTime *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 duration, uint256 lastDepositedTime)
func (_CakePool *CakePoolFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*CakePoolDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakePoolDepositIterator{contract: _CakePool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 duration, uint256 lastDepositedTime)
func (_CakePool *CakePoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CakePoolDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolDeposit)
				if err := _CakePool.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 duration, uint256 lastDepositedTime)
func (_CakePool *CakePoolFilterer) ParseDeposit(log types.Log) (*CakePoolDeposit, error) {
	event := new(CakePoolDeposit)
	if err := _CakePool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolFreeFeeUserIterator is returned from FilterFreeFeeUser and is used to iterate over the raw logs and unpacked data for FreeFeeUser events raised by the CakePool contract.
type CakePoolFreeFeeUserIterator struct {
	Event *CakePoolFreeFeeUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolFreeFeeUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolFreeFeeUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolFreeFeeUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolFreeFeeUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolFreeFeeUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolFreeFeeUser represents a FreeFeeUser event raised by the CakePool contract.
type CakePoolFreeFeeUser struct {
	User common.Address
	Free bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFreeFeeUser is a free log retrieval operation binding the contract event 0x3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d44.
//
// Solidity: event FreeFeeUser(address indexed user, bool indexed free)
func (_CakePool *CakePoolFilterer) FilterFreeFeeUser(opts *bind.FilterOpts, user []common.Address, free []bool) (*CakePoolFreeFeeUserIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var freeRule []interface{}
	for _, freeItem := range free {
		freeRule = append(freeRule, freeItem)
	}

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "FreeFeeUser", userRule, freeRule)
	if err != nil {
		return nil, err
	}
	return &CakePoolFreeFeeUserIterator{contract: _CakePool.contract, event: "FreeFeeUser", logs: logs, sub: sub}, nil
}

// WatchFreeFeeUser is a free log subscription operation binding the contract event 0x3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d44.
//
// Solidity: event FreeFeeUser(address indexed user, bool indexed free)
func (_CakePool *CakePoolFilterer) WatchFreeFeeUser(opts *bind.WatchOpts, sink chan<- *CakePoolFreeFeeUser, user []common.Address, free []bool) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var freeRule []interface{}
	for _, freeItem := range free {
		freeRule = append(freeRule, freeItem)
	}

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "FreeFeeUser", userRule, freeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolFreeFeeUser)
				if err := _CakePool.contract.UnpackLog(event, "FreeFeeUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreeFeeUser is a log parse operation binding the contract event 0x3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d44.
//
// Solidity: event FreeFeeUser(address indexed user, bool indexed free)
func (_CakePool *CakePoolFilterer) ParseFreeFeeUser(log types.Log) (*CakePoolFreeFeeUser, error) {
	event := new(CakePoolFreeFeeUser)
	if err := _CakePool.contract.UnpackLog(event, "FreeFeeUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the CakePool contract.
type CakePoolHarvestIterator struct {
	Event *CakePoolHarvest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolHarvest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolHarvest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolHarvest represents a Harvest event raised by the CakePool contract.
type CakePoolHarvest struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed sender, uint256 amount)
func (_CakePool *CakePoolFilterer) FilterHarvest(opts *bind.FilterOpts, sender []common.Address) (*CakePoolHarvestIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Harvest", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakePoolHarvestIterator{contract: _CakePool.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed sender, uint256 amount)
func (_CakePool *CakePoolFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *CakePoolHarvest, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Harvest", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolHarvest)
				if err := _CakePool.contract.UnpackLog(event, "Harvest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHarvest is a log parse operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed sender, uint256 amount)
func (_CakePool *CakePoolFilterer) ParseHarvest(log types.Log) (*CakePoolHarvest, error) {
	event := new(CakePoolHarvest)
	if err := _CakePool.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the CakePool contract.
type CakePoolInitIterator struct {
	Event *CakePoolInit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolInit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolInit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolInit represents a Init event raised by the CakePool contract.
type CakePoolInit struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_CakePool *CakePoolFilterer) FilterInit(opts *bind.FilterOpts) (*CakePoolInitIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return &CakePoolInitIterator{contract: _CakePool.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_CakePool *CakePoolFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *CakePoolInit) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolInit)
				if err := _CakePool.contract.UnpackLog(event, "Init", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInit is a log parse operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_CakePool *CakePoolFilterer) ParseInit(log types.Log) (*CakePoolInit, error) {
	event := new(CakePoolInit)
	if err := _CakePool.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the CakePool contract.
type CakePoolLockIterator struct {
	Event *CakePoolLock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolLock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolLock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolLock represents a Lock event raised by the CakePool contract.
type CakePoolLock struct {
	Sender         common.Address
	LockedAmount   *big.Int
	Shares         *big.Int
	LockedDuration *big.Int
	BlockTimestamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x2b943276e5d747f6f7dd46d3b880d8874cb8d6b9b88ca1903990a2738e7dc7a1.
//
// Solidity: event Lock(address indexed sender, uint256 lockedAmount, uint256 shares, uint256 lockedDuration, uint256 blockTimestamp)
func (_CakePool *CakePoolFilterer) FilterLock(opts *bind.FilterOpts, sender []common.Address) (*CakePoolLockIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Lock", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakePoolLockIterator{contract: _CakePool.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x2b943276e5d747f6f7dd46d3b880d8874cb8d6b9b88ca1903990a2738e7dc7a1.
//
// Solidity: event Lock(address indexed sender, uint256 lockedAmount, uint256 shares, uint256 lockedDuration, uint256 blockTimestamp)
func (_CakePool *CakePoolFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *CakePoolLock, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Lock", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolLock)
				if err := _CakePool.contract.UnpackLog(event, "Lock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLock is a log parse operation binding the contract event 0x2b943276e5d747f6f7dd46d3b880d8874cb8d6b9b88ca1903990a2738e7dc7a1.
//
// Solidity: event Lock(address indexed sender, uint256 lockedAmount, uint256 shares, uint256 lockedDuration, uint256 blockTimestamp)
func (_CakePool *CakePoolFilterer) ParseLock(log types.Log) (*CakePoolLock, error) {
	event := new(CakePoolLock)
	if err := _CakePool.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the CakePool contract.
type CakePoolNewAdminIterator struct {
	Event *CakePoolNewAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewAdmin represents a NewAdmin event raised by the CakePool contract.
type CakePoolNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address admin)
func (_CakePool *CakePoolFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*CakePoolNewAdminIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewAdminIterator{contract: _CakePool.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address admin)
func (_CakePool *CakePoolFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CakePoolNewAdmin) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewAdmin)
				if err := _CakePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address admin)
func (_CakePool *CakePoolFilterer) ParseNewAdmin(log types.Log) (*CakePoolNewAdmin, error) {
	event := new(CakePoolNewAdmin)
	if err := _CakePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewBoostContractIterator is returned from FilterNewBoostContract and is used to iterate over the raw logs and unpacked data for NewBoostContract events raised by the CakePool contract.
type CakePoolNewBoostContractIterator struct {
	Event *CakePoolNewBoostContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewBoostContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewBoostContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewBoostContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewBoostContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewBoostContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewBoostContract represents a NewBoostContract event raised by the CakePool contract.
type CakePoolNewBoostContract struct {
	BoostContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewBoostContract is a free log retrieval operation binding the contract event 0x8f49a182922022d9119a1a6aeeca151b4a5665e86bd61c1ff32e152d459558b2.
//
// Solidity: event NewBoostContract(address boostContract)
func (_CakePool *CakePoolFilterer) FilterNewBoostContract(opts *bind.FilterOpts) (*CakePoolNewBoostContractIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewBoostContract")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewBoostContractIterator{contract: _CakePool.contract, event: "NewBoostContract", logs: logs, sub: sub}, nil
}

// WatchNewBoostContract is a free log subscription operation binding the contract event 0x8f49a182922022d9119a1a6aeeca151b4a5665e86bd61c1ff32e152d459558b2.
//
// Solidity: event NewBoostContract(address boostContract)
func (_CakePool *CakePoolFilterer) WatchNewBoostContract(opts *bind.WatchOpts, sink chan<- *CakePoolNewBoostContract) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewBoostContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewBoostContract)
				if err := _CakePool.contract.UnpackLog(event, "NewBoostContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBoostContract is a log parse operation binding the contract event 0x8f49a182922022d9119a1a6aeeca151b4a5665e86bd61c1ff32e152d459558b2.
//
// Solidity: event NewBoostContract(address boostContract)
func (_CakePool *CakePoolFilterer) ParseNewBoostContract(log types.Log) (*CakePoolNewBoostContract, error) {
	event := new(CakePoolNewBoostContract)
	if err := _CakePool.contract.UnpackLog(event, "NewBoostContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewBoostWeightIterator is returned from FilterNewBoostWeight and is used to iterate over the raw logs and unpacked data for NewBoostWeight events raised by the CakePool contract.
type CakePoolNewBoostWeightIterator struct {
	Event *CakePoolNewBoostWeight // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewBoostWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewBoostWeight)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewBoostWeight)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewBoostWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewBoostWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewBoostWeight represents a NewBoostWeight event raised by the CakePool contract.
type CakePoolNewBoostWeight struct {
	BoostWeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBoostWeight is a free log retrieval operation binding the contract event 0x7666dfff8c3377938e522b4eed3aff079973a976f95969db60a406d49f40da4e.
//
// Solidity: event NewBoostWeight(uint256 boostWeight)
func (_CakePool *CakePoolFilterer) FilterNewBoostWeight(opts *bind.FilterOpts) (*CakePoolNewBoostWeightIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewBoostWeight")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewBoostWeightIterator{contract: _CakePool.contract, event: "NewBoostWeight", logs: logs, sub: sub}, nil
}

// WatchNewBoostWeight is a free log subscription operation binding the contract event 0x7666dfff8c3377938e522b4eed3aff079973a976f95969db60a406d49f40da4e.
//
// Solidity: event NewBoostWeight(uint256 boostWeight)
func (_CakePool *CakePoolFilterer) WatchNewBoostWeight(opts *bind.WatchOpts, sink chan<- *CakePoolNewBoostWeight) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewBoostWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewBoostWeight)
				if err := _CakePool.contract.UnpackLog(event, "NewBoostWeight", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBoostWeight is a log parse operation binding the contract event 0x7666dfff8c3377938e522b4eed3aff079973a976f95969db60a406d49f40da4e.
//
// Solidity: event NewBoostWeight(uint256 boostWeight)
func (_CakePool *CakePoolFilterer) ParseNewBoostWeight(log types.Log) (*CakePoolNewBoostWeight, error) {
	event := new(CakePoolNewBoostWeight)
	if err := _CakePool.contract.UnpackLog(event, "NewBoostWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewDurationFactorIterator is returned from FilterNewDurationFactor and is used to iterate over the raw logs and unpacked data for NewDurationFactor events raised by the CakePool contract.
type CakePoolNewDurationFactorIterator struct {
	Event *CakePoolNewDurationFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewDurationFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewDurationFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewDurationFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewDurationFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewDurationFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewDurationFactor represents a NewDurationFactor event raised by the CakePool contract.
type CakePoolNewDurationFactor struct {
	DurationFactor *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewDurationFactor is a free log retrieval operation binding the contract event 0x9478eb023aac0a7d58a4e935377056bf27cf5b72a2300725f831817a8f62fbde.
//
// Solidity: event NewDurationFactor(uint256 durationFactor)
func (_CakePool *CakePoolFilterer) FilterNewDurationFactor(opts *bind.FilterOpts) (*CakePoolNewDurationFactorIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewDurationFactor")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewDurationFactorIterator{contract: _CakePool.contract, event: "NewDurationFactor", logs: logs, sub: sub}, nil
}

// WatchNewDurationFactor is a free log subscription operation binding the contract event 0x9478eb023aac0a7d58a4e935377056bf27cf5b72a2300725f831817a8f62fbde.
//
// Solidity: event NewDurationFactor(uint256 durationFactor)
func (_CakePool *CakePoolFilterer) WatchNewDurationFactor(opts *bind.WatchOpts, sink chan<- *CakePoolNewDurationFactor) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewDurationFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewDurationFactor)
				if err := _CakePool.contract.UnpackLog(event, "NewDurationFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDurationFactor is a log parse operation binding the contract event 0x9478eb023aac0a7d58a4e935377056bf27cf5b72a2300725f831817a8f62fbde.
//
// Solidity: event NewDurationFactor(uint256 durationFactor)
func (_CakePool *CakePoolFilterer) ParseNewDurationFactor(log types.Log) (*CakePoolNewDurationFactor, error) {
	event := new(CakePoolNewDurationFactor)
	if err := _CakePool.contract.UnpackLog(event, "NewDurationFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewDurationFactorOverdueIterator is returned from FilterNewDurationFactorOverdue and is used to iterate over the raw logs and unpacked data for NewDurationFactorOverdue events raised by the CakePool contract.
type CakePoolNewDurationFactorOverdueIterator struct {
	Event *CakePoolNewDurationFactorOverdue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewDurationFactorOverdueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewDurationFactorOverdue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewDurationFactorOverdue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewDurationFactorOverdueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewDurationFactorOverdueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewDurationFactorOverdue represents a NewDurationFactorOverdue event raised by the CakePool contract.
type CakePoolNewDurationFactorOverdue struct {
	DurationFactorOverdue *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewDurationFactorOverdue is a free log retrieval operation binding the contract event 0x18b6d179114082d7eda9837e15a39eb30032d5f3df00487a67541398f48fabfe.
//
// Solidity: event NewDurationFactorOverdue(uint256 durationFactorOverdue)
func (_CakePool *CakePoolFilterer) FilterNewDurationFactorOverdue(opts *bind.FilterOpts) (*CakePoolNewDurationFactorOverdueIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewDurationFactorOverdue")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewDurationFactorOverdueIterator{contract: _CakePool.contract, event: "NewDurationFactorOverdue", logs: logs, sub: sub}, nil
}

// WatchNewDurationFactorOverdue is a free log subscription operation binding the contract event 0x18b6d179114082d7eda9837e15a39eb30032d5f3df00487a67541398f48fabfe.
//
// Solidity: event NewDurationFactorOverdue(uint256 durationFactorOverdue)
func (_CakePool *CakePoolFilterer) WatchNewDurationFactorOverdue(opts *bind.WatchOpts, sink chan<- *CakePoolNewDurationFactorOverdue) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewDurationFactorOverdue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewDurationFactorOverdue)
				if err := _CakePool.contract.UnpackLog(event, "NewDurationFactorOverdue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDurationFactorOverdue is a log parse operation binding the contract event 0x18b6d179114082d7eda9837e15a39eb30032d5f3df00487a67541398f48fabfe.
//
// Solidity: event NewDurationFactorOverdue(uint256 durationFactorOverdue)
func (_CakePool *CakePoolFilterer) ParseNewDurationFactorOverdue(log types.Log) (*CakePoolNewDurationFactorOverdue, error) {
	event := new(CakePoolNewDurationFactorOverdue)
	if err := _CakePool.contract.UnpackLog(event, "NewDurationFactorOverdue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewMaxLockDurationIterator is returned from FilterNewMaxLockDuration and is used to iterate over the raw logs and unpacked data for NewMaxLockDuration events raised by the CakePool contract.
type CakePoolNewMaxLockDurationIterator struct {
	Event *CakePoolNewMaxLockDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewMaxLockDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewMaxLockDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewMaxLockDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewMaxLockDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewMaxLockDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewMaxLockDuration represents a NewMaxLockDuration event raised by the CakePool contract.
type CakePoolNewMaxLockDuration struct {
	MaxLockDuration *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewMaxLockDuration is a free log retrieval operation binding the contract event 0xcab2f3455b51b6ca5377e84fccd0f890b6f6ca36c02e18b6d36cb34f469fe4fe.
//
// Solidity: event NewMaxLockDuration(uint256 maxLockDuration)
func (_CakePool *CakePoolFilterer) FilterNewMaxLockDuration(opts *bind.FilterOpts) (*CakePoolNewMaxLockDurationIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewMaxLockDuration")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewMaxLockDurationIterator{contract: _CakePool.contract, event: "NewMaxLockDuration", logs: logs, sub: sub}, nil
}

// WatchNewMaxLockDuration is a free log subscription operation binding the contract event 0xcab2f3455b51b6ca5377e84fccd0f890b6f6ca36c02e18b6d36cb34f469fe4fe.
//
// Solidity: event NewMaxLockDuration(uint256 maxLockDuration)
func (_CakePool *CakePoolFilterer) WatchNewMaxLockDuration(opts *bind.WatchOpts, sink chan<- *CakePoolNewMaxLockDuration) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewMaxLockDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewMaxLockDuration)
				if err := _CakePool.contract.UnpackLog(event, "NewMaxLockDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewMaxLockDuration is a log parse operation binding the contract event 0xcab2f3455b51b6ca5377e84fccd0f890b6f6ca36c02e18b6d36cb34f469fe4fe.
//
// Solidity: event NewMaxLockDuration(uint256 maxLockDuration)
func (_CakePool *CakePoolFilterer) ParseNewMaxLockDuration(log types.Log) (*CakePoolNewMaxLockDuration, error) {
	event := new(CakePoolNewMaxLockDuration)
	if err := _CakePool.contract.UnpackLog(event, "NewMaxLockDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the CakePool contract.
type CakePoolNewOperatorIterator struct {
	Event *CakePoolNewOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewOperator represents a NewOperator event raised by the CakePool contract.
type CakePoolNewOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address operator)
func (_CakePool *CakePoolFilterer) FilterNewOperator(opts *bind.FilterOpts) (*CakePoolNewOperatorIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewOperator")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewOperatorIterator{contract: _CakePool.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address operator)
func (_CakePool *CakePoolFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *CakePoolNewOperator) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewOperator)
				if err := _CakePool.contract.UnpackLog(event, "NewOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOperator is a log parse operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address operator)
func (_CakePool *CakePoolFilterer) ParseNewOperator(log types.Log) (*CakePoolNewOperator, error) {
	event := new(CakePoolNewOperator)
	if err := _CakePool.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewOverdueFeeIterator is returned from FilterNewOverdueFee and is used to iterate over the raw logs and unpacked data for NewOverdueFee events raised by the CakePool contract.
type CakePoolNewOverdueFeeIterator struct {
	Event *CakePoolNewOverdueFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewOverdueFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewOverdueFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewOverdueFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewOverdueFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewOverdueFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewOverdueFee represents a NewOverdueFee event raised by the CakePool contract.
type CakePoolNewOverdueFee struct {
	OverdueFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewOverdueFee is a free log retrieval operation binding the contract event 0xf4bd1c5978320077e792afbb3911e8cab1325ce28a6b3e67f9067a1d80692961.
//
// Solidity: event NewOverdueFee(uint256 overdueFee)
func (_CakePool *CakePoolFilterer) FilterNewOverdueFee(opts *bind.FilterOpts) (*CakePoolNewOverdueFeeIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewOverdueFee")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewOverdueFeeIterator{contract: _CakePool.contract, event: "NewOverdueFee", logs: logs, sub: sub}, nil
}

// WatchNewOverdueFee is a free log subscription operation binding the contract event 0xf4bd1c5978320077e792afbb3911e8cab1325ce28a6b3e67f9067a1d80692961.
//
// Solidity: event NewOverdueFee(uint256 overdueFee)
func (_CakePool *CakePoolFilterer) WatchNewOverdueFee(opts *bind.WatchOpts, sink chan<- *CakePoolNewOverdueFee) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewOverdueFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewOverdueFee)
				if err := _CakePool.contract.UnpackLog(event, "NewOverdueFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOverdueFee is a log parse operation binding the contract event 0xf4bd1c5978320077e792afbb3911e8cab1325ce28a6b3e67f9067a1d80692961.
//
// Solidity: event NewOverdueFee(uint256 overdueFee)
func (_CakePool *CakePoolFilterer) ParseNewOverdueFee(log types.Log) (*CakePoolNewOverdueFee, error) {
	event := new(CakePoolNewOverdueFee)
	if err := _CakePool.contract.UnpackLog(event, "NewOverdueFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewPerformanceFeeIterator is returned from FilterNewPerformanceFee and is used to iterate over the raw logs and unpacked data for NewPerformanceFee events raised by the CakePool contract.
type CakePoolNewPerformanceFeeIterator struct {
	Event *CakePoolNewPerformanceFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewPerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewPerformanceFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewPerformanceFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewPerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewPerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewPerformanceFee represents a NewPerformanceFee event raised by the CakePool contract.
type CakePoolNewPerformanceFee struct {
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPerformanceFee is a free log retrieval operation binding the contract event 0xefeafcf03e479a9566d7ef321b4816de0ba19cfa3cd0fae2f8c5f4a0afb342c4.
//
// Solidity: event NewPerformanceFee(uint256 performanceFee)
func (_CakePool *CakePoolFilterer) FilterNewPerformanceFee(opts *bind.FilterOpts) (*CakePoolNewPerformanceFeeIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewPerformanceFee")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewPerformanceFeeIterator{contract: _CakePool.contract, event: "NewPerformanceFee", logs: logs, sub: sub}, nil
}

// WatchNewPerformanceFee is a free log subscription operation binding the contract event 0xefeafcf03e479a9566d7ef321b4816de0ba19cfa3cd0fae2f8c5f4a0afb342c4.
//
// Solidity: event NewPerformanceFee(uint256 performanceFee)
func (_CakePool *CakePoolFilterer) WatchNewPerformanceFee(opts *bind.WatchOpts, sink chan<- *CakePoolNewPerformanceFee) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewPerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewPerformanceFee)
				if err := _CakePool.contract.UnpackLog(event, "NewPerformanceFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPerformanceFee is a log parse operation binding the contract event 0xefeafcf03e479a9566d7ef321b4816de0ba19cfa3cd0fae2f8c5f4a0afb342c4.
//
// Solidity: event NewPerformanceFee(uint256 performanceFee)
func (_CakePool *CakePoolFilterer) ParseNewPerformanceFee(log types.Log) (*CakePoolNewPerformanceFee, error) {
	event := new(CakePoolNewPerformanceFee)
	if err := _CakePool.contract.UnpackLog(event, "NewPerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewPerformanceFeeContractIterator is returned from FilterNewPerformanceFeeContract and is used to iterate over the raw logs and unpacked data for NewPerformanceFeeContract events raised by the CakePool contract.
type CakePoolNewPerformanceFeeContractIterator struct {
	Event *CakePoolNewPerformanceFeeContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewPerformanceFeeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewPerformanceFeeContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewPerformanceFeeContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewPerformanceFeeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewPerformanceFeeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewPerformanceFeeContract represents a NewPerformanceFeeContract event raised by the CakePool contract.
type CakePoolNewPerformanceFeeContract struct {
	PerformanceFeeContract *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewPerformanceFeeContract is a free log retrieval operation binding the contract event 0xc5d25457b67b87678c987375af13f6e50beb3ad7bfd009da26766ae986eaa20d.
//
// Solidity: event NewPerformanceFeeContract(uint256 performanceFeeContract)
func (_CakePool *CakePoolFilterer) FilterNewPerformanceFeeContract(opts *bind.FilterOpts) (*CakePoolNewPerformanceFeeContractIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewPerformanceFeeContract")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewPerformanceFeeContractIterator{contract: _CakePool.contract, event: "NewPerformanceFeeContract", logs: logs, sub: sub}, nil
}

// WatchNewPerformanceFeeContract is a free log subscription operation binding the contract event 0xc5d25457b67b87678c987375af13f6e50beb3ad7bfd009da26766ae986eaa20d.
//
// Solidity: event NewPerformanceFeeContract(uint256 performanceFeeContract)
func (_CakePool *CakePoolFilterer) WatchNewPerformanceFeeContract(opts *bind.WatchOpts, sink chan<- *CakePoolNewPerformanceFeeContract) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewPerformanceFeeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewPerformanceFeeContract)
				if err := _CakePool.contract.UnpackLog(event, "NewPerformanceFeeContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPerformanceFeeContract is a log parse operation binding the contract event 0xc5d25457b67b87678c987375af13f6e50beb3ad7bfd009da26766ae986eaa20d.
//
// Solidity: event NewPerformanceFeeContract(uint256 performanceFeeContract)
func (_CakePool *CakePoolFilterer) ParseNewPerformanceFeeContract(log types.Log) (*CakePoolNewPerformanceFeeContract, error) {
	event := new(CakePoolNewPerformanceFeeContract)
	if err := _CakePool.contract.UnpackLog(event, "NewPerformanceFeeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewTreasuryIterator is returned from FilterNewTreasury and is used to iterate over the raw logs and unpacked data for NewTreasury events raised by the CakePool contract.
type CakePoolNewTreasuryIterator struct {
	Event *CakePoolNewTreasury // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewTreasury)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewTreasury)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewTreasury represents a NewTreasury event raised by the CakePool contract.
type CakePoolNewTreasury struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewTreasury is a free log retrieval operation binding the contract event 0xafa147634b29e2c7bd53ce194256b9f41cfb9ba3036f2b822fdd1d965beea086.
//
// Solidity: event NewTreasury(address treasury)
func (_CakePool *CakePoolFilterer) FilterNewTreasury(opts *bind.FilterOpts) (*CakePoolNewTreasuryIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewTreasury")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewTreasuryIterator{contract: _CakePool.contract, event: "NewTreasury", logs: logs, sub: sub}, nil
}

// WatchNewTreasury is a free log subscription operation binding the contract event 0xafa147634b29e2c7bd53ce194256b9f41cfb9ba3036f2b822fdd1d965beea086.
//
// Solidity: event NewTreasury(address treasury)
func (_CakePool *CakePoolFilterer) WatchNewTreasury(opts *bind.WatchOpts, sink chan<- *CakePoolNewTreasury) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewTreasury")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewTreasury)
				if err := _CakePool.contract.UnpackLog(event, "NewTreasury", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasury is a log parse operation binding the contract event 0xafa147634b29e2c7bd53ce194256b9f41cfb9ba3036f2b822fdd1d965beea086.
//
// Solidity: event NewTreasury(address treasury)
func (_CakePool *CakePoolFilterer) ParseNewTreasury(log types.Log) (*CakePoolNewTreasury, error) {
	event := new(CakePoolNewTreasury)
	if err := _CakePool.contract.UnpackLog(event, "NewTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewUnlockFreeDurationIterator is returned from FilterNewUnlockFreeDuration and is used to iterate over the raw logs and unpacked data for NewUnlockFreeDuration events raised by the CakePool contract.
type CakePoolNewUnlockFreeDurationIterator struct {
	Event *CakePoolNewUnlockFreeDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewUnlockFreeDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewUnlockFreeDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewUnlockFreeDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewUnlockFreeDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewUnlockFreeDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewUnlockFreeDuration represents a NewUnlockFreeDuration event raised by the CakePool contract.
type CakePoolNewUnlockFreeDuration struct {
	UnlockFreeDuration *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewUnlockFreeDuration is a free log retrieval operation binding the contract event 0xf84bf2b901cfc02956d4e69556d7448cef4ea13587e7714dba7c6d697091e7ad.
//
// Solidity: event NewUnlockFreeDuration(uint256 unlockFreeDuration)
func (_CakePool *CakePoolFilterer) FilterNewUnlockFreeDuration(opts *bind.FilterOpts) (*CakePoolNewUnlockFreeDurationIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewUnlockFreeDuration")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewUnlockFreeDurationIterator{contract: _CakePool.contract, event: "NewUnlockFreeDuration", logs: logs, sub: sub}, nil
}

// WatchNewUnlockFreeDuration is a free log subscription operation binding the contract event 0xf84bf2b901cfc02956d4e69556d7448cef4ea13587e7714dba7c6d697091e7ad.
//
// Solidity: event NewUnlockFreeDuration(uint256 unlockFreeDuration)
func (_CakePool *CakePoolFilterer) WatchNewUnlockFreeDuration(opts *bind.WatchOpts, sink chan<- *CakePoolNewUnlockFreeDuration) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewUnlockFreeDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewUnlockFreeDuration)
				if err := _CakePool.contract.UnpackLog(event, "NewUnlockFreeDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewUnlockFreeDuration is a log parse operation binding the contract event 0xf84bf2b901cfc02956d4e69556d7448cef4ea13587e7714dba7c6d697091e7ad.
//
// Solidity: event NewUnlockFreeDuration(uint256 unlockFreeDuration)
func (_CakePool *CakePoolFilterer) ParseNewUnlockFreeDuration(log types.Log) (*CakePoolNewUnlockFreeDuration, error) {
	event := new(CakePoolNewUnlockFreeDuration)
	if err := _CakePool.contract.UnpackLog(event, "NewUnlockFreeDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewVCakeContractIterator is returned from FilterNewVCakeContract and is used to iterate over the raw logs and unpacked data for NewVCakeContract events raised by the CakePool contract.
type CakePoolNewVCakeContractIterator struct {
	Event *CakePoolNewVCakeContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewVCakeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewVCakeContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewVCakeContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewVCakeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewVCakeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewVCakeContract represents a NewVCakeContract event raised by the CakePool contract.
type CakePoolNewVCakeContract struct {
	VCake common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewVCakeContract is a free log retrieval operation binding the contract event 0x5352e27b0414343d9438a1c6e9d04c65c7cb4d91f44920afee588f91717893f1.
//
// Solidity: event NewVCakeContract(address VCake)
func (_CakePool *CakePoolFilterer) FilterNewVCakeContract(opts *bind.FilterOpts) (*CakePoolNewVCakeContractIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewVCakeContract")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewVCakeContractIterator{contract: _CakePool.contract, event: "NewVCakeContract", logs: logs, sub: sub}, nil
}

// WatchNewVCakeContract is a free log subscription operation binding the contract event 0x5352e27b0414343d9438a1c6e9d04c65c7cb4d91f44920afee588f91717893f1.
//
// Solidity: event NewVCakeContract(address VCake)
func (_CakePool *CakePoolFilterer) WatchNewVCakeContract(opts *bind.WatchOpts, sink chan<- *CakePoolNewVCakeContract) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewVCakeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewVCakeContract)
				if err := _CakePool.contract.UnpackLog(event, "NewVCakeContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVCakeContract is a log parse operation binding the contract event 0x5352e27b0414343d9438a1c6e9d04c65c7cb4d91f44920afee588f91717893f1.
//
// Solidity: event NewVCakeContract(address VCake)
func (_CakePool *CakePoolFilterer) ParseNewVCakeContract(log types.Log) (*CakePoolNewVCakeContract, error) {
	event := new(CakePoolNewVCakeContract)
	if err := _CakePool.contract.UnpackLog(event, "NewVCakeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewWithdrawFeeIterator is returned from FilterNewWithdrawFee and is used to iterate over the raw logs and unpacked data for NewWithdrawFee events raised by the CakePool contract.
type CakePoolNewWithdrawFeeIterator struct {
	Event *CakePoolNewWithdrawFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewWithdrawFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewWithdrawFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewWithdrawFee represents a NewWithdrawFee event raised by the CakePool contract.
type CakePoolNewWithdrawFee struct {
	WithdrawFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFee is a free log retrieval operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 withdrawFee)
func (_CakePool *CakePoolFilterer) FilterNewWithdrawFee(opts *bind.FilterOpts) (*CakePoolNewWithdrawFeeIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewWithdrawFeeIterator{contract: _CakePool.contract, event: "NewWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFee is a free log subscription operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 withdrawFee)
func (_CakePool *CakePoolFilterer) WatchNewWithdrawFee(opts *bind.WatchOpts, sink chan<- *CakePoolNewWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewWithdrawFee)
				if err := _CakePool.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewWithdrawFee is a log parse operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 withdrawFee)
func (_CakePool *CakePoolFilterer) ParseNewWithdrawFee(log types.Log) (*CakePoolNewWithdrawFee, error) {
	event := new(CakePoolNewWithdrawFee)
	if err := _CakePool.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewWithdrawFeeContractIterator is returned from FilterNewWithdrawFeeContract and is used to iterate over the raw logs and unpacked data for NewWithdrawFeeContract events raised by the CakePool contract.
type CakePoolNewWithdrawFeeContractIterator struct {
	Event *CakePoolNewWithdrawFeeContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewWithdrawFeeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewWithdrawFeeContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewWithdrawFeeContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewWithdrawFeeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewWithdrawFeeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewWithdrawFeeContract represents a NewWithdrawFeeContract event raised by the CakePool contract.
type CakePoolNewWithdrawFeeContract struct {
	WithdrawFeeContract *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFeeContract is a free log retrieval operation binding the contract event 0xcab352e118188b8a2f20a2e8c4ce1241ce2c1740aac4f17c5b0831e65824d8ef.
//
// Solidity: event NewWithdrawFeeContract(uint256 withdrawFeeContract)
func (_CakePool *CakePoolFilterer) FilterNewWithdrawFeeContract(opts *bind.FilterOpts) (*CakePoolNewWithdrawFeeContractIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewWithdrawFeeContract")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewWithdrawFeeContractIterator{contract: _CakePool.contract, event: "NewWithdrawFeeContract", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFeeContract is a free log subscription operation binding the contract event 0xcab352e118188b8a2f20a2e8c4ce1241ce2c1740aac4f17c5b0831e65824d8ef.
//
// Solidity: event NewWithdrawFeeContract(uint256 withdrawFeeContract)
func (_CakePool *CakePoolFilterer) WatchNewWithdrawFeeContract(opts *bind.WatchOpts, sink chan<- *CakePoolNewWithdrawFeeContract) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewWithdrawFeeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewWithdrawFeeContract)
				if err := _CakePool.contract.UnpackLog(event, "NewWithdrawFeeContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewWithdrawFeeContract is a log parse operation binding the contract event 0xcab352e118188b8a2f20a2e8c4ce1241ce2c1740aac4f17c5b0831e65824d8ef.
//
// Solidity: event NewWithdrawFeeContract(uint256 withdrawFeeContract)
func (_CakePool *CakePoolFilterer) ParseNewWithdrawFeeContract(log types.Log) (*CakePoolNewWithdrawFeeContract, error) {
	event := new(CakePoolNewWithdrawFeeContract)
	if err := _CakePool.contract.UnpackLog(event, "NewWithdrawFeeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolNewWithdrawFeePeriodIterator is returned from FilterNewWithdrawFeePeriod and is used to iterate over the raw logs and unpacked data for NewWithdrawFeePeriod events raised by the CakePool contract.
type CakePoolNewWithdrawFeePeriodIterator struct {
	Event *CakePoolNewWithdrawFeePeriod // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolNewWithdrawFeePeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolNewWithdrawFeePeriod)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolNewWithdrawFeePeriod)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolNewWithdrawFeePeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolNewWithdrawFeePeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolNewWithdrawFeePeriod represents a NewWithdrawFeePeriod event raised by the CakePool contract.
type CakePoolNewWithdrawFeePeriod struct {
	WithdrawFeePeriod *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFeePeriod is a free log retrieval operation binding the contract event 0xb89ddaddb7435be26824cb48d2d0186c9525a2e1ec057abcb502704cdc0686cc.
//
// Solidity: event NewWithdrawFeePeriod(uint256 withdrawFeePeriod)
func (_CakePool *CakePoolFilterer) FilterNewWithdrawFeePeriod(opts *bind.FilterOpts) (*CakePoolNewWithdrawFeePeriodIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "NewWithdrawFeePeriod")
	if err != nil {
		return nil, err
	}
	return &CakePoolNewWithdrawFeePeriodIterator{contract: _CakePool.contract, event: "NewWithdrawFeePeriod", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFeePeriod is a free log subscription operation binding the contract event 0xb89ddaddb7435be26824cb48d2d0186c9525a2e1ec057abcb502704cdc0686cc.
//
// Solidity: event NewWithdrawFeePeriod(uint256 withdrawFeePeriod)
func (_CakePool *CakePoolFilterer) WatchNewWithdrawFeePeriod(opts *bind.WatchOpts, sink chan<- *CakePoolNewWithdrawFeePeriod) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "NewWithdrawFeePeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolNewWithdrawFeePeriod)
				if err := _CakePool.contract.UnpackLog(event, "NewWithdrawFeePeriod", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewWithdrawFeePeriod is a log parse operation binding the contract event 0xb89ddaddb7435be26824cb48d2d0186c9525a2e1ec057abcb502704cdc0686cc.
//
// Solidity: event NewWithdrawFeePeriod(uint256 withdrawFeePeriod)
func (_CakePool *CakePoolFilterer) ParseNewWithdrawFeePeriod(log types.Log) (*CakePoolNewWithdrawFeePeriod, error) {
	event := new(CakePoolNewWithdrawFeePeriod)
	if err := _CakePool.contract.UnpackLog(event, "NewWithdrawFeePeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CakePool contract.
type CakePoolOwnershipTransferredIterator struct {
	Event *CakePoolOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolOwnershipTransferred represents a OwnershipTransferred event raised by the CakePool contract.
type CakePoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CakePool *CakePoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CakePoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CakePoolOwnershipTransferredIterator{contract: _CakePool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CakePool *CakePoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CakePoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolOwnershipTransferred)
				if err := _CakePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CakePool *CakePoolFilterer) ParseOwnershipTransferred(log types.Log) (*CakePoolOwnershipTransferred, error) {
	event := new(CakePoolOwnershipTransferred)
	if err := _CakePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the CakePool contract.
type CakePoolPauseIterator struct {
	Event *CakePoolPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolPause represents a Pause event raised by the CakePool contract.
type CakePoolPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CakePool *CakePoolFilterer) FilterPause(opts *bind.FilterOpts) (*CakePoolPauseIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &CakePoolPauseIterator{contract: _CakePool.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CakePool *CakePoolFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *CakePoolPause) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolPause)
				if err := _CakePool.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CakePool *CakePoolFilterer) ParsePause(log types.Log) (*CakePoolPause, error) {
	event := new(CakePoolPause)
	if err := _CakePool.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CakePool contract.
type CakePoolPausedIterator struct {
	Event *CakePoolPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolPaused represents a Paused event raised by the CakePool contract.
type CakePoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CakePool *CakePoolFilterer) FilterPaused(opts *bind.FilterOpts) (*CakePoolPausedIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CakePoolPausedIterator{contract: _CakePool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CakePool *CakePoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CakePoolPaused) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolPaused)
				if err := _CakePool.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CakePool *CakePoolFilterer) ParsePaused(log types.Log) (*CakePoolPaused, error) {
	event := new(CakePoolPaused)
	if err := _CakePool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the CakePool contract.
type CakePoolUnlockIterator struct {
	Event *CakePoolUnlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolUnlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolUnlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolUnlock represents a Unlock event raised by the CakePool contract.
type CakePoolUnlock struct {
	Sender         common.Address
	Amount         *big.Int
	BlockTimestamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address indexed sender, uint256 amount, uint256 blockTimestamp)
func (_CakePool *CakePoolFilterer) FilterUnlock(opts *bind.FilterOpts, sender []common.Address) (*CakePoolUnlockIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Unlock", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakePoolUnlockIterator{contract: _CakePool.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address indexed sender, uint256 amount, uint256 blockTimestamp)
func (_CakePool *CakePoolFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *CakePoolUnlock, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Unlock", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolUnlock)
				if err := _CakePool.contract.UnpackLog(event, "Unlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlock is a log parse operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address indexed sender, uint256 amount, uint256 blockTimestamp)
func (_CakePool *CakePoolFilterer) ParseUnlock(log types.Log) (*CakePoolUnlock, error) {
	event := new(CakePoolUnlock)
	if err := _CakePool.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the CakePool contract.
type CakePoolUnpauseIterator struct {
	Event *CakePoolUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolUnpause represents a Unpause event raised by the CakePool contract.
type CakePoolUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CakePool *CakePoolFilterer) FilterUnpause(opts *bind.FilterOpts) (*CakePoolUnpauseIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &CakePoolUnpauseIterator{contract: _CakePool.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CakePool *CakePoolFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *CakePoolUnpause) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolUnpause)
				if err := _CakePool.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CakePool *CakePoolFilterer) ParseUnpause(log types.Log) (*CakePoolUnpause, error) {
	event := new(CakePoolUnpause)
	if err := _CakePool.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CakePool contract.
type CakePoolUnpausedIterator struct {
	Event *CakePoolUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolUnpaused represents a Unpaused event raised by the CakePool contract.
type CakePoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CakePool *CakePoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CakePoolUnpausedIterator, error) {

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CakePoolUnpausedIterator{contract: _CakePool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CakePool *CakePoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CakePoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolUnpaused)
				if err := _CakePool.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CakePool *CakePoolFilterer) ParseUnpaused(log types.Log) (*CakePoolUnpaused, error) {
	event := new(CakePoolUnpaused)
	if err := _CakePool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakePoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CakePool contract.
type CakePoolWithdrawIterator struct {
	Event *CakePoolWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CakePoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakePoolWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CakePoolWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CakePoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakePoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakePoolWithdraw represents a Withdraw event raised by the CakePool contract.
type CakePoolWithdraw struct {
	Sender common.Address
	Amount *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_CakePool *CakePoolFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*CakePoolWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &CakePoolWithdrawIterator{contract: _CakePool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_CakePool *CakePoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CakePoolWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CakePool.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakePoolWithdraw)
				if err := _CakePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_CakePool *CakePoolFilterer) ParseWithdraw(log types.Log) (*CakePoolWithdraw, error) {
	event := new(CakePoolWithdraw)
	if err := _CakePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
