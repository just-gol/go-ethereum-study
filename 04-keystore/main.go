package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var quickNodeURL = "https://magical-wild-model.ethereum-sepolia.quiknode.pro/4291e8c1bb9318f2307ff3ff5174826487ef3940/"
var localURL = "http://localhost:8545"

func main() {
	//key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	passwork := "password1"
	//account, err := key.NewAccount(passwork)
	//if err != nil {
	//	log.Fatalf("Error creating account: %v", err)
	//}
	//fmt.Printf("account address: %v\n", account.Address)
	file, err := os.ReadFile("./wallet/UTC--2025-12-22T06-14-39.876452113Z--0e8f9a7e918785d086abee53cf15b77a8cf0982e")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(file, passwork)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private Key:", hex.EncodeToString(pData))

	pub := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public Key:", hexutil.Encode(pub))

	fmt.Println("Public address:", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
