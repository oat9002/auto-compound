package services

import "github.com/ethereum/go-ethereum/ethclient"

type ClientService struct{}

func NewClientService() *ClientService {
	clientSerivce := &ClientService{}

	return clientSerivce
}

func (c *ClientService) GetClient(networkUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(networkUrl)
}
