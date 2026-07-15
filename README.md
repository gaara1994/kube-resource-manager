# kube-resource-manager

Kubernetes 资源配置管理平台后端服务，提供 RESTful API 来集中管理多集群的 K8s 资源。

## 技术栈

| 层次 | 技术 |
|------|------|
| Web 框架 | Gin |
| CLI | Cobra |
| ORM | GORM（PostgreSQL） |
| 认证 | JWT（golang-jwt/v5）+ bcrypt |
| 缓存 | Redis（go-redis） |
| 消息队列 | Kafka（IBM/sarama） |
| 日志 | Uber Zap |
| 监控 | Prometheus |
| 配置 | TOML |
| API 文档 | Swagger（swaggo） |

## 项目结构

```
├── main.go                     # 入口
├── config.toml                 # 配置文件
├── cmd/root.go                 # Cobra CLI 根命令
├── config/config.go            # 配置结构体 + TOML 解析
├── routes/
│   ├── router.go               # 路由注册
│   ├── ctls.go                 # 控制器实例
│   ├── healthcheck.go          # 健康检查
│   └── prometheus.go           # Prometheus 指标
├── internal/
│   ├── controllers/            # 控制器层
│   ├── dao/                    # 数据访问层
│   ├── db/
│   │   ├── migrate.go          # GORM 初始化 + AutoMigrate
│   │   └── models/             # 数据模型
│   ├── dto/                    # 请求 DTO（含校验）
│   ├── request/                # 通用请求结构（分页）
│   ├── response/               # 统一响应格式
│   ├── errcodes/               # 业务错误码
│   ├── messaging/              # Kafka 生产者/消费者
│   └── redis/                  # Redis 客户端
├── pkg/logger/                 # Zap 日志
├── utils/auth/                 # JWT + bcrypt
├── docs/                       # Swagger 文档
└── document/                   # 需求文档 + 思维导图
```

## API 路由

### 公开接口（无需认证）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/healthz` | 健康检查 |
| GET | `/metrics` | Prometheus 指标 |
| GET | `/swagger/*any` | Swagger UI |
| POST | `/login` | 用户登录，返回 JWT |

### 需 JWT 认证（`/api/v1/...`）

| 方法 | 路径 | 说明 | 状态 |
|------|------|------|------|
| GET | `/cluster/:id` | 查询集群 | ✅ |
| POST | `/cluster` | 添加集群 | ✅ |
| PUT | `/cluster` | 修改集群 | ✅ |
| DELETE | `/cluster/:id` | 删除集群 | ✅ |
| GET | `/cluster/list` | 集群列表（分页+搜索） | ✅ |
| GET | `/namespace/:id` | 查询命名空间 | ✅ |
| POST | `/namespace` | 添加命名空间 | ✅ |
| PUT | `/namespace` | 修改命名空间 | ✅ |
| DELETE | `/namespace/:id` | 删除命名空间 | ✅ |
| GET | `/namespace/list` | 命名空间列表 | ✅ |
| GET | `/resource/type/:id` | 查询资源类型 | ✅ |
| POST | `/resource/type` | 添加资源类型 | ✅ |
| PUT | `/resource/type` | 修改资源类型 | ✅ |
| DELETE | `/resource/type/:id` | 删除资源类型 | ✅ |
| GET | `/resource/type/list` | 资源类型列表 | ✅ |
| GET | `/resource/config/:id` | 查询资源配置 | ✅ |
| POST | `/resource/config` | 添加资源配置 | ✅ |
| PUT | `/resource/config` | 修改资源配置 | ✅ |
| DELETE | `/resource/config/:id` | 删除资源配置 | ✅ |
| GET | `/resource/config/list` | 资源配置列表 | ✅ |
| POST | `/user` | 创建用户 | ✅ |

## 数据模型关系

```
KubernetesCluster（集群）
    └── KubernetesNamespace（命名空间） ← cluster_id 外键

KubernetesResourceConfig（资源配置）关联三个维度：
    ├── resource_cluster_id   → 集群 ID
    ├── resource_namespace_id → 命名空间 ID
    └── resource_type_id      → 资源类型 ID

用户权限体系（RBAC 五表）：
    User ←→ UserRole ←→ Role ←→ RolePermission ←→ Permission
```

---

## 已完成功能

