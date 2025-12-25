package handle

import (
	"crypto/ecdsa"
	"math/big"
	"net/http"

	"08-go-solidity-when/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type WhenHandler struct {
	svc service.WhenService
}

func NewWhenHandler(svc service.WhenService) *WhenHandler {
	return &WhenHandler{svc: svc}
}

func (h *WhenHandler) GetBalance(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.Query("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	owner := crypto.PubkeyToAddress(privateKey.PublicKey)
	balanceOf, err := h.svc.GetBalance(ctx.Request.Context(), contractAddress, owner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"balanceOf": balanceOf,
	})
}

func (h *WhenHandler) Allowance(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.Query("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	to := common.HexToAddress(ctx.Query("to"))
	allowance, err := h.svc.Allowance(ctx.Request.Context(), contractAddress, from, to)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"allowance": allowance,
	})
}

func (h *WhenHandler) TransferFrom(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	dst := common.HexToAddress(ctx.PostForm("dst"))
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	tx, err := h.svc.TransferFrom(ctx.Request.Context(), contractAddress, privateKey, dst, amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"hash": tx.Hash().Hex(),
	})
}

func (h *WhenHandler) Deposit(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	amount := big.NewInt(1000000000000000000)
	tx, err := h.svc.Deposit(ctx.Request.Context(), contractAddress, privateKey, amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"hash": tx.Hash().Hex(),
	})
}

func (h *WhenHandler) Approve(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	to := ctx.PostForm("guy")
	if to == "" {
		to = ctx.PostForm("to")
	}
	spender := common.HexToAddress(to)
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	tx, err := h.svc.Approve(ctx.Request.Context(), contractAddress, privateKey, spender, amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"hash": tx.Hash().Hex(),
	})
}

func parsePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	if len(hexKey) >= 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}
	return crypto.HexToECDSA(hexKey)
}
