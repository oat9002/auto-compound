package services

import "github.com/ethereum/go-ethereum/ethclient"

type clientService struct{}

func NewClientService() *clientService {
	clientSerivce := new(clientService)

	return clientSerivce
}

func (c *clientService) GetClient(networkUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(networkUrl)
}
