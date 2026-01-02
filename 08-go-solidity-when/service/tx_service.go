package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TxService interface {
	SendValueAndTrack(ctx context.Context, privateKey *ecdsa.PrivateKey, to common.Address, value *big.Int, confirmations uint64) (*types.Transaction, *types.Receipt, error)
}

type txService struct {
	client *ethclient.Client
}

func NewTxService(client *ethclient.Client) TxService {
	return &txService{client: client}
}

func (s *txService) SendValueAndTrack(ctx context.Context, privateKey *ecdsa.PrivateKey, to common.Address, value *big.Int, confirmations uint64) (*types.Transaction, *types.Receipt, error) {
	if privateKey == nil {
		return nil, nil, errors.New("private key is required")
	}
	if value == nil || value.Sign() < 0 {
		return nil, nil, errors.New("value must be >= 0")
	}

	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := s.client.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, nil, err
	}
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, err
	}
	chainID, err := s.client.ChainID(ctx)
	if err != nil {
		return nil, nil, err
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      21000,
		GasPrice: gasPrice,
	})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, nil, err
	}
	if err := s.client.SendTransaction(ctx, signedTx); err != nil {
		return nil, nil, err
	}

	receipt, err := waitReceipt(ctx, s.client, signedTx.Hash())
	if err != nil {
		return signedTx, nil, err
	}

	if confirmations > 0 {
		if err := waitConfirmations(ctx, s.client, receipt.BlockNumber, confirmations); err != nil {
			return signedTx, receipt, err
		}
	}

	return signedTx, receipt, nil
}

/*
waitReceipt 等“上链”
发送交易后不会立刻上链，receipt 只有在交易被打包后才会有。
这个函数用循环查询 TransactionReceipt：
没找到（pending）就等待一会儿再查
找到就返回 receipt（含成功/失败状态、区块号等）
*/
func waitReceipt(ctx context.Context, client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		receipt, err := client.TransactionReceipt(ctx, hash)
		if err == nil {
			return receipt, nil
		}
		if !errors.Is(err, ethereum.NotFound) {
			return nil, err
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
		}
	}
}

/*
waitConfirmations 等“更稳一点”
交易刚打包可能会被“链重组”回滚，所以很多系统会等 N 个确认数。
它就是不断查最新区块号，直到达到 receipt.BlockNumber + confirmations - 1 才返回。
*/
func waitConfirmations(ctx context.Context, client *ethclient.Client, blockNumber *big.Int, confirmations uint64) error {
	if confirmations == 0 {
		return nil
	}
	target := new(big.Int).Add(blockNumber, new(big.Int).SetUint64(confirmations-1))
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return err
		}
		if header.Number.Cmp(target) >= 0 {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}
