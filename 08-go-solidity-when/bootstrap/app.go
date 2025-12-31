package bootstrap

import (
	"context"

	"08-go-solidity-when/config"
	"08-go-solidity-when/handle"
	"08-go-solidity-when/routers"
	"08-go-solidity-when/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewApp(cfg config.Config) (*gin.Engine, error) {
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, err
	}
	whenService := service.NewWhenService(client)
	whenHandler := handle.NewWhenHandler(whenService, cfg.ContractAddress)
	transactionService := service.NewTxService(client)
	transactionHandler := handle.NewTxHandler(transactionService)
	approvalService := service.NewApprovalService()
	approvalHandler := handle.NewApprovalHandle(approvalService)
	depositService := service.NewDepositService()
	depositHandle := handle.NewDepositHandle(depositService)
	transferService := service.NewTransferService()
	transferHandle := handle.NewTransferHandle(transferService)
	withdrawService := service.NewWithdrawService()
	withdrawHandle := handle.NewWithdrawHandle(withdrawService)
	wsClient, err := ethclient.Dial(cfg.WSURL)
	if err != nil {
		return nil, err
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
	routers.ApiRoutersInit(r, whenHandler, transactionHandler, approvalHandler, depositHandle, transferHandle, withdrawHandle)
	return r, nil
}
