package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

// Get default TransactionOpts
func GetDefautlTransactionOpts(client *ethclient.Client, privateKeyStr string, chainId uint64, gasLimit uint64, gasPriceThreshold uint64) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	} else if gasPrice.Uint64() > gasPriceThreshold {
		return nil, fmt.Errorf("Gas price is above threshold, gas price: " + gasPrice.Text(10) + " wei")
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))
	if err != nil {
		return nil, err
	}

	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = gasLimit
	txOpts.GasPrice = gasPrice

	return txOpts, nil
}

func GetGasFree(client *ethclient.Client, tx *types.Transaction) (float64, error) {
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())

	if err != nil {
		panic("")
	}

	return math.Round(float64(receipt.GasUsed)*FromWei(tx.GasPrice())*math.Pow10(6)) / math.Pow10(6), nil
}
