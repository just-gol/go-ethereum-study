package handle

import (
	"08-go-solidity-when/service"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ListenerHandler struct {
	svc service.ListenerService
}

func NewListenerHandler(svc service.ListenerService) *ListenerHandler {
	return &ListenerHandler{svc: svc}
}

func (l ListenerHandler) MonitorEvent(ctx *gin.Context) {
	address := ctx.PostForm("contractAddress")
	event, err := l.svc.MonitorEvent(ctx.Request.Context(), common.HexToAddress(address))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
