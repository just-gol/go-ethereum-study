# 08-go-solidity-when

一个基于 Gin + go-ethereum 的示例项目，用来演示与 `When` 合约交互（`deposit/approve/transferFrom/allowance/balanceOf`）。

## 架构说明

- `main.go`
  - 启动 HTTP 服务
  - 读取配置，创建以太坊客户端、Service 和 Handler
- `routers/`
  - 路由注册与分发
- `handle/`
  - 负责解析请求参数、校验、返回 JSON
  - 不直接写链交互逻辑
- `service/`
  - 负责合约调用与交易构造
  - 封装 go-ethereum 的细节，便于测试和复用
- `contract/`
  - Solidity 合约源码
- `gen/`
  - abigen 生成的 Go 绑定代码

## Handler 的作用

Handler（`handle/when_handler.go`）只做“HTTP 相关”的事情：
1. 解析参数
2. 调用 Service
3. 返回 JSON

这样可以让业务逻辑集中在 Service，Handler 更轻量、易测试。

## Service 的作用

Service（`service/when_service.go`）负责“链交互逻辑”：
1. 创建合约实例
2. 构造交易（nonce/gas/签名）
3. 发交易或读链上状态

这样 Handler 不需要关心 go-ethereum 的细节。

## 请求流程

```
HTTP 请求
  -> Router
    -> Handler（参数解析）
      -> Service（合约交互）
        -> Ethereum 节点
```

## 区块回放与定时同步

项目在启动后会自动补齐历史区块事件，并定时回放补漏，流程如下：

1) 读取上次同步区块（`sync_state` 表）。  
2) 没有记录时，用配置里的 `start_block` 作为起点。  
3) 查询链上最新区块 `latest`，减去 `confirmations` 后作为回放上限。  
4) 使用 `Filter*` 在 `[last+1, latest]` 区间回放事件并入库。  
5) 实时监听 `Watch*`，持续入库新事件。  
6) 定时任务按 `replay_interval_seconds` 继续回放补漏。

幂等策略：用 `event_log` 表的 `(tx_hash, log_index)` 做唯一约束，避免重复入库。

## 配置文件

默认读取 `08-go-solidity-when/config.json`，也可以用环境变量覆盖：

- `WHEN_CONFIG_PATH`：配置文件路径
- `WHEN_RPC_URL`：HTTP RPC 地址
- `WHEN_WS_URL`：WebSocket RPC 地址
- `WHEN_CONTRACT_ADDRESS`：合约地址

## 测试运行方式

注意：测试会触发 `models` 包初始化，需要配置 `APP_CONFIG_PATH` 指向 `config/app.ini`。
可在命令行输入,这样当前 shell 里就不用每次都写了
```bash
export APP_CONFIG_PATH=/home/lsy/go-ethereum-study/08-go-solidity-when/config/app.ini
```
方式一：进入模块目录运行

```bash
cd /home/lsy/go-ethereum-study/08-go-solidity-when
APP_CONFIG_PATH=/home/lsy/go-ethereum-study/08-go-solidity-when/config/app.ini \
go test ./handle -run TestWhenHandlerGetBalance
```

方式二：使用 `-C` 指定工作目录（Go 1.20+）

```bash
APP_CONFIG_PATH=/home/lsy/go-ethereum-study/08-go-solidity-when/config/app.ini \
go test -C /home/lsy/go-ethereum-study/08-go-solidity-when ./handle -run TestWhenHandlerGetBalance
```

方式三：运行所有测试

```bash
cd /home/lsy/go-ethereum-study/08-go-solidity-when
APP_CONFIG_PATH=/home/lsy/go-ethereum-study/08-go-solidity-when/config/app.ini \
go test ./...
```
