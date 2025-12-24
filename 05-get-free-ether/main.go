package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var quickNodeURL = "https://magical-wild-model.ethereum-sepolia.quiknode.pro/4291e8c1bb9318f2307ff3ff5174826487ef3940/"
var localURL = "http://localhost:8545"

func main() {
	//key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	//passwork := "password"
	//_, err := key.NewAccount(passwork)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = key.NewAccount(passwork)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 0303d1f5f03ba235fd91190a998e24eca120708a
	// 838089eabffb00a1ce873705173771475cc86b3c
	client, err := ethclient.Dial(quickNodeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("0303d1f5f03ba235fd91190a998e24eca120708a")
	a2 := common.HexToAddress("838089eabffb00a1ce873705173771475cc86b3c")
	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}
	fb1 := new(big.Float)
	fb1.SetString(b1.String())
	fmt.Println("balance1", new(big.Float).Quo(fb1, big.NewFloat(math.Pow10(18))))
	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}
	fb2 := new(big.Float)
	fb2.SetString(b2.String())
	fmt.Println("balance2", new(big.Float).Quo(fb2, big.NewFloat(math.Pow10(18))))
	//getBlockInfo(1, client)
}

func getBlockInfo(blockNumber int64, client *ethclient.Client) {
	newInt := big.NewInt(blockNumber)
	block, err := client.BlockByNumber(context.Background(), newInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("blockNumber:", block.Number().Uint64())
	fmt.Println("time:", block.Time())
	fmt.Println("difficulty:", block.Difficulty().Uint64())
	fmt.Println("hash:", block.Hash().Hex())
	fmt.Println("transactionsLen:", len(block.Transactions()))
}
