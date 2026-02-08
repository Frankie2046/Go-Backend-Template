# Go-Backend-Template

适用于个人 / 中小型后端 Web 项目脚手架，默认基于 Fiber + Wire + Zap + Swagger。

---

## 技术栈

- HTTP Framework: Fiber
- Language: Go
- Config: YAML
- DI: Google Wire
- Observability: Zap
- API Docs: Swagger (swag)

---

## 目录结构

```text
.
├── cmd/                # 程序入口
│   └── main.go
│
├── internal/
│   ├── api/            # 接口层（HTTP / RPC）
│   │   ├── handler/    # 请求处理（薄）
│   │   ├── middleware/
│   │   ├── app.go       # Fiber app 装配
│   │   ├── route.go
│   │   └── swagger.go
│   │
│   ├── service/        # 业务逻辑（核心）
│   │
│   ├── repo/           # 数据访问（DB / Cache）
│   │
│   ├── model/          # 业务模型
│   │
│   ├── config/         # 配置加载
│   └── wire/           # 依赖注入
│
├── configs/             # 配置文件
│   └── config.yaml
│
├── pkg/                 # 通用工具库
│   └── logger/          # zap 初始化
│
├── scripts/             # 脚本
│   └── swag.sh
│
├── test/                # 集成 / E2E 测试
│
├── docs/                # swagger 输出（生成）
│
├── go.mod
└── README.md
```

---
## 架构说明
- 入口（cmd/main.go）
  - 调用 wire 注入，启动服务
- API 层（internal/api）
  - middleware：请求日志
  - handler：参数解析与响应
  - route/app：路由注册与 Fiber 装配
- Service 层（internal/service）
  - 业务规则与流程编排
- Repo 层（internal/repo）
  - 数据访问与外部依赖
- Model 层（internal/model）
  - 业务数据结构
- Config（internal/config）
  - 从 `configs/config.yaml` 读取配置
- DI（internal/wire）
  - 统一实例创建与依赖注入

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
## 使用方式

### 本地运行
1) 准备配置
```sh
cp configs/config.yaml configs/config.yaml
```
2) 启动服务
```sh
go run cmd/main.go
```

### Swagger
生成文档：
```sh
./scripts/swag.sh
```
访问：
```
http://localhost:8080/swagger/index.html
```

### Docker
```sh
docker build -t go-backend .
docker run --rm -p 8080:8080 go-backend
```

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
