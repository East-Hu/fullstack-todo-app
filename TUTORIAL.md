# 全栈 Todo App 完整入门教程 (Docker + MySQL 版)

> **适用人群**: 有编程基础（算法方向），零开发经验，想要入门 Web 全栈开发。
>
> **技术栈**: React (前端) + Go / Gin / GORM / MySQL (后端) + Docker (数据库环境) + DBeaver (数据库管理)

---

## 目录

1. [项目整体架构](#1-项目整体架构)
2. [如何启动项目](#2-如何启动项目)
3. [数据库管理 (DBeaver)](#3-数据库管理-dbeaver)
4. [Docker基础与使用](#4-docker基础与使用)
5. [Go 语言基础速览](#5-go-语言基础速览)
6. [后端详解: Gin + GORM + MySQL](#6-后端详解-gin--gorm--mysql)
7. [前端详解: React + Vite + Axios](#7-前端详解-react--vite--axios)
8. [前后端如何通信](#8-前后端如何通信)
9. [常用命令速查](#9-常用命令速查)

---

## 1. 项目整体架构

```
全栈 Todo App 工作流程:

┌──────────────┐   HTTP 请求   ┌──────────────┐   SQL 操作   ┌───────────────┐
│   浏览器      │ ──────────→  │  Go 后端      │ ──────────→ │  MySQL 容器    │
│   (React)    │ ←────────── │  (Gin+GORM)  │ ←────────── │  (Docker)     │
│   :5173      │   JSON 响应   │  :8080       │   查询结果   │  :3307->3306  │
└──────────────┘              └──────────────┘              └───────────────┘
                                     ↑ (代码中连接)
                          "root:root@tcp(127.0.0.1:3307)/todo_app"
```

### 1.1 关键变化：引入 Docker 与 MySQL

- **原设计**: 使用 SQLite，数据存放在本地文件 `todos.db`。简单但功能有限。
- **新设计**: 使用 **MySQL**，它是工业界最标准的关系型数据库。
- **Docker 的作用**: 我们不需要在 Mac 上繁琐地安装 MySQL 服务，而是通过 Docker 启动一个 "容器" (Container)。这个容器就像一个轻量级的虚拟机，里面预装好了 MySQL，我们用完即扔，不污染本机环境。

---

## 2. 如何启动项目

### 前提条件

确保已安装:
- **Go** ≥ 1.25
- **Node.js** ≥ 18
- **Docker Desktop** (Mac 上通过 `brew install --cask docker` 安装并启动)
- **DBeaver** (数据库管理工具，Mac 上通过 `brew install --cask dbeaver-community` 安装)

### 第一步：启动数据库 (MySQL)

在项目根目录下运行：

```bash
cd /Users/east/AntiGravity_projects/FullStackLearning/fullstack-todo-app
docker compose up -d
```

- `docker compose` 是 Docker 的编排工具。
- `up` 表示启动。
- `-d` (detach) 表示在后台运行。
- 它会自动下载 MySQL 镜像并在端口 **3307** 启动数据库。

> ⚠️ **注意**: `fullstack-todo-app` 这个 Docker 容器就是你的数据库服务器。**千万不要关掉它**，否则后端就连不上数据库了。如果你要停止开发，可以用 `docker compose down` 来优雅关闭。

### 第二步：启动后端

```bash
cd backend
go run main.go
```
看到 `✅ MySQL connected and migrated!` 表示成功连接到 Docker 中的 MySQL。

### 第三步：启动前端

```bash
cd frontend
npm run dev
```

**访问: http://localhost:5173**，你可以在页面上点点点来添加数据。

---

## 3. 数据库管理 (DBeaver)

因为 Docker 里的 MySQL 没有图形界面，我们可以用 **DBeaver** (类似 Navicat 或 MySQL Workbench) 来连接它，查看和管理数据。

### 3.1 连接步骤

1. 打开 **DBeaver**。
2. 点击左上角的 **插头图标** (New Connection)，选择 **MySQL**，点击 Next。
3. 填写连接信息（参考 `docker-compose.yml`）：
   - **Server Host**: `127.0.0.1` (localhost)
   - **Port**: `3307` (注意不是默认的 3306，因为我们做了映射)
   - **Database**: `todo_app`
   - **Username**: `todouser`
   - **Password**: `todopass`
4. 点击 **Test Connection** (测试连接)。如果提示成功，点击 Finish。

### 3.2 如何查看/通过 SQL 添加数据

连接成功后，在左侧导航栏展开：
`todo_app` -> `Databases` -> `todo_app` -> `Tables` -> `todos`

- **查看数据**: 双击 `todos` 表，切换到 **Data** 标签页，就能看到你刚才在网页上添加的数据。
- **写 SQL**: 右键数据库 -> **SQL Editor** -> **Open SQL Script**。
  
  ```sql
  -- 比如手动插入一条数据
  INSERT INTO todos (created_at, updated_at, title, completed) 
  VALUES (NOW(), NOW(), 'Learn SQL via DBeaver', false);
  ```
  执行后刷新网页，你就能看到这条新数据了！

---

## 4. Docker 基础与使用

**Docker 是什么？**
> 想象你在搬家。以前你需要把床、桌子、锅碗瓢盆一个个搬到新房子，还得重新组装（在每台电脑上重新安装 MySQL, Redis, 配置环境，很容易出错）。
> 
> Docker 就像一个**集装箱**。你把所有家具打包进集装箱（镜像 Image）。到了新房子，直接把集装箱放下（容器 Container），打开就能住。无论搬到哪里，里面的一切都是你打包时的样子。

**核心概念:**
1. **镜像 (Image)**: 软件的安装包（只读）。例如 `mysql:8.0` 镜像包含了运行 MySQL 所需的一切文件。
2. **容器 (Container)**: 镜像运行起来的实例（进程）。就像面向对象里的 Class (镜像) 和 Object (容器)。
3. **Docker Compose**: 用一个 YAML 文件管理多个容器。

**常用命令:**

| 命令 | 说明 |
|------|------|
| `docker compose up -d` | 根据 yml 启动所有服务 (后台运行) |
| `docker compose down` | 停止并移除容器 |
| `docker compose logs -f` | 查看容器日志 (排查数据库错误用) |
| `docker ps` | 列出正在运行的容器 |
| `docker exec -it <name> bash` | 进入容器内部 (像 SSH 连进去一样) |

---

## 5. Go 语言基础速览

> 参见原教程，Go 语法部分通用。唯一的区别是导入了不同的数据库驱动。

---

## 6. 后端详解: Gin + GORM + MySQL

### 6.1 核心改动：main.go 连接 MySQL

```go
import (
    "gorm.io/driver/mysql"  // 导入 MySQL 驱动
    "gorm.io/gorm"
)

func main() {
    // DSN (Data Source Name) 数据源名称
    // 格式: 用户名:密码@tcp(地址:端口)/数据库名?参数
    // 注意端口是 3307 (因为我们在 Docker 中映射到了这个宿主机端口)
    dsn := "todouser:todopass@tcp(127.0.0.1:3307)/todo_app?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to MySQL:", err)
    }
    
    // ... 其他逻辑不变 (Gin, Routes)
}
```

**为什么 GORM 很强？**
- 你会发现，除了连接字符串 (`start options`) 和驱动导入变了，**其他的代码 (Models, Controllers, Routes) 一行都不用改！**
- GORM 屏蔽了底层数据库的差异（SQL 方言不同），让你用一套 Go 代码操作 MySQL、PostgreSQL 或 SQLite。

### 6.2 MySQL 在 Docker 中的配置

查看 `docker-compose.yml`：

```yaml
services:
  mysql:
    image: mysql:8.0             # 使用官方 MySQL 8.0 镜像
    container_name: todo-mysql
    environment:                 # 环境变量设置初始账号密码
      MYSQL_ROOT_PASSWORD: root123
      MYSQL_DATABASE: todo_app   # 自动创建这个库
      MYSQL_USER: todouser       # 创建普通用户
      MYSQL_PASSWORD: todopass
    ports:
      - "3307:3306"              # 端口映射: 把容器内的 3306 映射到本机的 3307
                                 # 这样本机虽然没装 MySQL，但访问 localhost:3307 就是访问容器里的 MySQL
    volumes:
      - mysql_data:/var/lib/mysql # 数据持久化: 即使删了容器，数据还在 docker volume 里

volumes:
  mysql_data:
```

---

## 7. 前端详解: React + Vite + Axios

> 前端代码**完全不需要修改**。
> 
> 前端只负责发 HTTP 请求给后端 API。后端把数据存在文件里 (SQLite) 还是存到 Docker 容器里 (MySQL)，前端完全感知不到，也不关心。这就是**前后端分离**架构的优势。

---

## 8. 前后端如何通信

### HTTP 请求/响应的完整流程

以"添加一个 Todo"为例:

```
1. 用户在 React 输入 "Learn Go"，点击"添加"

2. React 调用:
   axios.post('/api/todos', { title: "Learn Go" })

3. Vite 代理转发到:
   POST http://localhost:8080/api/todos
   Content-Type: application/json
   Body: {"title": "Learn Go"}

4. Gin 路由匹配到:
   api.POST("/todos", todoCtrl.CreateTodo)

5. CreateTodo 函数执行:
   a. c.ShouldBindJSON(&todo)    → 解析 JSON 到 todo 结构体
   b. tc.DB.Create(&todo)        → INSERT INTO todos ... (发给 Docker MySQL)
   c. c.JSON(201, todo)          → 返回 JSON 响应

6. React 收到响应:
   res.data = { ID: 1, title: "Learn Go", completed: false, ... }

7. React 更新 state:
   setTodos([...todos, res.data])

8. React 重新渲染 → 页面上显示新的 Todo
```

---

## 9. 常用命令速查

### 调试 MySQL

如果你不想用 DBeaver，想装酷用命令行看看数据库里到底存了什么：

```bash
# 进入 MySQL 容器内部
docker exec -it todo-mysql mysql -u todouser -ptodopass todo_app

# 在 MySQL 命令行中:
mysql> show tables;
mysql> select * from todos;
mysql> exit
```

### 项目开发流程

1. **启动环境**: `docker compose up -d`
2. **开发后端**: 修改 Go 代码 → `Ctrl+C` 停掉后端 → `go run main.go` 重启
3. **开发前端**: 修改 React 代码 → 浏览器自动刷新
4. **结束工作**: `docker compose down` (可选，停止数据库释放资源)
