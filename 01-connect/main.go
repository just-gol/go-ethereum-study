package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

var quickNodeURL = "https://magical-wild-model.ethereum-sepolia.quiknode.pro/4291e8c1bb9318f2307ff3ff5174826487ef3940/"
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
