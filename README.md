# API Gateway 编程面试

## 背景

你将在 **40 分钟** 内，借助 AI 工具设计并实现一个轻量级 API Gateway 的核心功能。

你可以使用任何 AI 编码工具（Claude Code / Cursor / Copilot / 其他均可），技术栈不限（Go / Node.js / Python / Rust / Java 等均可）。

---

## 环境说明

本项目已提供 **3 个后端微服务**，通过 Docker Compose 一键启动：

```bash
docker compose up --build
```

启动后可用的后端服务：

| 服务 | 地址 | 说明 |
|------|------|------|
| user-service | `http://localhost:8081` | 用户服务，返回用户列表 |
| order-service | `http://localhost:8082` | 订单服务，返回订单列表 |
| product-service | `http://localhost:8083` | 商品服务，返回商品列表 |

你可以用 curl 验证服务是否正常：

```bash
curl http://localhost:8081/
curl http://localhost:8082/
curl http://localhost:8083/
```

每个服务还提供 `/health` 端点用于健康检查。

---

## 你的任务

在 `gateway/` 目录下实现一个 API Gateway，对外监听一个统一端口（建议 `8080`），将请求按路由规则转发到上述后端服务。

### 必做功能（MVP）

1. **路由转发** — 根据请求路径将流量转发到对应的后端服务
   - `GET /api/users/**` → `http://localhost:8081`
   - `GET /api/orders/**` → `http://localhost:8082`
   - `GET /api/products/**` → `http://localhost:8083`

2. **中间件 / 插件机制** — 设计一个可扩展的中间件管道，让新功能可以通过"添加中间件"的方式接入，而不是修改核心转发逻辑

3. **至少实现一个中间件**（从以下任选）：
   - 请求日志（记录方法、路径、耗时、状态码）
   - 简单鉴权（如 header 中的 API Key 校验）
   - 限流（如固定窗口/令牌桶）

### 选做功能（加分项）

- 路由配置外部化（从文件/环境变量加载，参考项目根目录的 `routes.yml`）
- 熔断器（后端不可用时快速失败）
- 请求/响应改写（添加 header、修改 body 等）
- 更多中间件
- 负载均衡

---

## 交付要求

- 代码可运行
- 能用 curl 演示核心流程，例如：
  ```bash
  # 通过 gateway 访问用户服务
  curl http://localhost:8080/api/users/

  # 通过 gateway 访问订单服务
  curl http://localhost:8080/api/orders/

  # 验证中间件生效（如鉴权被拒绝）
  curl -i http://localhost:8080/api/users/  # 无 API Key → 401
  curl -H "X-API-Key: secret" http://localhost:8080/api/users/  # 有 Key → 200
  ```
- 能清晰解释你的架构设计决策

---

## 时间分配建议

| 阶段 | 时间 | 内容 |
|------|------|------|
| 需求理解 & 设计 | ~8 min | 阅读本文档，构思架构，与面试官对齐方案 |
| AI 协作实现 | ~22 min | 在 `gateway/` 目录下编码实现 |
| 演示 & 答辩 | ~10 min | curl 演示 + 面试官提问 |

---

## 注意事项

- **不要修改** `services/` 目录下的后端服务代码
- 所有你的代码都应该写在 `gateway/` 目录下
- 完成后请将你的代码提交为一个 PR
- 面试过程中面试官会全程观察你的屏幕，重点关注你**如何与 AI 协作**，而不仅仅是最终结果

---

## 项目结构

```
.
├── README.md              ← 你正在读的文件
├── routes.yml             ← 路由配置示例（可选使用）
├── docker-compose.yml     ← 一键启动后端服务
├── services/
│   ├── user-service/      ← 用户服务 (:8081)
│   ├── order-service/     ← 订单服务 (:8082)
│   └── product-service/   ← 商品服务 (:8083)
└── gateway/               ← 你的工作目录（当前为空）
```
