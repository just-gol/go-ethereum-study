package main

import (
	todo "07-go-solidity-learn/gen"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var quickNodeURL = "https://magical-wild-model.ethereum-sepolia.quiknode.pro/4291e8c1bb9318f2307ff3ff5174826487ef3940/"
var localURL = "http://localhost:8545"

func main() {
	file, err := os.ReadFile("./wallet/UTC--2025-12-22T07-29-15.490482896Z--0303d1f5f03ba235fd91190a998e24eca120708a")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(file, "password")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(quickNodeURL)
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	fmt.Println("Chain ID:", chainID)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x6196F8D81833457b3cef9bCBeD18B823b682A916")
	tx, err := todo.NewTodo(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	t.GasPrice = gasPrice
	t.GasLimit = 3000000
	add, err := tx.Add(t, "First task")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("add:", add.Hash())
	pubkeyAddress := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	tasks, err := tx.List(&bind.CallOpts{
		From: pubkeyAddress,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list:", tasks)

	//update, err := tx.Update(t, big.NewInt(0), "update task content")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("update:", update.Hash())
}
