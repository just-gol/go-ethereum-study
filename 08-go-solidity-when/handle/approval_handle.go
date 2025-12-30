package handle

import (
	"strconv"

	"08-go-solidity-when/models"
	"08-go-solidity-when/service"

	"github.com/gin-gonic/gin"
)

type ApprovalHandle struct {
	svc service.ApprovalService
}

func NewApprovalHandle(svc service.ApprovalService) *ApprovalHandle {
	return &ApprovalHandle{svc: svc}
}

func (h *ApprovalHandle) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	approval, err := h.svc.GetOne(id)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, "成功", approval)
}

func (h *ApprovalHandle) GetPage(ctx *gin.Context) {
	var approval = models.Approval{}
	err := ctx.ShouldBindQuery(approval)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	list, count, pageSize, pageNum, err := h.svc.GetPage(approval)
	if err != nil {
		models.Error(ctx, err.Error())
	}
	models.PageSuccess(ctx, "成功", list, pageNum, pageSize, count)
}
