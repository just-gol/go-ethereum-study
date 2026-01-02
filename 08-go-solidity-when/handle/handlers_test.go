package handle

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"08-go-solidity-when/models"
	"08-go-solidity-when/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type mockWhenService struct {
	getBalanceFn   func(context.Context, common.Address, common.Address) (*big.Int, error)
	allowanceFn    func(context.Context, common.Address, common.Address, common.Address) (*big.Int, error)
	transferFromFn func(context.Context, common.Address, *ecdsa.PrivateKey, common.Address, common.Address, *big.Int) (*types.Transaction, error)
	depositFn      func(context.Context, common.Address, *ecdsa.PrivateKey, *big.Int) (*types.Transaction, error)
	approveFn      func(context.Context, common.Address, *ecdsa.PrivateKey, common.Address, *big.Int) (*types.Transaction, error)
	withdrawFn     func(context.Context, common.Address, *ecdsa.PrivateKey, *big.Int) (*types.Transaction, error)
}

func (m mockWhenService) GetBalance(ctx context.Context, contractAddress, owner common.Address) (*big.Int, error) {
	if m.getBalanceFn == nil {
		return nil, errors.New("unexpected GetBalance")
	}
	return m.getBalanceFn(ctx, contractAddress, owner)
}

func (m mockWhenService) Allowance(ctx context.Context, contractAddress, owner, spender common.Address) (*big.Int, error) {
	if m.allowanceFn == nil {
		return nil, errors.New("unexpected Allowance")
	}
	return m.allowanceFn(ctx, contractAddress, owner, spender)
}

func (m mockWhenService) TransferFrom(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	if m.transferFromFn == nil {
		return nil, errors.New("unexpected TransferFrom")
	}
	return m.transferFromFn(ctx, contractAddress, privateKey, src, dst, amount)
}

func (m mockWhenService) Deposit(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	if m.depositFn == nil {
		return nil, errors.New("unexpected Deposit")
	}
	return m.depositFn(ctx, contractAddress, privateKey, amount)
}

func (m mockWhenService) Approve(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	if m.approveFn == nil {
		return nil, errors.New("unexpected Approve")
	}
	return m.approveFn(ctx, contractAddress, privateKey, spender, amount)
}

func (m mockWhenService) Withdraw(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	if m.withdrawFn == nil {
		return nil, errors.New("unexpected Withdraw")
	}
	return m.withdrawFn(ctx, contractAddress, privateKey, amount)
}

type mockTxService struct {
	sendFn func(context.Context, *ecdsa.PrivateKey, common.Address, *big.Int, uint64) (*types.Transaction, *types.Receipt, error)
}

func (m mockTxService) SendValueAndTrack(ctx context.Context, privateKey *ecdsa.PrivateKey, to common.Address, value *big.Int, confirmations uint64) (*types.Transaction, *types.Receipt, error) {
	if m.sendFn == nil {
		return nil, nil, errors.New("unexpected SendValueAndTrack")
	}
	return m.sendFn(ctx, privateKey, to, value, confirmations)
}

type mockApprovalService struct {
	getOneFn  func(int) (models.Approval, error)
	getPageFn func(models.Approval) ([]models.Approval, int64, int, int, error)
}

func (m mockApprovalService) GetOne(id int) (models.Approval, error) {
	if m.getOneFn == nil {
		return models.Approval{}, errors.New("unexpected GetOne")
	}
	return m.getOneFn(id)
}

func (m mockApprovalService) GetPage(approval models.Approval) ([]models.Approval, int64, int, int, error) {
	if m.getPageFn == nil {
		return nil, 0, 0, 0, errors.New("unexpected GetPage")
	}
	return m.getPageFn(approval)
}

type mockDepositService struct {
	getOneFn  func(int) (models.Deposit, error)
	getPageFn func(models.Deposit) ([]models.Deposit, int64, int, int, error)
}

func (m mockDepositService) GetOne(id int) (models.Deposit, error) {
	if m.getOneFn == nil {
		return models.Deposit{}, errors.New("unexpected GetOne")
	}
	return m.getOneFn(id)
}

func (m mockDepositService) GetPage(deposit models.Deposit) ([]models.Deposit, int64, int, int, error) {
	if m.getPageFn == nil {
		return nil, 0, 0, 0, errors.New("unexpected GetPage")
	}
	return m.getPageFn(deposit)
}

type mockTransferService struct {
	getOneFn  func(int) (models.Transfer, error)
	getPageFn func(models.Transfer) ([]models.Transfer, int64, int, int, error)
}

func (m mockTransferService) GetOne(id int) (models.Transfer, error) {
	if m.getOneFn == nil {
		return models.Transfer{}, errors.New("unexpected GetOne")
	}
	return m.getOneFn(id)
}

func (m mockTransferService) GetPage(transfer models.Transfer) ([]models.Transfer, int64, int, int, error) {
	if m.getPageFn == nil {
		return nil, 0, 0, 0, errors.New("unexpected GetPage")
	}
	return m.getPageFn(transfer)
}

type mockWithdrawService struct {
	getOneFn  func(int) (models.Withdraw, error)
	getPageFn func(models.Withdraw) ([]models.Withdraw, int64, int, int, error)
}

func (m mockWithdrawService) GetOne(id int) (models.Withdraw, error) {
	if m.getOneFn == nil {
		return models.Withdraw{}, errors.New("unexpected GetOne")
	}
	return m.getOneFn(id)
}

