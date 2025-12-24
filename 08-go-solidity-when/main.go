package main

import (
	when "08-go-solidity-when/gen"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var localURL = "http://localhost:8545"

func main() {
	privateKeyString := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyString[2:])
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatal(err)
	}
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		log.Fatal(err)
	}
	// 合约地址
	address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	// 创建实例
	newWhen, err := when.NewWhen(address, client)
	if err != nil {
		log.Fatal(err)
	}

	dst := common.HexToAddress("0303d1f5f03ba235fd91190a998e24eca120708a")
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)
	auth.Value = big.NewInt(0)
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	transaction, err := newWhen.TransferFrom(auth, crypto.PubkeyToAddress(privateKey.PublicKey), dst, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction: %v\n", transaction)

	// 查询地址余额
	balanceOf, err := newWhen.BalanceOf(&bind.CallOpts{}, dst)
	if err != nil {
		if revertErr, ok := err.(interface{ ErrorData() string }); ok {
			data := revertErr.ErrorData()
			fmt.Println("Revert data:", data)
		}
		log.Fatal(err)
	}
	fmt.Printf("PubkeyToAddress: %v,BalanceOf: %s\n", dst.Hex(), balanceOf)
}
