# 🚀 全栈 Todo 应用 (Go + React + Docker)

> 一个现代化的全栈应用示例，展示了如何使用 **Go (Gin)** 后端、**React** 前端和 **MySQL** 数据库，并通过 **Docker Compose** 进行集成。

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg?style=flat&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB.svg?style=flat&logo=react)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED.svg?style=flat&logo=docker)

[🇺🇸 English](./README.md) | [🇨🇳 简体中文](./README.zh-CN.md)

---

## 📖 简介

本项目是为那些想要从算法/基础编程转型到 **全栈开发** 的开发者准备的最佳入门案例。它实现了一个完整的 CRUD (增删改查) 流程，并采用了清晰的架构设计。

### ✨ 核心特性
- **前端**: 使用 **React** 和 **Vite** 构建，体验极速 UI。
- **后端**: 基于 **Go (Golang)** 和 **Gin** 框架的 RESTful API。
- **数据库**: 使用 **MySQL 8.0**，通过 **GORM** (ORM 库) 进行无缝数据操作。
- **容器化**: 整个数据库环境运行在 **Docker** 中，保持本机环境整洁。
- **热重载**: 开发过程中前后端均支持即时反馈。

---

## 🛠 技术栈

| 组件 | 技术 | 说明 |
|-----------|------------|-------------|
| **前端** | React, Vite, Axios | 现代 UI 库搭配极速构建工具 |
| **后端** | Go, Gin, GORM | 高性能编译型语言 |
| **数据库** | MySQL 8.0 | 工业级标准关系型数据库 |
| **基础设施** | Docker Compose | 开发环境容器编排 |

---

## 🚀 快速开始

### 准备工作
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Go](https://go.dev/) (v1.21+)
- [Node.js](https://nodejs.org/) (v18+)

### 1️⃣ 克隆仓库
```bash
git clone https://github.com/YOUR_USERNAME/fullstack-todo-app.git
cd fullstack-todo-app
```

### 2️⃣ 启动数据库 (Docker)
我们使用 Docker 来运行 MySQL，免去手动安装的烦恼。
```bash
docker compose up -d
```
> 这将在端口 `3307` 启动 MySQL (映射自容器的 3306)。

### 3️⃣ 启动后端
```bash
cd backend
go mod tidy       # 安装 Go 依赖
go run main.go    # 启动 API 服务器 (http://localhost:8080)
```

### 4️⃣ 启动前端
打开一个 **新终端**:
```bash
cd frontend
npm install       # 安装 Node 依赖
npm run dev       # 启动 UI (http://localhost:5173)
```

🎉 **搞定！访问 [http://localhost:5173](http://localhost:5173) 查看你的应用。**

---

## 📂 项目结构

```bash
fullstack-todo-app/
├── docker-compose.yml    # 数据库配置
├── TUTORIAL.md           # 📚 详细的保姆级教程 (重点推荐)
├── ADVANCED_ROADMAP.md   # 🚀 进阶学习路线 (Mini-Amazon)
├── backend/
│   ├── main.go           # 入口文件 (服务器 & 数据库连接)
│   ├── models/           # 数据库模型 (GORM 结构体)
│   ├── controllers/      # 业务逻辑 & 处理器
│   └── routes/           # API 路由定义
└── frontend/
    ├── src/
    │   ├── App.jsx       # 主 UI 组件
    │   ├── main.jsx      # React 入口点
    └── vite.config.js    # 构建配置
```

---

## 📚 学习资源

本仓库包含一份详尽的内部教程：
- **[TUTORIAL.md](./TUTORIAL.md)**: 完整指南，解释每一行代码、Docker 设置以及 Go 基础语法。
- **[ADVANCED_ROADMAP.md](./ADVANCED_ROADMAP.md)**: 带你通过该项目进阶到下一阶段 (认证、Redis、微服务) 的路线图。

---

## 🤝 贡献代码

欢迎 Fork 本项目并提交 Pull Request。
你的任何贡献都将 **备受感激**。

1. Fork 本项目
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

---

## 📄 许可证

本项目基于 MIT 许可证分发。详情请参阅 `LICENSE` 文件。
