package main

import (
	"context"
	"log"

	"08-go-solidity-when/config"
	"08-go-solidity-when/handle"
	"08-go-solidity-when/routers"
	"08-go-solidity-when/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		log.Fatal(err)
	}
	whenService := service.NewWhenService(client)
	handler := handle.NewWhenHandler(whenService, cfg.ContractAddress)
	txService := service.NewTxService(client)
	txHandler := handle.NewTxHandler(txService)
	approvalService := service.NewApprovalService()
	approvalHandle := handle.NewApprovalHandle(approvalService)

	wsClient, err := ethclient.Dial(cfg.WSURL)
	if err != nil {
		log.Fatal(err)
	}
	listenerSvc := service.NewListenerService(wsClient)
	if cfg.ContractAddress != "" {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		go func() {
			listenerSvc.MonitorEvent(context.Background(), contractAddress)
		}()
	}
	r := gin.Default()
	r.Use(cors.Default())
	routers.ApiRoutersInit(r, handler, txHandler, approvalHandle)
	_ = r.Run()
}
