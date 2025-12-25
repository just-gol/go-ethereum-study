package main

import (
	"log"

	"08-go-solidity-when/handle"
	"08-go-solidity-when/routers"
	"08-go-solidity-when/service"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	whenService := service.NewWhenService(client)
	handler := handle.NewWhenHandler(whenService)

	r := gin.Default()
	r.Use(cors.Default())
	routers.ApiRoutersInit(r, handler)
	_ = r.Run()
}
