package main

import (
	todo "07-go-solidity-learn/gen"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
	client, err := ethclient.Dial("quickNodeURL")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(3000000)
	deployTodo, transaction, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deployTodo.Hex())
	fmt.Println(transaction.Hash().Hex())
}
