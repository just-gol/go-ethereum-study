package handle

import (
	"math/big"
	"net/http"
	"strconv"

	"08-go-solidity-when/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TxHandler struct {
	svc service.TxService
}

func NewTxHandler(svc service.TxService) *TxHandler {
	return &TxHandler{svc: svc}
}

func (h *TxHandler) SendValueAndTrack(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	toStr := ctx.PostForm("to")
	if toStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "to is required"})
		return
	}
	valueStr := ctx.PostForm("valueWei")
	if valueStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "valueWei is required"})
		return
	}
	value, ok := new(big.Int).SetString(valueStr, 10)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid valueWei"})
		return
	}
	confirmations := uint64(1)
	if c := ctx.PostForm("confirmations"); c != "" {
		parsed, err := strconv.ParseUint(c, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid confirmations"})
			return
		}
		confirmations = parsed
	}

	tx, receipt, err := h.svc.SendValueAndTrack(ctx.Request.Context(), privateKey, common.HexToAddress(toStr), value, confirmations)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"hash":        tx.Hash().Hex(),
		"status":      receipt.Status,
		"blockNumber": receipt.BlockNumber.String(),
	})
}
