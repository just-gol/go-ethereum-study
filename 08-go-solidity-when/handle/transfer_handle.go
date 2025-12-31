package handle

import (
	"08-go-solidity-when/models"
	"08-go-solidity-when/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransferHandle struct {
	svc service.TransferService
}

func NewTransferHandle(svc service.TransferService) *TransferHandle {
	return &TransferHandle{svc: svc}
}

func (t *TransferHandle) GetOne(ctx *gin.Context) {
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

func (t *TransferHandle) GetPage(ctx *gin.Context) {
	var transfer models.Transfer
	err := ctx.ShouldBindQuery(&transfer)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	page, count, pageSize, pageNum, err := t.svc.GetPage(transfer)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.PageSuccess(ctx, "成功", page, pageSize, pageNum, count)
}
