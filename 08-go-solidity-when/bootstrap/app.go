package bootstrap

import (
	"context"
	"log"
	"time"

	"08-go-solidity-when/config"
	"08-go-solidity-when/handle"
	"08-go-solidity-when/models"
	"08-go-solidity-when/routers"
	"08-go-solidity-when/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewApp(cfg config.Config) (*gin.Engine, error) {
	// 1) 初始化 HTTP RPC 客户端（读链/发交易）
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
	// 2) 确保事件索引与同步进度表存在
	// AutoMigrate 是 GORM（Go 常用 ORM）提供的“自动迁移”能力：根据你传入的模型结构体，自动在数据库里创建/更新对应的表结构。
	if err := models.DB.AutoMigrate(&models.EventLog{}, &models.SyncState{}); err != nil {
		return nil, err
	}
	// 3) 初始化 WS 客户端（实时订阅事件）
	wsClient, err := ethclient.Dial(cfg.WSURL)
	if err != nil {
		return nil, err
	}
	listenerSvc := service.NewListenerService(wsClient)
	if cfg.ContractAddress != "" {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		go func() {
			// 启动回放：先补历史，再定时补漏
			if err := listenerSvc.ReplayFromLast(context.Background(), contractAddress, cfg.StartBlock, cfg.Confirmations); err != nil {
				log.Println("replay startup error:", err)
				return
			}
			listenerSvc.StartReplayLoop(context.Background(), contractAddress, cfg.StartBlock, cfg.Confirmations, time.Duration(cfg.ReplayIntervalSecond)*time.Second)
		}()
		go func() {
			// 启动实时监听
			listenerSvc.MonitorEvent(context.Background(), contractAddress)
		}()
	}

	// 4) 初始化 HTTP 路由
	r := gin.Default()
	r.Use(cors.Default())
	routers.ApiRoutersInit(r, whenHandler, transactionHandler, approvalHandler, depositHandle, transferHandle, withdrawHandle)
	return r, nil
}
