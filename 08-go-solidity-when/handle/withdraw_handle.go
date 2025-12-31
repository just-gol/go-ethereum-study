package handle

import (
	"08-go-solidity-when/models"
	"08-go-solidity-when/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WithdrawHandle struct {
	svc service.WithdrawService
}

func NewWithdrawHandle(svc service.WithdrawService) *WithdrawHandle {
	return &WithdrawHandle{svc: svc}
}

func (t *WithdrawHandle) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	deposit, err := t.svc.GetOne(id)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, "成功", deposit)
}

func (t *WithdrawHandle) GetPage(ctx *gin.Context) {
	var withdraw models.Withdraw
	err := ctx.ShouldBindQuery(&withdraw)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	page, count, pageSize, pageNum, err := t.svc.GetPage(withdraw)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.PageSuccess(ctx, "成功", page, pageSize, pageNum, count)
}
