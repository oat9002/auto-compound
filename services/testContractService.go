package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oat9002/auto-compound/contracts"
)

type testContractService struct {
	contract *contracts.MyContract
}

func NewTestContractService(client *ethclient.Client) (*testContractService, error) {
	contract, err := getMyContractContract(client)
	service := &testContractService{contract: contract}

	return service, err
}

func getMyContractContract(client *ethclient.Client) (*contracts.MyContract, error) {
	contractAddress := common.HexToAddress("0x239cf537496d8DC38AD9ed1dE7D37bBCa69417D7")

	return contracts.NewMyContract(contractAddress, client)
}
