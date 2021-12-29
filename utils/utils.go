package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

//Get gas fee
func GetGasFee(client *ethclient.Client, txHash common.Hash) (float64, error) {
	errCh := make(chan error)
	gasFeeCh := make(chan float64)
	qCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-qCh:
				gasFeeCh <- 0
				errCh <- fmt.Errorf(fmt.Sprintf("Transaction %s is still in pending state", txHash))
				return
			default:
				tx, isPending, err := client.TransactionByHash(context.Background(), txHash)

				if err != nil {
					gasFeeCh <- 0
					errCh <- err
					return
				}

				if isPending {
					time.Sleep(5 * time.Second)
					continue
				}

				receipt, err := client.TransactionReceipt(context.Background(), txHash)

				if err != nil {
					gasFeeCh <- 0
					errCh <- err
					return
				}

				gasFeeCh <- math.Round(float64(receipt.GasUsed)*FromWei(tx.GasPrice())*math.Pow10(6)) / math.Pow10(6)
				errCh <- nil
				return
			}
		}
	}()

	go func() {
		time.Sleep(10 * time.Minute)
		qCh <- struct{}{}
	}()

	gasFee := <-gasFeeCh
	err := <-errCh

	return gasFee, err
}
