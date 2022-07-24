package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ClientService struct{}

func NewClientService() *ClientService {
	clientSerivce := &ClientService{}

	return clientSerivce
}

func (c *ClientService) GetClient(networkUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(networkUrl)

	if err != nil {
		return nil, fmt.Errorf("cannot create eth client, %w", err)
	}

	return client, nil
}
