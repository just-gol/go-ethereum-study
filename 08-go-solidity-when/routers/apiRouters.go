package routers

import (
	"08-go-solidity-when/handle"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine, handler *handle.WhenHandler, txHandler *handle.TxHandler) {
	group := r.Group("/api")
	{
		group.GET("/getBalance", handler.GetBalance)
		group.POST("/transferFrom", handler.TransferFrom)
		group.POST("/deposit", handler.Deposit)
		group.POST("/approve", handler.Approve)
		group.GET("/allowance", handler.Allowance)
		group.POST("/withdraw", handler.Withdraw)
		group.POST("/sendValue", txHandler.SendValueAndTrack)
	}
}
