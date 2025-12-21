package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

var quickNodeURL = "https://magical-wild-model.ethereum-sepolia.quiknode.pro/4291e8c1bb9318f2307ff3ff5174826487ef3940/"
var localURL = "http://localhost:8545"

func main() {
	client, err := ethclient.Dial(quickNodeURL)
	if err != nil {
		log.Fatalf("Error to connect:%v", err)
	}
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Error to get a blockNum:%v", err)
	}
	fmt.Printf("the Block number: %d\n", blockNumber)

	addr := "0x00640f45FEC08927A437EA023CfA13Cb7354d926"
	address := common.HexToAddress(addr)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance:%v", err)
	}
	fmt.Printf("balance: %v\n", balance)
	// 1eth  =10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Printf("fBalance: %v\n", fBalance)
	// 10 * 10 ...18
	balanceETH := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("balanceETH: %v\n", balanceETH)

}
