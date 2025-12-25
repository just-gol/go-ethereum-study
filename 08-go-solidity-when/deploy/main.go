package main

import (
	when "08-go-solidity-when/gen"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var localURL = "http://localhost:8545"

func main() {
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyStr := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyStr[2:])
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(3000000)
	deployWhen, transaction, _, err := when.DeployWhen(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deployWhen.Hex())
	fmt.Println(transaction.Hash().Hex())
}