func (m mockWithdrawService) GetPage(withdraw models.Withdraw) ([]models.Withdraw, int64, int, int, error) {
	if m.getPageFn == nil {
		return nil, 0, 0, 0, errors.New("unexpected GetPage")
	}
	return m.getPageFn(withdraw)
}

func TestWhenHandlerGetBalance(t *testing.T) {
	gin.SetMode(gin.TestMode)
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	privHex := "0x" + hex.EncodeToString(crypto.FromECDSA(privateKey))
	contract := common.HexToAddress("0x0000000000000000000000000000000000000001")

	mock := mockWhenService{
		getBalanceFn: func(ctx context.Context, contractAddress, owner common.Address) (*big.Int, error) {
			if contractAddress != contract {
				t.Fatalf("unexpected contract: %s", contractAddress.Hex())
			}
			return big.NewInt(7), nil
		},
	}
	handler := NewWhenHandler(mock, "")

	r := gin.New()
	r.GET("/getBalance", handler.GetBalance)

	req := httptest.NewRequest("GET", "/getBalance?privateKeyString="+privHex+"&contractAddress="+contract.Hex(), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "balanceOf") {
		t.Fatalf("expected balanceOf in response, got %s", w.Body.String())
	}
}

func TestWhenHandlerApproveUsesToParam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	privHex := "0x" + hex.EncodeToString(crypto.FromECDSA(privateKey))
	contract := common.HexToAddress("0x0000000000000000000000000000000000000002")
	spender := common.HexToAddress("0x0000000000000000000000000000000000000003")

	mock := mockWhenService{
		approveFn: func(ctx context.Context, contractAddress common.Address, key *ecdsa.PrivateKey, gotSpender common.Address, amount *big.Int) (*types.Transaction, error) {
			if gotSpender != spender {
				t.Fatalf("unexpected spender: %s", gotSpender.Hex())
			}
			return types.NewTx(&types.LegacyTx{Nonce: 1, To: &contract}), nil
		},
	}
	handler := NewWhenHandler(mock, "")

	r := gin.New()
	r.POST("/approve", handler.Approve)

	form := url.Values{}
	form.Set("privateKeyString", privHex)
	form.Set("contractAddress", contract.Hex())
	form.Set("to", spender.Hex())

	req := httptest.NewRequest("POST", "/approve", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "hash") {
		t.Fatalf("expected hash in response, got %s", w.Body.String())
	}
}

func TestTxHandlerMissingTo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := mockTxService{
		sendFn: func(ctx context.Context, privateKey *ecdsa.PrivateKey, to common.Address, value *big.Int, confirmations uint64) (*types.Transaction, *types.Receipt, error) {
			return nil, nil, errors.New("should not be called")
		},
	}
	handler := NewTxHandler(mock)

	r := gin.New()
	r.POST("/sendValue", handler.SendValueAndTrack)

	form := url.Values{}
	form.Set("privateKeyString", "0x00")
	form.Set("valueWei", "1")

	req := httptest.NewRequest("POST", "/sendValue", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestApprovalHandleGetOne(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := mockApprovalService{
		getOneFn: func(id int) (models.Approval, error) {
			if id != 1 {
				t.Fatalf("unexpected id: %d", id)
			}
			return models.Approval{Id: 1, Src: "0x1"}, nil
		},
	}
	handler := NewApprovalHandle(mock)

	r := gin.New()
	r.GET("/approval", handler.GetOne)

	req := httptest.NewRequest("GET", "/approval?id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "\"code\":200") {
		t.Fatalf("expected code 200 in response, got %s", w.Body.String())
	}
}

func TestDepositHandleGetOne(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := mockDepositService{
		getOneFn: func(id int) (models.Deposit, error) {
			if id != 1 {
				t.Fatalf("unexpected id: %d", id)
			}
			return models.Deposit{Id: 1, Dst: "0x1"}, nil
		},
	}
	handler := NewDepositHandle(mock)

	r := gin.New()
	r.GET("/deposit", handler.GetOne)

	req := httptest.NewRequest("GET", "/deposit?id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "\"code\":200") {
		t.Fatalf("expected code 200 in response, got %s", w.Body.String())
	}
}

func TestTransferHandleGetOne(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := mockTransferService{
		getOneFn: func(id int) (models.Transfer, error) {
			if id != 1 {
				t.Fatalf("unexpected id: %d", id)
			}
			return models.Transfer{Id: 1, Src: "0x1"}, nil
		},
	}
	handler := NewTransferHandle(mock)

	r := gin.New()
	r.GET("/transfer", handler.GetOne)

	req := httptest.NewRequest("GET", "/transfer?id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "\"code\":200") {
		t.Fatalf("expected code 200 in response, got %s", w.Body.String())
	}
}

func TestWithdrawHandleGetOne(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := mockWithdrawService{
		getOneFn: func(id int) (models.Withdraw, error) {
			if id != 1 {
				t.Fatalf("unexpected id: %d", id)
			}
			return models.Withdraw{Id: 1, Src: "0x1"}, nil
		},
	}
	handler := NewWithdrawHandle(mock)

	r := gin.New()
	r.GET("/withdraw", handler.GetOne)

	req := httptest.NewRequest("GET", "/withdraw?id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "\"code\":200") {
		t.Fatalf("expected code 200 in response, got %s", w.Body.String())
	}
}

var _ service.WhenService = mockWhenService{}
var _ service.TxService = mockTxService{}
var _ service.ApprovalService = mockApprovalService{}
var _ service.DepositService = mockDepositService{}
var _ service.TransferService = mockTransferService{}
var _ service.WithdrawService = mockWithdrawService{}
