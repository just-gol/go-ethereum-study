package handle

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"net/http"
	"strconv"

	"08-go-solidity-when/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type WhenHandler struct {
	svc             service.WhenService
	defaultContract common.Address
	hasDefault      bool
}

func NewWhenHandler(svc service.WhenService, defaultContract string) *WhenHandler {
	handler := &WhenHandler{svc: svc}
	if defaultContract != "" {
		handler.defaultContract = common.HexToAddress(defaultContract)
		handler.hasDefault = true
	}
	return handler
}

func (h *WhenHandler) GetBalance(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.Query("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractAddress, err := h.resolveContractAddress(ctx.Query("contractAddress"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	contractAddress, err := h.resolveContractAddress(ctx.Query("contractAddress"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	contractAddress, err := h.resolveContractAddress(ctx.PostForm("contractAddress"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dst := common.HexToAddress(ctx.PostForm("dst"))
	src := common.HexToAddress(ctx.PostForm("src"))
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	tx, err := h.svc.TransferFrom(ctx.Request.Context(), contractAddress, privateKey, src, dst, amount)
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
	contractAddress, err := h.resolveContractAddress(ctx.PostForm("contractAddress"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	contractAddress, err := h.resolveContractAddress(ctx.PostForm("contractAddress"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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

func (h *WhenHandler) Withdraw(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyString"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractAddress, err := h.resolveContractAddress(ctx.PostForm("contractAddress"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	amount := ctx.PostForm("amount")
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	tx, err := h.svc.Withdraw(ctx.Request.Context(), contractAddress, privateKey, new(big.Int).Mul(big.NewInt(int64(amountInt)), big.NewInt(1e18)))
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

func (h *WhenHandler) resolveContractAddress(address string) (common.Address, error) {
	if address == "" {
		if h.hasDefault {
			return h.defaultContract, nil
		}
		return common.Address{}, errors.New("contractAddress is required")
	}
	return common.HexToAddress(address), nil
}
