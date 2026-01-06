package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var localURL = "http://localhost:8545"

func main() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(privateKey)
	fmt.Printf("私钥:%v \n", hexutil.Encode(pData))
	// 公钥由私钥生成
	puData := crypto.FromECDSAPub(&privateKey.PublicKey)
	fmt.Printf("公钥:%v \n", hexutil.Encode(puData))
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Printf("公钥地址:%v \n", address)

}
