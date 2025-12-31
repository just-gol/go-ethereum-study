package routers

import (
	"08-go-solidity-when/handle"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine, handler *handle.WhenHandler, txHandler *handle.TxHandler,
	approvalHandler *handle.ApprovalHandle, depositHandler *handle.DepositHandle,
	transferHandler *handle.TransferHandle) {
	group := r.Group("/api")
	{
		group.GET("/getBalance", handler.GetBalance)
		group.POST("/transferFrom", handler.TransferFrom)
		group.POST("/deposit", handler.Deposit)
		group.POST("/approve", handler.Approve)
		group.GET("/allowance", handler.Allowance)
		group.POST("/withdraw", handler.Withdraw)
		group.POST("/sendValue", txHandler.SendValueAndTrack)
		group.GET("/approval/getOne", approvalHandler.GetOne)
		group.GET("/approval/getPage", approvalHandler.GetPage)
		group.GET("/deposit/getOne", depositHandler.GetOne)
		group.GET("/deposit/getPage", depositHandler.GetPage)
		group.GET("/transfer/getOne", transferHandler.GetOne)
		group.GET("/transfer/getPage", transferHandler.GetPage)
	}
}
