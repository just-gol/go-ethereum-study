# 合约交互 API

Base path: `/api`

## When 合约交互

### GET `/api/getBalance`
- 作用：查询当前私钥地址的合约余额
- Query 参数：
  - `privateKeyString`（必填）
  - `contractAddress`（可选，默认读取配置）
- 返回：`balanceOf`

### GET `/api/allowance`
- 作用：查询当前私钥地址对 `to` 的授权额度
- Query 参数：
  - `privateKeyString`（必填）
  - `contractAddress`（可选）
  - `to`（必填，被授权地址）
- 返回：`allowance`

### POST `/api/deposit`
- 作用：向合约存入 ETH（调用 `deposit`）
- Form 参数：
  - `privateKeyString`（必填）
  - `contractAddress`（可选）
- 备注：固定存入 `1e18` wei
- 返回：`hash`

### POST `/api/withdraw`
- 作用：从合约提取 ETH（调用 `withdraw`）
- Form 参数：
  - `privateKeyString`（必填）
  - `contractAddress`（可选）
  - `amount`（必填，整数，单位 ETH，内部会乘 `1e18`）
- 返回：`hash`

### POST `/api/approve`
- 作用：授权额度（调用 `approve`）
- Form 参数：
  - `privateKeyString`（必填）
  - `contractAddress`（可选）
  - `guy` 或 `to`（必填，被授权地址）
- 备注：固定授权 `1e18` wei
- 返回：`hash`

### POST `/api/transferFrom`
- 作用：转账（调用 `transferFrom`）
- Form 参数：
  - `privateKeyString`（必填）
  - `contractAddress`（可选）
  - `src`（必填，发送方）
  - `dst`（必填，接收方）
- 备注：固定转账 `1e18` wei
- 返回：`hash`

### POST `/api/sendValue`
- 作用：直接转 ETH 并跟踪交易状态
- Form 参数：
  - `privateKeyString`（必填）
  - `to`（必填）
  - `valueWei`（必填，单位 wei）
  - `confirmations`（可选，默认 1）
- 返回：`hash`、`status`、`blockNumber`

## 事件入库查询

### GET `/api/approval/getOne`
- Query 参数：`id`

### GET `/api/approval/getPage`
- Query 参数：`id`、`src`、`guy`、`wad`、`pageNumber`、`pageSize`

### GET `/api/deposit/getOne`
- Query 参数：`id`

### GET `/api/deposit/getPage`
- Query 参数：`id`、`dst`、`wad`、`pageNumber`、`pageSize`

### GET `/api/transfer/getOne`
- Query 参数：`id`

### GET `/api/transfer/getPage`
- Query 参数：`id`、`src`、`dst`、`wad`、`pageNumber`、`pageSize`

### GET `/api/withdraw/getOne`
- Query 参数：`id`

### GET `/api/withdraw/getPage`
- Query 参数：`id`、`src`、`wad`、`pageNumber`、`pageSize`