- [x] 项目骨架搭建（Gin + Cobra + GORM + Swagger）
- [x] 配置文件加载（TOML）
- [x] 日志初始化（Zap）
- [x] PostgreSQL 连接 + GORM AutoMigrate
- [x] 统一响应格式 `{code, message, data}`
- [x] 业务错误码体系（按模块分段）
- [x] 集群 CRUD + 分页列表查询
- [x] 命名空间 CRUD + 分页列表查询
- [x] 资源类型 CRUD + 分页列表查询
- [x] 资源配置 CRUD + 分页列表查询
- [x] 用户注册（bcrypt 密码哈希）
- [x] 用户登录（JWT 签发）
- [x] JWT 认证中间件
- [x] 健康检查 + Prometheus 指标端点
- [x] Kafka 生产者/消费者封装
- [x] Redis 客户端封装
- [x] RBAC 数据模型建表（Role、Permission、UserRole、RolePermission）

---

## 待完成功能

### 高优先级

- [ ] **接入 K8s API**：当前只是把集群信息存库，并未真正调用 K8s API。需要引入 `client-go`，实现下发 YAML 配置到真实集群
- [ ] **修复错误码冲突**：`ResourceConfigErrPost`、`UserErrPost`、`RoleErrPost`、`PermissionErrPost` 都从 40400 用 `iota` 起算，导致错误码重复，应改为不同基准值
- [ ] **JWT secret 使用配置文件**：`utils/auth/auth.go` 中 `jwtSecret` 硬编码为 `"kube-resource-manager"`，应改为读取 `config.toml` 的 `auth.jwt_secret`
- [ ] **修复 DAO 层 List 方法**：NamespaceDao、ResourceConfigDao、ResourceTypeDao、UserDao 的 List 方法使用了 `cluster_name`、`status` 等字段过滤，但这些字段在对应模型中不存在（复制粘贴残留），应根据实际模型字段修正

### 中优先级

- [ ] **Redis 接入主流程**：`internal/redis/` 已封装好，但 `main.go` 未调用 `InitRedis()`，缓存层未启用
- [ ] **Kafka 接入主流程**：`internal/messaging/` 已封装好，但未在业务中实际使用。可用于 YAML 配置变更事件的异步通知
- [ ] **角色管理 API**：Role 模型已有，但无 Controller 和 DAO
- [ ] **权限管理 API**：Permission 模型已有，但无 Controller 和 DAO
- [ ] **用户-角色关联 API**：UserRole 模型已有，但无业务逻辑
- [ ] **角色-权限关联 API**：RolePermission 模型已有，但无业务逻辑
- [ ] **用户管理完善**：目前 UserController 只有 POST（注册），缺少 GET（查询）、PUT（修改）、DELETE（删除）、列表查询
- [ ] **配置文件字段 `mode` vs `Mode`**：`config.toml` 中 `[server]` 下的 `Mode` 大小写与 Go 结构体 tag 不一致（`mode` vs `Mode`），虽然 TOML 不区分大小写，但建议统一
- [ ] **验证 `main.go` 双 `r.Run` 问题**：当前 `main.go` 里 `r.Run(":8080")` 后会阻塞，第二个 `r.Run` 永远不会执行

### 低优先级

- [ ] **Cobra CLI 完善**：`cmd/root.go` 目前只有骨架，可添加子命令（如 `serve`、`migrate`、`version`）
- [ ] **Dockerfile 修复**：当前直接 `ADD main /app/main`，缺少 `go build` 步骤，镜像构建流程不完整
- [ ] **单元测试**：目前仅有 Kafka 和 Redis 的测试文件，Controller/DAO/Service 层均无测试
- [ ] **日志级别从配置读取**：`pkg/logger/logger.go` 中日志配置是硬编码的 `NewProductionConfig()`，未使用 `config.toml` 中的 `[logging]` 配置
- [ ] **请求参数校验**：Controller 中大量 `//校验` 注释但无实际校验逻辑（如集群名唯一性、YAML 格式校验）
- [ ] **MySQL 支持**：配置中定义了 MySQL，但 `migrate.go` 仅实现了 PostgreSQL 连接
- [ ] **API 文档完善**：Swagger 注解目前仅 `KubernetesClusterController.GET` 有，其他接口缺少文档注解
- [ ] **CORS 配置**：生产环境可能需要跨域支持
- [ ] **优雅关闭**：`main.go` 缺少 signal 处理和 graceful shutdown

---

## 快速开始

```bash
# 1. 确保 PostgreSQL 可用，修改 config.toml 中的数据库连接信息

# 2. 运行
make run
# 或
go build -o main && ./main

# 3. 访问 Swagger
# http://localhost:8080/swagger/index.html
```

## 依赖安全

当前 `golang.org/x/crypto` 版本受 CVE-2026-46595 影响（仅 ssh 子包，项目实际只用了 bcrypt，无实际风险）。升级命令：

```bash
go get golang.org/x/crypto@v0.52.0
```
