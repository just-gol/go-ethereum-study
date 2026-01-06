package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var localURL = "http://localhost:8545"

func main() {
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Block number: %d\n", blockNumber)
}
