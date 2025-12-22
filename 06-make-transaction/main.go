package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	client, err := ethclient.Dial(localURL)
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

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)
	legacyTx := &types.LegacyTx{
		Nonce:    nonce,
		To:       &a2,
		Value:    amount,
		Gas:      21000,
		GasPrice: gasPrice,
		Data:     []byte{},
	}
	//tx := types.NewTx(legacyTx)
	//fmt.Println("Created Transaction", tx)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.ReadFile("./wallet/UTC--2025-12-22T07-29-15.490482896Z--0303d1f5f03ba235fd91190a998e24eca120708a")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(file, "password")
	if err != nil {
		log.Fatal(err)
	}
	// 这是一个用于签名已构造的交易的方法。通常，它会将一个已构建好的交易对象（types.Transaction）与指定的私钥进行签名。
	//types.SignTx()

	// 这是一个更高层的签名方法，它直接创建一个新的交易并对其进行签名。与 SignTx() 不同，SignNewTx() 内部会处理部分交易的构建。
	/*types.NewEIP155Signer(chainID) 适用于标准的以太坊交易，特别是 Legacy Transaction。
	适用于 普通的以太坊交易，比如：
	发送 ETH 转账
	调用合约（标准的 Legacy 交易）
	*/
	//types.NewEIP2930Signer(chainID) 适用于需要访问列表的交易，通常用于更复杂的交易场景，能够优化 Gas。
	tx, err := types.SignNewTx(key.PrivateKey, types.NewEIP155Signer(chainID), legacyTx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent:%s\n", tx.Hash().Hex())
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
