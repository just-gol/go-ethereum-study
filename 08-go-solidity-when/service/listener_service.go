package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"08-go-solidity-when/models"

	when "08-go-solidity-when/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type ListenerService interface {
	MonitorEvent(ctx context.Context, contractAddress common.Address)
	// ReplayFromLast 从上次同步区块回放到最新确认区块。
	ReplayFromLast(ctx context.Context, contractAddress common.Address, startBlock uint64, confirmations uint64) error
	// StartReplayLoop 定时回放，补齐重启或断线期间的缺失数据。
	StartReplayLoop(ctx context.Context, contractAddress common.Address, startBlock uint64, confirmations uint64, interval time.Duration)
}

type listenerService struct {
	client *ethclient.Client
}

func NewListenerService(client *ethclient.Client) ListenerService {
	return &listenerService{client: client}
}

func (l *listenerService) MonitorEvent(ctx context.Context, contractAddress common.Address) {
	// 实时订阅：使用 abigen 的 Watch* 直接解析事件字段。
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
			l.handleApproval(ev)
		case ev := <-transfers:
			l.handleTransfer(ev)
		case ev := <-deposit:
			l.handleDeposit(ev)
		case ev := <-withdraw:
			l.handleWithdraw(ev)
		}
	}
}

func (l *listenerService) ReplayFromLast(ctx context.Context, contractAddress common.Address, startBlock uint64, confirmations uint64) error {
	// 读取上次同步区块；首次启动则用配置的起始区块。
	key := syncKey(contractAddress)
	last, err := l.getSyncBlock(key)
	if err != nil {
		return err
	}
	if last == 0 && startBlock > 0 {
		last = startBlock - 1
	}

	// 仅回放到最新确认区块，避免链重组带来的回滚。
	latestHeader, err := l.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	latest := latestHeader.Number.Uint64()
	if confirmations > 1 && latest >= confirmations-1 {
		latest -= confirmations - 1
	}

	if last >= latest {
		return nil
	}

	// 回放区间：[last+1, latest]
	return l.replayRange(ctx, contractAddress, last+1, latest)
}

func (l *listenerService) StartReplayLoop(ctx context.Context, contractAddress common.Address, startBlock uint64, confirmations uint64, interval time.Duration) {
	// 定时补数据：防止 WS 断线或服务重启导致漏事件。
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := l.ReplayFromLast(ctx, contractAddress, startBlock, confirmations); err != nil {
				log.Println("replay loop error:", err)
			}
		}
	}
}

