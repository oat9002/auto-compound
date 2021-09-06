package utils

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// convert wei
func FromWei(data *big.Int) float64 {
	toReturn := float64(data.Int64()) / math.Pow10(18)

	return math.Round(toReturn*math.Pow10(6)) / math.Pow10(6)
}

// Get default CallOpts
func GetDefaultCallOpts(address common.Address) *bind.CallOpts {
	return &bind.CallOpts{Pending: false, From: address, BlockNumber: nil, Context: nil}
}
