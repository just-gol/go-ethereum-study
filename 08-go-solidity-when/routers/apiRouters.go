package routers

import (
	"08-go-solidity-when/controller"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	group := r.Group("/api")
	{
		group.GET("/getBalance", controller.ApiController{}.GetBalance)
		group.POST("/transferFrom", controller.ApiController{}.TransferFrom)
		group.POST("/deposit", controller.ApiController{}.Deposit)
		group.POST("/approve", controller.ApiController{}.Approve)
		group.GET("/allowance", controller.ApiController{}.Allowance)
	}
}