func (l *listenerService) replayRange(ctx context.Context, contractAddress common.Address, start uint64, end uint64) error {
	// 历史回放：按区块范围 Filter* 批量拉取事件。
	w, err := when.NewWhen(contractAddress, l.client)
	if err != nil {
		return err
	}
	endCopy := end

	approvalIter, err := w.FilterApproval(&bind.FilterOpts{Start: start, End: &endCopy, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer approvalIter.Close()
	for approvalIter.Next() {
		l.handleApproval(approvalIter.Event)
	}
	if err := approvalIter.Error(); err != nil {
		return err
	}

	transferIter, err := w.FilterTransfer(&bind.FilterOpts{Start: start, End: &endCopy, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer transferIter.Close()
	for transferIter.Next() {
		l.handleTransfer(transferIter.Event)
	}
	if err := transferIter.Error(); err != nil {
		return err
	}

	depositIter, err := w.FilterDeposit(&bind.FilterOpts{Start: start, End: &endCopy, Context: ctx}, nil)
	if err != nil {
		return err
	}
	defer depositIter.Close()
	for depositIter.Next() {
		l.handleDeposit(depositIter.Event)
	}
	if err := depositIter.Error(); err != nil {
		return err
	}

	withdrawIter, err := w.FilterWithdrawal(&bind.FilterOpts{Start: start, End: &endCopy, Context: ctx}, nil)
	if err != nil {
		return err
	}
	defer withdrawIter.Close()
	for withdrawIter.Next() {
		l.handleWithdraw(withdrawIter.Event)
	}
	if err := withdrawIter.Error(); err != nil {
		return err
	}
	// 即使 end 这段区间没有任何事件，也把进度推进到 end，避免下次重复回放空区块
	return l.setSyncBlock(syncKey(contractAddress), end)
}

func (l *listenerService) handleApproval(ev *when.WhenApproval) {
	if ev == nil {
		return
	}
	// 幂等入库：txHash + logIndex 唯一。
	ok, err := l.recordEvent(ev.Raw, "Approval")
	if err != nil || !ok {
		return
	}
	fmt.Printf("Approval src=%s guy=%s wad=%s\n", ev.Src.Hex(), ev.Guy.Hex(), ev.Wad.String())
	approval := models.Approval{
		Src: ev.Src.String(),
		Guy: ev.Guy.String(),
		Wad: ev.Wad.String(),
	}
	models.DB.Create(&approval)
	_ = l.setSyncBlock(syncKey(ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) handleTransfer(ev *when.WhenTransfer) {
	if ev == nil {
		return
	}
	// 幂等入库：txHash + logIndex 唯一。
	ok, err := l.recordEvent(ev.Raw, "Transfer")
	if err != nil || !ok {
		return
	}
	fmt.Printf("Transfer src=%s dst=%s wad=%s\n", ev.Src.Hex(), ev.Dst.Hex(), ev.Wad.String())
	transfer := models.Transfer{
		Src: ev.Src.String(),
		Dst: ev.Dst.String(),
		Wad: ev.Wad.String(),
	}
	models.DB.Create(&transfer)
	_ = l.setSyncBlock(syncKey(ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) handleDeposit(ev *when.WhenDeposit) {
	if ev == nil {
		return
	}
	// 幂等入库：txHash + logIndex 唯一。
	ok, err := l.recordEvent(ev.Raw, "Deposit")
	if err != nil || !ok {
		return
	}
	fmt.Printf("Deposit dst=%s wad=%s\n", ev.Dst.Hex(), ev.Wad.String())
	deposit := models.Deposit{
		Dst: ev.Dst.String(),
		Wad: ev.Wad.String(),
	}
	models.DB.Create(&deposit)
	_ = l.setSyncBlock(syncKey(ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) handleWithdraw(ev *when.WhenWithdrawal) {
	if ev == nil {
		return
	}
	// 幂等入库：txHash + logIndex 唯一。
	ok, err := l.recordEvent(ev.Raw, "Withdrawal")
	if err != nil || !ok {
		return
	}
	fmt.Printf("Withdraw src=%s wad=%s\n", ev.Src.Hex(), ev.Wad.String())
	withdraw := models.Withdraw{
		Src: ev.Src.String(),
		Wad: ev.Wad.String(),
	}
	models.DB.Create(&withdraw)
	_ = l.setSyncBlock(syncKey(ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) recordEvent(logEntry types.Log, eventName string) (bool, error) {
	// FirstOrCreate 保证回放和实时订阅不会重复落库。
	entry := models.EventLog{
		TxHash:      logEntry.TxHash.Hex(),
		LogIndex:    logEntry.Index,
		BlockNumber: logEntry.BlockNumber,
		Event:       eventName,
		Contract:    logEntry.Address.Hex(),
	}

	// FirstOrCreate 查到就返回,没查到就写入
	result := models.DB.Where("tx_hash = ? AND log_index = ?", entry.TxHash, entry.LogIndex).FirstOrCreate(&entry)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (l *listenerService) getSyncBlock(key string) (uint64, error) {
	// 同步进度写在 DB 里，重启后可续跑。
	var state models.SyncState
	err := models.DB.Where("name = ?", key).First(&state).Error
	if err == nil {
		return state.BlockNumber, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	return 0, err
}

func (l *listenerService) setSyncBlock(key string, block uint64) error {
	// Upsert 最新已处理区块。
	state := models.SyncState{Name: key, BlockNumber: block}
	//Assign GORM 会在“查到已存在记录”的情况下做更新，没查到就创建;从而达到“更新最新区块号”的目的
	return models.DB.Where("name = ?", key).Assign(models.SyncState{BlockNumber: block}).FirstOrCreate(&state).Error
}

func syncKey(contractAddress common.Address) string {
	// 规范化地址生成稳定的 key。
	return "when:" + strings.ToLower(contractAddress.Hex())
}
