# Go-Backend-Template


适用于个人 / 中小型后端 Web 项目脚手架。
# Project Name

后端 Web 服务脚手架，面向长期演进的业务项目。

---

## 技术栈

- HTTP Framework: Gin / Echo（可替换）
- Language: Go
- Storage: MySQL / Redis
- Config: YAML + Env
- Observability: Log / Metrics（可选）

---

## 目录结构

```text
.
├── cmd/                # 程序入口
│   └── server/
│       └── main.go
│
├── internal/
│   ├── api/            # 接口层（HTTP / RPC）
│   │   ├── handler/    # 请求处理（薄）
│   │   ├── middleware/
│   │   └── router.go
│   │
│   ├── service/        # 业务逻辑（核心）
│   │
│   ├── repo/           # 数据访问（DB / Cache）
│   │
│   ├── model/          # 业务模型
│   │
│   └── config/         # 配置加载
│
├── configs/             # 配置文件
│
├── pkg/                 # 通用工具库
│
├── scripts/             # 脚本
│
├── test/                # 集成 / E2E 测试
│
├── go.mod
└── README.md
```

---
## 分层说明
- api
- 负责协议解析、参数校验、返回结果
- 不写业务逻辑
- service
- 业务规则与流程编排
- 业务逻辑只允许存在于这一层
- repo
- 数据访问与外部依赖
- 不包含业务判断
- model
- 业务数据结构
- 小项目中可合并 DO / DTO

---
## 本地运行
go run cmd/server/main.go

---
## 测试
go test ./...

---
## 设计原则
- 业务逻辑集中
- 协议与业务解耦
- 可平滑演进为大型项目
- 避免过度设计

---
## 扩展建议
当出现以下情况时可升级结构：
- 业务模块增多
- 多存储实现
- 需要 mock 测试
- 团队规模扩大
升级路径不破坏现有代码。