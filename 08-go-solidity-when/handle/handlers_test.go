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

func TestWhenHandlerGetBalance(t *testing.T) { // 定义一个单元测试：测试 WhenHandler 的 GetBalance 接口行为
	gin.SetMode(gin.TestMode) // 将 Gin 切换到测试模式，避免测试时输出多余日志等

	privateKey, err := crypto.GenerateKey() // 生成一对新的 ECDSA 私钥（以太坊账户的私钥）
	if err != nil {                         // 如果生成失败
		t.Fatalf("generate key: %v", err) // 直接让测试失败并打印错误
	}

	// 将私钥转换成 16 进制字符串，并加上 0x 前缀（模拟请求参数 privateKeyString）
	privHex := "0x" + hex.EncodeToString(crypto.FromECDSA(privateKey))

	// 构造一个合约地址（这里用的是一个固定地址，模拟请求里的 contractAddress 参数）
	contract := common.HexToAddress("0x0000000000000000000000000000000000000001")

	// 构造一个 mock service（假实现），用于替代真实链上查询逻辑，从而让测试可控、可重复
	mock := mockWhenService{
		// 注入一个函数，用来模拟 getBalance 的行为
		getBalanceFn: func(ctx context.Context, contractAddress, owner common.Address) (*big.Int, error) {
			// 断言：handler 调用 service 时传入的 contractAddress 必须等于我们期望的地址
			if contractAddress != contract {
				t.Fatalf("unexpected contract: %s", contractAddress.Hex()) // 若不一致，测试失败
			}

			// 返回一个固定余额 7（big.Int 是因为链上金额通常很大）
			return big.NewInt(7), nil
		},
	}

	handler := NewWhenHandler(mock, "") // 用 mock service 创建 handler；第二个参数是配置（这里传空字符串）

	r := gin.New() // 创建一个全新的 Gin Engine（路由器），用于测试路由

	r.GET("/getBalance", handler.GetBalance) // 注册 GET 路由：访问 /getBalance 时走 handler.GetBalance

	// 构造一条 HTTP 请求：
	// - Method: GET
	// - Path: /getBalance
	// - Query 参数：privateKeyString=... & contractAddress=...
	req := httptest.NewRequest(
		"GET",
		"/getBalance?privateKeyString="+privHex+"&contractAddress="+contract.Hex(),
		nil, // GET 请求一般没有 body，这里传 nil
	)

	w := httptest.NewRecorder() // 创建一个响应记录器，用于捕获 Gin 返回的响应（状态码/响应体等）

	r.ServeHTTP(w, req) // 让 Gin 路由器处理这次请求，并把结果写入 w

	if w.Code != 200 { // 断言：HTTP 状态码应该是 200 OK
		t.Fatalf("expected 200, got %d", w.Code) // 如果不是 200，测试失败
	}

	// 断言：响应体中应包含 "balanceOf"
	// 这里通常意味着 handler 返回内容里包含对 ERC20 balanceOf 的调用信息或说明文本
	if !strings.Contains(w.Body.String(), "balanceOf") {
		t.Fatalf("expected balanceOf in response, got %s", w.Body.String()) // 不包含则测试失败
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
