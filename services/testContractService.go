package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oat9002/auto-compound/contracts"
	"github.com/oat9002/auto-compound/utils"
)

type TestContractService struct {
	client   *ethclient.Client
	contract *contracts.MyContract
	chainId  uint64
}

func NewTestContractService(client *ethclient.Client, chainId uint64) (*TestContractService, error) {
	contract, err := getMyContractContract(client)
	service := &TestContractService{
		client:   client,
		contract: contract,
		chainId:  chainId,
	}

	return service, err
}

func getMyContractContract(client *ethclient.Client) (*contracts.MyContract, error) {
	contractAddress := common.HexToAddress("0x239cf537496d8DC38AD9ed1dE7D37bBCa69417D7")

	return contracts.NewMyContract(contractAddress, client)
}

func (t *TestContractService) Increase(privateKey string) (*types.Transaction, error) {
	txOpts, err := utils.GetDefautlTransactionOpts(t.client, privateKey, t.chainId)

	if err != nil {
		return nil, err
	}

	transaction, err := t.contract.Increase(txOpts)

	return transaction, err
}
