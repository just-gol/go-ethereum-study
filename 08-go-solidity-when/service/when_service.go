package service

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	when "08-go-solidity-when/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type WhenService interface {
	GetBalance(ctx context.Context, contractAddress, owner common.Address) (*big.Int, error)
	Allowance(ctx context.Context, contractAddress, owner, spender common.Address) (*big.Int, error)
	TransferFrom(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, dst common.Address, amount *big.Int) (*types.Transaction, error)
	Deposit(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error)
	Approve(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, spender common.Address, amount *big.Int) (*types.Transaction, error)
	Withdraw(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error)
}

type whenService struct {
	client *ethclient.Client
}

func NewWhenService(client *ethclient.Client) WhenService {
	return &whenService{client: client}
}

func (s *whenService) GetBalance(ctx context.Context, contractAddress, owner common.Address) (*big.Int, error) {
	w, err := when.NewWhen(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return w.BalanceOf(&bind.CallOpts{Context: ctx}, owner)
}

func (s *whenService) Allowance(ctx context.Context, contractAddress, owner, spender common.Address) (*big.Int, error) {
	w, err := when.NewWhen(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return w.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
}

func (s *whenService) TransferFrom(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	w, err := when.NewWhen(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	chainID, err := s.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	nonce, err := s.client.PendingNonceAt(ctx, crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return nil, err
	}
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)
	auth.Value = big.NewInt(0)

	return w.TransferFrom(auth, crypto.PubkeyToAddress(privateKey.PublicKey), dst, amount)
}

func (s *whenService) Deposit(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	chainID, err := s.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	nonce, err := s.client.PendingNonceAt(ctx, crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return nil, err
	}
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	legacyTx := &types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddress,
		Value:    amount,
		Gas:      300000,
		GasPrice: gasPrice,
		Data:     nil,
	}
	tx, err := types.SignNewTx(privateKey, types.NewEIP155Signer(chainID), legacyTx)
	if err != nil {
		return nil, err
	}
	if err := s.client.SendTransaction(ctx, tx); err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *whenService) Approve(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	w, err := when.NewWhen(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	chainID, err := s.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return w.Approve(auth, spender, amount)
}

func (s *whenService) Withdraw(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	w, err := when.NewWhen(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	chainID, err := s.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return w.Withdraw(auth, amount)
}
