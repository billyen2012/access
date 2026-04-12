# API Gateway Coding Interview

## Overview

You have **40 minutes** to design and implement the core functionality of a lightweight API Gateway using AI tools.

You may use any AI coding tool (Claude Code / Cursor / Copilot / etc.) and any tech stack (Go / Node.js / Python / Rust / Java / etc.).

---

## Environment

This project provides **3 backend microservices**, ready to launch with Docker Compose:

```bash
docker compose up --build
```

Available backend services after startup:

| Service | Address | Description |
|---------|---------|-------------|
| user-service | `http://localhost:8081` | Returns a list of users |
| order-service | `http://localhost:8082` | Returns a list of orders |
| product-service | `http://localhost:8083` | Returns a list of products |

Verify the services are running:

```bash
curl http://localhost:8081/
curl http://localhost:8082/
curl http://localhost:8083/
```

Each service also exposes a `/health` endpoint for health checks.

---

## Your Task

Implement an API Gateway in the `gateway/` directory. The gateway should listen on a single port (recommended `8080`) and forward requests to the appropriate backend service based on routing rules.

### Required Features (MVP)

1. **Request Routing** — Forward traffic to the correct backend based on request path
   - `GET /api/users/**` → `http://localhost:8081`
   - `GET /api/orders/**` → `http://localhost:8082`
   - `GET /api/products/**` → `http://localhost:8083`

2. **Middleware / Plugin Pipeline** — Design an extensible middleware pipeline so that new features can be added by plugging in a middleware, rather than modifying the core routing logic

3. **At least one middleware** (choose from the following):
   - Request logging (method, path, latency, status code)
   - Simple authentication (e.g. API Key validation via header)
   - Rate limiting (e.g. fixed window / token bucket)

### Optional Features (Bonus)

- Externalized route configuration (load from file / env vars — see `routes.yml` in the project root)
- Circuit breaker (fail fast when a backend is unavailable)
- Request / response rewriting (add headers, modify body, etc.)
- Additional middlewares
- Load balancing

---

## Deliverables

- Working, runnable code
- Demonstrate the core flow with curl, for example:
  ```bash
  # Access user service through the gateway
  curl http://localhost:8080/api/users/

  # Access order service through the gateway
  curl http://localhost:8080/api/orders/

  # Verify middleware works (e.g. auth rejection)
  curl -i http://localhost:8080/api/users/          # No API Key → 401
  curl -H "X-API-Key: secret" http://localhost:8080/api/users/  # With Key → 200
  ```
- Be prepared to clearly explain your architectural decisions

---

## Suggested Time Allocation

| Phase | Time | Activity |
|-------|------|----------|
| Understand requirements & design | ~8 min | Read this document, plan your architecture, align with the interviewer |
| Implement with AI | ~22 min | Code your solution in the `gateway/` directory |
| Demo & Q&A | ~10 min | curl demo + interviewer questions |

---

## Important Notes

- **Do NOT modify** any code under the `services/` directory
- All your code should be written in the `gateway/` directory
- Submit your code as a Pull Request when finished
- The interviewer will observe your screen throughout — the focus is on **how you collaborate with AI**, not just the final result

---

## Project Structure

```
.
├── README.md              ← You are here
├── routes.yml             ← Example route config (optional)
├── docker-compose.yml     ← One-click backend startup
├── services/
│   ├── user-service/      ← User service (:8081)
│   ├── order-service/     ← Order service (:8082)
│   └── product-service/   ← Product service (:8083)
└── gateway/               ← Your working directory (currently empty)
```
