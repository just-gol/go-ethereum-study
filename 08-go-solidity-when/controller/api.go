package controller

import (
	when "08-go-solidity-when/gen"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
}

var localURL = "http://localhost:8545"

func (con ApiController) GetBalance(ctx *gin.Context) {
	//privateKeyString := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKeyString := ctx.Query("privateKeyString")

	privateKey, err := crypto.HexToECDSA(privateKeyString[2:])
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatal(err)
	}
	// 合约地址 0x5FbDB2315678afecb367f032d93F642f64180aa3
	contractAddress := ctx.Query("contractAddress")
	address := common.HexToAddress(contractAddress)
	// 创建实例
	newWhen, err := when.NewWhen(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// 当前调用者的地址
	dst := crypto.PubkeyToAddress(privateKey.PublicKey)
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
	ctx.JSON(200, gin.H{
		"balanceOf": balanceOf,
	})
}
func (con ApiController) TransferFrom(ctx *gin.Context) {
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatal(err)
	}
	//privateKeyString := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKeyString := ctx.PostForm("privateKeyString")
	// 合约地址 0x5FbDB2315678afecb367f032d93F642f64180aa3
	contractAddress := ctx.PostForm("contractAddress")

	privateKey, err := crypto.HexToECDSA(privateKeyString[2:])
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	newWhen, err := when.NewWhen(address, client)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)
	auth.Value = big.NewInt(0)
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	// 目标地址
	dst := common.HexToAddress(ctx.PostForm("dst"))
	transaction, err := newWhen.TransferFrom(auth, crypto.PubkeyToAddress(privateKey.PublicKey), dst, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction: %v\n", transaction)
}

func (con ApiController) Deposit(ctx *gin.Context) {
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyString := ctx.PostForm("privateKeyString")
	privateKey, err := crypto.HexToECDSA(privateKeyString[2:])
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		log.Fatal(err)
	}
	depositAmount := big.NewInt(1000000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	legacyTx := &types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddress,
		Value:    depositAmount,
		Gas:      300000,
		GasPrice: gasPrice,
		Data:     nil,
	}
	tx, err := types.SignNewTx(privateKey, types.NewEIP155Signer(chainID), legacyTx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent:%s\n", tx.Hash().Hex())
}

func (con ApiController) Approve(ctx *gin.Context) {
	client, err := ethclient.Dial(localURL)
	if err != nil {
		log.Fatal(err)
	}
	//privateKeyString := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKeyString := ctx.PostForm("privateKeyString")
	// 合约地址 0x5FbDB2315678afecb367f032d93F642f64180aa3
	contractAddress := ctx.PostForm("contractAddress")
	address := common.HexToAddress(contractAddress)
	newWhen, err := when.NewWhen(address, client)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(privateKeyString[2:])
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	to := ctx.PostForm("guy")
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	approve, err := newWhen.Approve(transactOpts, common.HexToAddress(to), amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Approve: %v\n", approve.Hash().Hex())
}
