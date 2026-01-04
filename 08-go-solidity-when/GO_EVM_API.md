# Go 与 EVM 交互常用 API

以下整理的是本项目里常用的 go-ethereum 接口，按用途分类，便于快速查阅。

## 1) 客户端与网络

- `ethclient.Dial(rpcURL)`  
  连接到以太坊节点（HTTP/WS）。

- `client.NetworkID(ctx)`  
  返回网络 ID（历史上常与 chainID 相同，但不是签名专用）。

- `client.ChainID(ctx)`  
  返回 EIP-155 chainID，用于交易签名（推荐用这个）。

## 2) 账户与地址

- `crypto.HexToECDSA(hex)`  
  私钥字符串（hex）→ `*ecdsa.PrivateKey`。

- `crypto.PubkeyToAddress(privateKey.PublicKey)`  
  公钥 → 地址（用于 from）。

## 3) 交易相关

- `client.PendingNonceAt(ctx, from)`  
  获取该地址下一笔交易的 nonce。

- `client.SuggestGasPrice(ctx)`  
  建议 GasPrice（legacy 交易）。

- `types.NewTx(&types.LegacyTx{...})`  
  构造 legacy 交易体。

- `types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)`  
  签名交易。

- `client.SendTransaction(ctx, signedTx)`  
  发送交易到节点。

- `client.TransactionReceipt(ctx, txHash)`  
  查询交易回执（是否上链、状态、区块号等）。

## 4) 合约绑定与调用

- `bind.NewKeyedTransactorWithChainID(privateKey, chainID)`  
  生成 `TransactOpts`（包含 From + Signer）。

- `when.NewWhen(contractAddress, client)`  
  生成合约绑定实例（abigen 生成）。

- `contractMethod(auth, args...)`  
  执行合约写操作（会发交易）。

- `contractMethod(&bind.CallOpts{Context: ctx}, args...)`  
  执行合约读操作（本地调用）。

## 5) 事件监听与回放

- `Watch*`（如 `WatchApproval`）  
  WS 实时订阅事件，直接返回结构化事件。

- `Filter*`（如 `FilterApproval`）  
  按区块范围拉历史事件（回放）。

- `Parse*`（如 `ParseApproval`）  
  把原始 `types.Log` 解析为事件结构体。

## 6) 典型流程

**发送交易**：  
`HexToECDSA -> PubkeyToAddress -> PendingNonceAt -> SuggestGasPrice -> NewTx -> SignTx -> SendTransaction`

**调用合约**：  
`NewKeyedTransactorWithChainID -> NewWhen -> ContractMethod`

**监听事件**：  
`NewWhen -> Watch*`（实时） / `Filter*`（回放）
