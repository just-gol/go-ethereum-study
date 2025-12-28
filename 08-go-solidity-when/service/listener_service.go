package service

import (
	"context"
	"fmt"
	"log"

	when "08-go-solidity-when/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	// SubscribeFilterLogs
	//拿到的是原始 types.Log（topic + data）。
	//需要你自己用 ABI 解析事件。
	/**
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
	*/
	// WatchApproval 直接拿到结构化事件 WhenApproval（字段已解析）。
	w, err := when.NewWhen(contractAddress, l.client)
	if err != nil {
		log.Println("listener init error:", err)
		return
	}
	sink := make(chan *when.WhenApproval)
	subApproval, err := w.WatchApproval(&bind.WatchOpts{Context: ctx}, sink, nil, nil)
	if err != nil {
		log.Println("watch approval error:", err)
		return
	}
	defer subApproval.Unsubscribe()
	transfers := make(chan *when.WhenTransfer)
	subTransaction, err := w.WatchTransfer(&bind.WatchOpts{Context: ctx}, transfers, nil, nil)
	if err != nil {
		log.Println("watch transfer error:", err)
		return
	}
	defer subTransaction.Unsubscribe()
	deposit := make(chan *when.WhenDeposit)
	subDeposit, err := w.WatchDeposit(&bind.WatchOpts{Context: ctx}, deposit, nil)
	if err != nil {
		log.Println("watch deposit error:", err)
		return
	}
	defer subDeposit.Unsubscribe()
	withdraw := make(chan *when.WhenWithdrawal)
	subWithdraw, err := w.WatchWithdrawal(&bind.WatchOpts{Context: ctx}, withdraw, nil)
	if err != nil {
		log.Println("watch withdraw error:", err)
		return
	}
	defer subWithdraw.Unsubscribe()
	for {
		select {
		case <-ctx.Done():
			log.Println("listener stopped:", ctx.Err())
			return
		case err := <-subApproval.Err():
			log.Println("approval sub error:", err)
			return
		case err := <-subTransaction.Err():
			log.Println("transfer sub error:", err)
			return
		case err := <-subDeposit.Err():
			log.Println("deposit sub error:", err)
			return
		case err := <-subWithdraw.Err():
			log.Println("withdraw sub error:", err)
			return
		case ev := <-sink:
			fmt.Printf("Approval src=%s guy=%s wad=%s\n", ev.Src.Hex(), ev.Guy.Hex(), ev.Wad.String())
			//approval := models.Approval{
			//	Src: ev.Src,
			//	Guy: ev.Guy,
			//	Wad: ev.Wad,
			//}
		case ev := <-transfers:
			//  event Transfer(address indexed src, address indexed dst, uint wad);
			fmt.Printf("Transfer src=%s dst=%s wad=%s\n", ev.Src.Hex(), ev.Dst.Hex(), ev.Wad.String())
		//transfer := models.Transfer{
		//	Src: ev.Src,
		//	Dst: ev.Dst,
		//	Wad: ev.Wad,
		//}
		case ev := <-deposit:
			fmt.Printf("Deposit dst=%s wad=%s\n", ev.Dst.Hex(), ev.Wad.String())
		case ev := <-withdraw:
			fmt.Printf("withdraw src=%s wad=%s\n", ev.Src.Hex(), ev.Wad.String())

		}
	}
}
