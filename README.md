# go-ethereum-study

这个仓库是个人学习与练手项目集合，主要围绕 go-ethereum 进行合约交互、交易构造、事件监听与回放等实践。

包含内容概览：
- 基础连接与账户管理
- 合约调用与交易发送
- 事件监听与历史回放
- 后端服务与数据入库

具体实现见各子目录（如 `08-go-solidity-when`）。

## 08-go-solidity-when

基于 Gin + go-ethereum 的合约交互与事件索引示例项目，包含：
- `When` 合约调用（deposit/withdraw/approve/transferFrom/allowance/balanceOf）
- 交易发送与回执跟踪（演示交易生命周期）
- 事件监听与历史回放（Watch + Filter）
- 事件入库与幂等去重（txHash + logIndex）
