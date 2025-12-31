package handle

import (
	"08-go-solidity-when/models"
	"08-go-solidity-when/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepositHandle struct {
	svc service.DepositService
}

func NewDepositHandle(svc service.DepositService) *DepositHandle {
	return &DepositHandle{svc: svc}
}

func (d *DepositHandle) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	deposit, err := d.svc.GetOne(id)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, "成功", deposit)
}

func (d *DepositHandle) GetPage(ctx *gin.Context) {
	var deposit models.Deposit
	err := ctx.ShouldBindQuery(&deposit)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	page, count, pageSize, pageNum, err := d.svc.GetPage(deposit)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.PageSuccess(ctx, "成功", page, pageSize, pageNum, count)
}
