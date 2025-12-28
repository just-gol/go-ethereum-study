package service

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ListenerService interface {
	MonitorEvent(ctx context.Context, contractAddress common.Address)
}

type listenerService struct {
	client *ethclient.Client
}

func NewListenerService(client *ethclient.Client) ListenerService {
	return &listenerService{client: client}
}

func (l *listenerService) MonitorEvent(ctx context.Context, contractAddress common.Address) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := l.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("error:", err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}
