# 全栈 Todo 应用 (Go + React + Docker + JWT 认证)

> 一个带有**用户认证系统**的现代化全栈应用，展示了如何使用 **Go (Gin)** 后端、**React** 前端、**MySQL** 数据库和 **JWT 认证**，并通过 **Docker Compose** 进行集成。

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg?style=flat&logo=go)
![React](https://img.shields.io/badge/React-19-61DAFB.svg?style=flat&logo=react)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED.svg?style=flat&logo=docker)

[English](./README.md) | [简体中文](./README.zh-CN.md)

---

## 简介

本项目是为那些想要从算法/基础编程转型到 **全栈开发** 的开发者准备的最佳入门案例。它实现了完整的 CRUD 流程，并带有 **JWT 用户认证**，确保每个用户只能访问自己的数据。

### 核心特性
- **用户认证**: 注册/登录，JWT Token，bcrypt 密码哈希
- **数据隔离**: 每个用户只能看到自己的 Todo
- **前端**: 使用 **React 19**、**React Router** 和 **Vite** 构建
- **后端**: 基于 **Go** 和 **Gin** 框架的 RESTful API
- **数据库**: 使用 **MySQL 8.0**，通过 **GORM** (ORM 库) 操作
- **容器化**: 数据库运行在 **Docker** 中，保持本机环境整洁
- **环境变量**: 敏感配置通过 `.env` 文件管理

---

## 技术栈

| 组件 | 技术 | 说明 |
|------|------|------|
| **前端** | React 19, React Router, Vite, Axios | 现代 UI + 客户端路由 |
| **后端** | Go, Gin, GORM, JWT | 高性能 API + 认证中间件 |
| **认证** | bcrypt, golang-jwt | 密码哈希 + JSON Web Token |
| **数据库** | MySQL 8.0 | 工业级标准关系型数据库 |
| **基础设施** | Docker Compose | 开发环境容器编排 |

---

## API 接口

| 方法 | 路径 | 说明 | 需要认证 |
|------|------|------|:-------:|
| POST | /api/register | 注册新用户 | 否 |
| POST | /api/login | 登录获取 JWT Token | 否 |
| GET | /api/todos | 获取当前用户的 Todo 列表 | 是 |
| POST | /api/todos | 创建 Todo | 是 |
| PUT | /api/todos/:id | 更新 Todo | 是 |
| DELETE | /api/todos/:id | 删除 Todo | 是 |

---

## 快速开始

### 准备工作
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Go](https://go.dev/) (v1.21+)
- [Node.js](https://nodejs.org/) (v18+)

### 1. 克隆仓库
```bash
git clone https://github.com/YOUR_USERNAME/fullstack-todo-app.git
cd fullstack-todo-app
```

### 2. 启动数据库 (Docker)
```bash
docker compose up -d
```

### 3. 启动后端
```bash
cd backend
go mod tidy       # 安装 Go 依赖
go run main.go    # 启动 API 服务器 (http://localhost:8080)
```

### 4. 启动前端
打开一个**新终端**:
```bash
cd frontend
npm install       # 安装 Node 依赖
npm run dev       # 启动 UI (http://localhost:5173)
```

**访问 [http://localhost:5173](http://localhost:5173)** — 注册账号，开始管理你的 Todo！

---

## 项目结构

```
fullstack-todo-app/
├── docker-compose.yml          # MySQL 数据库配置
├── backend/
│   ├── .env                    # 环境变量 (数据库、JWT 密钥)
│   ├── main.go                 # 入口文件 (数据库连接、服务启动)
│   ├── models/
│   │   ├── user.go             # 用户模型 (用户名、密码哈希)
│   │   └── todo.go             # Todo 模型 (标题、完成状态、用户ID)
│   ├── controllers/
│   │   ├── auth.go             # 注册/登录处理器
│   │   └── todo.go             # CRUD 处理器 (按用户过滤)
│   ├── middleware/
│   │   └── auth.go             # JWT 验证中间件
│   └── routes/
│       └── routes.go           # 路由定义
└── frontend/
    ├── src/
    │   ├── main.jsx            # React 入口 (BrowserRouter)
    │   ├── App.jsx             # 路由容器 (认证状态管理)
    │   └── pages/
    │       ├── Login.jsx       # 登录/注册页面
    │       └── Todos.jsx       # Todo 管理页面
    └── vite.config.js          # 构建配置 (API 代理)
```

---

## 学习资源

本仓库包含一份详尽的教程：
- **[TUTORIAL.md](./TUTORIAL.md)**: 完整指南 (23 章)，覆盖每一行代码 —— 从 Go 基础语法、Gin 框架，到 JWT 认证原理，再到 React Router 和完整请求生命周期。

---

## 贡献代码

欢迎 Fork 本项目并提交 Pull Request。

1. Fork 本项目
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

---

## 许可证

本项目基于 MIT 许可证分发。详情请参阅 `LICENSE` 文件。
