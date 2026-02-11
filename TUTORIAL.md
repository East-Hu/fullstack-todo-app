# 全栈 Todo App 完整入门教程 (Docker + MySQL + JWT 认证)

> **适用人群**: 有编程基础（算法方向），零开发经验，想要入门 Web 全栈开发。
>
> **技术栈**: React / React Router (前端) + Go / Gin / GORM / JWT / MySQL (后端) + Docker (数据库环境) + DBeaver (数据库管理)

---

## 目录

**Part I: 环境与项目启动**
1. [项目整体架构](#1-项目整体架构)
2. [如何启动项目](#2-如何启动项目)
3. [数据库管理 (DBeaver)](#3-数据库管理-dbeaver)
4. [Docker 基础与使用](#4-docker基础与使用)
5. [多设备协作 (Git + GitHub)](#5-多设备协作-git--github)

**Part II: 后端从零到一**
6. [Go 语言硬核入门](#6-go-语言硬核入门)
7. [Web 框架 Gin 深度解析](#7-web-框架-gin-深度解析)
8. [ORM 框架 GORM 实战](#8-orm-框架-gorm-实战)
9. [环境变量与项目配置](#9-环境变量与项目配置)
10. [后端入口逐行拆解 (main.go)](#10-后端入口逐行拆解-maingo)
11. [数据模型设计 (models/)](#11-数据模型设计-models)
12. [用户认证系统 — 理论篇](#12-用户认证系统--理论篇)
13. [用户认证系统 — 代码篇 (controllers/auth.go)](#13-用户认证系统--代码篇-controllersauthgo)
14. [JWT 中间件实现 (middleware/auth.go)](#14-jwt-中间件实现-middlewareauthgo)
15. [Todo 业务逻辑 (controllers/todo.go)](#15-todo-业务逻辑-controllerstodogo)
16. [路由系统 (routes/routes.go)](#16-路由系统-routesroutesgo)

**Part III: 前端从零到一**
17. [React 核心思想](#17-react-核心思想)
18. [Vite 与工程化](#18-vite-与工程化)
19. [Axios 与前后端通信](#19-axios-与前后端通信)
20. [React Router 前端路由](#20-react-router-前端路由)
21. [前端代码逐行拆解](#21-前端代码逐行拆解)

**Part IV: 融会贯通**
22. [一个完整请求的生命周期](#22-一个完整请求的生命周期)
23. [从这里启程](#23-从这里启程)

---

## 1. 项目整体架构

```
全栈 Todo App 工作流程 (带用户认证):

┌──────────────┐   HTTP + JWT   ┌──────────────┐   SQL 操作   ┌───────────────┐
│   浏览器      │ ──────────→   │  Go 后端      │ ──────────→ │  MySQL 容器    │
│  (React)     │ ←──────────   │ (Gin+GORM)   │ ←────────── │  (Docker)     │
│  :5173       │  JSON 响应     │  :8080       │   查询结果   │  :3307->3306  │
└──────────────┘               └──────────────┘              └───────────────┘
                                      ↑
                            配置来自 backend/.env
```

### 1.1 关键特性

- **Docker + MySQL**: 数据库运行在 Docker 容器中，不污染本机环境
- **JWT 用户认证**: 每个用户只能看到和操作自己的 Todo
- **环境变量管理**: 敏感配置（密码、密钥）不硬编码在代码里

### 1.2 项目文件结构

```
fullstack-todo-app/
├── docker-compose.yml          # Docker 编排：MySQL 数据库配置
├── .gitignore                  # Git 忽略规则
│
├── backend/                    # ===== Go 后端 =====
│   ├── .env                    # 环境变量 (数据库密码、JWT 密钥)
│   ├── main.go                 # 入口文件：连接数据库，启动服务
│   ├── go.mod / go.sum         # Go 依赖管理
│   ├── models/
│   │   ├── user.go             # User 模型 (用户表)
│   │   └── todo.go             # Todo 模型 (待办事项表)
│   ├── controllers/
│   │   ├── auth.go             # 认证控制器 (注册/登录)
│   │   └── todo.go             # Todo 控制器 (增删改查)
│   ├── middleware/
│   │   └── auth.go             # JWT 验证中间件
│   └── routes/
│       └── routes.go           # 路由配置 (URL → 处理函数)
│
└── frontend/                   # ===== React 前端 =====
    ├── package.json            # npm 依赖管理
    ├── vite.config.js          # Vite 构建配置 (含代理)
    ├── index.html              # HTML 入口
    └── src/
        ├── main.jsx            # React 入口 (挂载 Router)
        ├── App.jsx             # 路由容器 (管理登录状态)
        ├── index.css           # 全局样式
        └── pages/
            ├── Login.jsx       # 登录/注册页面
            └── Todos.jsx       # Todo 主页面
```

---

## 2. 如何启动项目

### 前提条件

确保已安装:
- **Go** >= 1.21
- **Node.js** >= 18
- **Docker Desktop** (Mac 上通过 `brew install --cask docker` 安装并启动)
- **DBeaver** (数据库管理工具，Mac 上通过 `brew install --cask dbeaver-community` 安装)

### 第一步：启动数据库 (MySQL)

在项目根目录下运行：

```bash
cd /path/to/fullstack-todo-app
docker compose up -d
```

- `docker compose` 是 Docker 的编排工具。
- `up` 表示启动。
- `-d` (detach) 表示在后台运行。
- 它会自动下载 MySQL 镜像并在端口 **3307** 启动数据库。

> **注意**: `fullstack-todo-app` 这个 Docker 容器就是你的数据库服务器。**千万不要关掉它**，否则后端就连不上数据库了。如果你要停止开发，可以用 `docker compose down` 来优雅关闭。

### 第二步：启动后端

```bash
cd backend
go run main.go
```
看到 `MySQL connected and migrated!` 表示成功连接到 Docker 中的 MySQL。

### 第三步：启动前端

```bash
cd frontend
npm run dev
```

**访问: http://localhost:5173**，你会看到登录页面。注册一个账号后就可以创建你自己的 Todo 了！

### 2.1 如何关闭项目

当你结束学习时，建议按以下顺序关闭，以节省电脑资源：

1. **关闭前端/后端**: 在运行终端中按 `Ctrl + C` 即可停止服务。
2. **关闭数据库**:
   ```bash
   docker compose down
   ```
   > **注意**: 如果你明天还要继续写代码，**完全可以不运行这句命令**，让数据库一直在后台挂着也没问题（就像你不会每次关机都卸载微信一样）。
   > 但如果你觉得电脑变卡了，或者想彻底清理环境，再运行它。
   >
   > 即使你删除了容器，数据（比如你创建的 Todo）因为保存在 `volumes` 里，下次启动时依然都在。

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
`todo_app` -> `Databases` -> `todo_app` -> `Tables`

你会看到两张表：
- **`users`** — 存储用户信息（用户名、密码哈希）
- **`todos`** — 存储待办事项，每条 todo 通过 `user_id` 关联到某个用户

```sql
-- 查看所有用户
SELECT id, username, created_at FROM users;

-- 查看所有 Todo（包含用户归属）
SELECT t.id, t.title, t.completed, u.username
FROM todos t JOIN users u ON t.user_id = u.id;
```

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

## 5. 多设备协作 (Git + GitHub)

既然你有几台电脑想同步开发，最标准的方法就是使用 **Git** 和 **GitHub**。

我已经在你的项目里初始化了 Git 仓库，并配置好了 `.gitignore` (忽略了垃圾文件)。

### 5.1 在当前电脑上 (上传代码)

1. **去 GitHub 创建仓库**: 登录 GitHub，点右上角 `+` -> `New repository`，起名 `fullstack-todo-app`，不要勾选 Initialize with README/gitignore。
2. **关联远程库**: 在项目根目录下运行 (替换成你的地址):
   ```bash
   git remote add origin https://github.com/YOUR_USERNAME/fullstack-todo-app.git
   git branch -M main
   git push -u origin main
   ```

### 5.2 在其他电脑上 (下载代码)

1. **克隆代码**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/fullstack-todo-app.git
   cd fullstack-todo-app
   ```

2. **启动环境** (无论哪台电脑，步骤都一样):
   ```bash
   # 1. 启动数据库 (Docker 保证环境一致!)
   docker compose up -d

   # 2. 启动后端
   cd backend
   go mod tidy      # 第一次需要下载依赖
   go run main.go

   # 3. 启动前端
   cd ../frontend
   npm install      # 第一次需要下载依赖
   npm run dev
   ```

### 5.3 日常同步流程

- **在电脑 A 写完代码**:
  ```bash
  git add .
  git commit -m "完成了一些功能"
  git push
  ```

- **在电脑 B 开始工作前**:
  ```bash
  git pull
  # 如果有新依赖:
  # go mod tidy (后端)
  # npm install (前端)
  ```

---

# Part II: 后端从零到一

> 从这里开始，我们深入每一个文件的代码。建议你一边读教程，一边打开对应的源文件对照阅读。

---

## 6. Go 语言硬核入门

Go 语言是 Google 设计的，它的哲学是 **"少即是多" (Less is More)**。它没有 Java 那么多花里胡哨的特性（没有继承，没有异常捕获），但它**快**，**稳**，**简单**。

### 6.1 变量与强类型

Go 是强类型语言。一旦定义了 `age` 是 `int`，你就不能把 "18" (string) 赋值给它。

```go
// 1. 完整声明 (显式指定类型)
var name string = "Inkka"

// 2. 类型推导 (编译器自动推断类型)
var age = 18

// 3. 简短声明 (最常用！只能在函数内部使用)
email := "east@example.com"
// := 的意思是：声明变量并赋值，编译器自动推断类型
```

**在本项目中的例子** (`backend/main.go`):
```go
serverPort := os.Getenv("SERVER_PORT")  // serverPort 被推断为 string 类型
if serverPort == "" {
    serverPort = "8080"  // 默认值
}
```

### 6.2 结构体 (Struct) —— Go 的面向对象

Go 没有 `class`，只有 `struct`。把 `struct` 看作是数据的**蓝图**。

```go
type Todo struct {
    // 字段名首字母大写 (Title) = Public (公开)，其他包可以访问
    // 字段名首字母小写 (title) = Private (私有)，只有当前包能看

    ID        uint   `gorm:"primaryKey"` // 反引号里的叫做 Tag (标签)
    Title     string `json:"title"`      // 告诉 JSON 解析器：转成 JSON 时 Key 叫 "title"
    Completed bool   `json:"completed"`
}
```

**为什么需要 Tag？** 因为 Go 的字段名是 `Title` (大写开头)，但前端 JSON 惯例是 `title` (小写)。Tag 就是翻译规则。

### 6.3 方法 (Method) —— 给 struct 绑定行为

Go 没有 class 的方法，但可以给 struct 绑定函数：

```go
// (tc *TodoController) 叫做 "接收者 (Receiver)"
// 意思是：这个函数属于 TodoController 这个 struct
func (tc *TodoController) GetTodos(c *gin.Context) {
    // tc.DB 就像 Python 里的 self.db
    tc.DB.Find(&todos)
}
```

**在本项目中**: 打开 `backend/controllers/todo.go`，所有函数都以 `(tc *TodoController)` 开头，说明它们是 `TodoController` 的方法，可以通过 `tc.DB` 访问数据库连接。

### 6.4 包 (Package) 与导入 (Import)

Go 用 `package` 来组织代码。**每个文件夹就是一个包**。

```
backend/
├── main.go              → package main     (程序入口)
├── models/todo.go       → package models   (数据模型)
├── controllers/todo.go  → package controllers (业务逻辑)
├── middleware/auth.go   → package middleware  (中间件)
└── routes/routes.go     → package routes    (路由配置)
```

使用其他包时要 `import`：
```go
import (
    "fullstack-todo-app/models"      // 自己写的包，用 module名/包名
    "github.com/gin-gonic/gin"       // 第三方包，用完整路径
    "fmt"                            // Go 标准库
)
```

### 6.5 指针 (Pointer) —— `*` 和 `&` 是什么

在项目代码中你会看到很多 `*` 和 `&`，不要被吓到：

```go
// & 取地址：把变量的内存地址给出去
todo := models.Todo{Title: "Learn Go"}
db.Create(&todo)  // 传入 todo 的地址，这样 db.Create 可以修改 todo（比如填入自增 ID）

// * 解引用：通过地址找到值
func NewTodoController(db *gorm.DB) *TodoController {
    // db 参数的类型是 *gorm.DB —— "指向 gorm.DB 的指针"
    // 这样函数内和函数外操作的是同一个 db 对象，不是副本
    return &TodoController{DB: db}  // 返回 TodoController 的地址
}
```

**简单理解**: `&` 是 "拿到地址"，`*类型` 是 "这是个地址类型"。传指针的目的是避免复制、允许修改原始数据。

### 6.6 错误处理 —— Go 的特色

其他语言用 `try-catch`，Go 认为**错误是一种普通的返回值**。

```go
// 函数可以返回多个值：(正常结果, 错误)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
    // err 不是 nil，说明出错了！
    log.Fatal("连接失败:", err)  // Fatal = 打印错误 + 退出程序
}
// 走到这里说明 err == nil，连接成功，放心用 db
```

这就是为什么 Go 代码里满屏 `if err != nil` —— 它强迫你**立刻处理**每一个可能出错的地方。

### 6.7 接口 (Interface) —— 简单提一下

你不需要现在完全理解，但项目中有用到：

```go
// gin.HandlerFunc 实际上就是一个函数签名：
// type HandlerFunc func(*Context)
// 任何接收 *gin.Context 参数的函数，都可以作为 handler 使用
```

---

## 7. Web 框架 Gin 深度解析

Gin 是 Go 生态里最流行的 HTTP Web 框架。它的核心是 **Context (上下文)**。

### 7.1 什么是 `*gin.Context`?

想象你在餐厅吃饭。
- **Request**: 你点菜。
- **Response**: 服务员上菜。
- **Context**: 就是那个**服务员**。
  - 她拿着你的点菜单 (Request Data)。
  - 她负责把菜端给你 (Response Writer)。
  - 她还知道你是几号桌 (Metadata)。

在 Gin 的每个处理函数里，`c *gin.Context` 都是主角：

```go
func CreateTodo(c *gin.Context) {
    // 1. 从 Request Body 里读取 JSON 数据
    var todo models.Todo
    c.ShouldBindJSON(&todo)  // 把 JSON → Go struct

    // 2. 从 URL 参数中取值
    id := c.Param("id")     // /api/todos/:id 里的 :id

    // 3. 从中间件存的值中取数据
    userID := c.MustGet("userID")  // 我们的 JWT 中间件会存入 userID

    // 4. 给客户端返回 JSON 响应
    c.JSON(200, todo)        // 状态码 200 + JSON 数据
    c.JSON(401, gin.H{"error": "Unauthorized"})  // gin.H 是 map 的简写
}
```

### 7.2 路由 (Router) —— URL 的指路牌

路由的作用：把 URL + HTTP 方法 → 映射到处理函数。

```go
r := gin.Default()  // 创建路由引擎（自带日志和错误恢复中间件）

// 基础路由
r.GET("/ping", handlePing)       // GET  /ping  → handlePing()
r.POST("/api/todos", createTodo) // POST /api/todos → createTodo()

// 路由分组 (Group) —— 共享前缀和中间件
api := r.Group("/api")           // /api 前缀
{
    api.POST("/login", login)    // POST /api/login
    api.GET("/todos", getTodos)  // GET  /api/todos
}
```

### 7.3 中间件 (Middleware) —— 请求的"安检门"

中间件是在请求到达处理函数之前（或之后）执行的函数。就像机场安检：

```
请求进来 → [CORS中间件] → [JWT中间件] → [你的处理函数] → 响应出去
```

```go
// 全局中间件：所有请求都会经过
r.Use(cors.New(...))

// 路由组中间件：只有这个组的请求会经过
authorized := api.Group("/")
authorized.Use(middleware.JWTAuth())  // 只有 /api/todos 等需要验证
{
    authorized.GET("/todos", ...)
}
```

中间件里的 `c.Next()` 和 `c.Abort()`：
- `c.Next()` → "安检通过，放行"
- `c.Abort()` → "安检不通过，拦截"

---

## 8. ORM 框架 GORM 实战

写原生 SQL (`SELECT * FROM ...`) 很容易拼错字符串，而且有 SQL 注入风险。ORM (Object Relational Mapping) 让我们用操作 **Go 对象** 的方式来操作 **数据库**。

### 8.1 模型与表的映射关系

```
Go struct          ←→    数据库表
─────────────────────────────────
type User struct   ←→    users 表
Username string    ←→    username 列
type Todo struct   ←→    todos 表
Title string       ←→    title 列
UserID uint        ←→    user_id 列 (外键)
```

GORM 会自动把 `CamelCase` 字段名转为 `snake_case` 列名。

### 8.2 gorm.Model —— 内置的基础字段

```go
// GORM 提供的 gorm.Model 长这样：
type Model struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time      // 自动记录创建时间
    UpdatedAt time.Time      // 自动记录更新时间
    DeletedAt gorm.DeletedAt // 软删除（不会真的删，只是标记一下）
}

// 我们的 struct 嵌入了它：
type Todo struct {
    gorm.Model               // 继承 ID, CreatedAt, UpdatedAt, DeletedAt
    Title     string         // 额外字段
    Completed bool
    UserID    uint           // 外键
}
```

### 8.3 自动迁移 (AutoMigrate)

```go
db.AutoMigrate(&models.User{}, &models.Todo{})
```

这句话会自动检查数据库：
- 表不存在？→ 创建
- 字段少了？→ 加列
- **注意**: 为安全起见，它**不会**自动删列

### 8.4 CRUD 操作对应的 GORM 方法

```go
// CREATE: 插入一条记录
db.Create(&todo)
// SQL: INSERT INTO todos (title, completed, user_id) VALUES (?, ?, ?)

// READ: 查询
db.Where("user_id = ?", userID).Find(&todos)
// SQL: SELECT * FROM todos WHERE user_id = 1 AND deleted_at IS NULL

db.First(&todo, id)
// SQL: SELECT * FROM todos WHERE id = 1 LIMIT 1

// UPDATE: 更新
db.Save(&todo)
// SQL: UPDATE todos SET title=?, completed=?, updated_at=? WHERE id=?

// DELETE: 软删除（标记 deleted_at，不是真的删）
db.Delete(&todo)
// SQL: UPDATE todos SET deleted_at = NOW() WHERE id = ?
```

**`?` 占位符** 是 GORM 防止 SQL 注入的方式。它不会直接把值拼进字符串。

---

## 9. 环境变量与项目配置

### 9.1 为什么需要环境变量？

看这行代码：
```go
// 错误示范 ❌ 数据库密码硬编码在代码里
dsn := "todouser:todopass@tcp(127.0.0.1:3307)/todo_app"
```

问题：
1. **安全性**: 代码推到 GitHub，全世界都能看到你的数据库密码
2. **灵活性**: 不同环境（开发/生产）的配置不同，每次都要改代码？
3. **协作**: 每个开发者的本地配置可能不一样

解决方案：把配置放到 **`.env` 文件** 里，`.gitignore` 把它排除在 Git 之外。

### 9.2 .env 文件长什么样

打开 `backend/.env`:

```env
DB_USER=todouser        # 数据库用户名
DB_PASS=todopass        # 数据库密码
DB_HOST=127.0.0.1       # 数据库地址
DB_PORT=3307            # 数据库端口
DB_NAME=todo_app        # 数据库名
JWT_SECRET=your-secret-key-change-in-production   # JWT 签名密钥
SERVER_PORT=8080        # 后端服务端口
```

### 9.3 godotenv —— 加载 .env 的库

```go
import "github.com/joho/godotenv"

// 在 main() 最开始调用
godotenv.Load()  // 读取当前目录下的 .env 文件，注入到环境变量中

// 之后就可以用 os.Getenv() 取值了
dbUser := os.Getenv("DB_USER")  // "todouser"
secret := os.Getenv("JWT_SECRET")  // "your-secret-key-change-in-production"
```

### 9.4 .gitignore 的作用

打开项目根目录的 `.gitignore`，你会看到：
```
backend/.env
```

这行的意思是：Git **不会**追踪 `.env` 文件。这样你把代码推到 GitHub 时，密码不会泄露。

当别人 clone 你的项目时，他们需要自己创建 `.env` 文件并填入自己的配置。

---

## 10. 后端入口逐行拆解 (main.go)

> 打开 `backend/main.go`，对照下面的注释逐行阅读。

```go
package main  // Go 程序的入口文件必须在 main 包中

import (
    "fmt"
    "log"
    "os"

    "fullstack-todo-app/models"   // 我们写的数据模型
    "fullstack-todo-app/routes"   // 我们写的路由配置

    "github.com/joho/godotenv"    // 第三方包：加载 .env 文件
    "gorm.io/driver/mysql"        // GORM 的 MySQL 驱动
    "gorm.io/gorm"                // GORM ORM 框架
)

func main() {
    // ① 加载 .env 文件
    // 如果找不到 .env（比如在生产环境用 Docker 注入环境变量），只是打印警告，不会崩溃
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }

    // ② 读取配置
    // os.Getenv() 从环境变量中取值，返回 string
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    serverPort := os.Getenv("SERVER_PORT")
    if serverPort == "" {
        serverPort = "8080"  // 如果没设置，用默认值
    }

    // ③ 拼接 DSN (Data Source Name) —— 数据库连接字符串
    // 格式: 用户名:密码@tcp(主机:端口)/数据库?参数
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPass, dbHost, dbPort, dbName)

    // ④ 建立数据库连接
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to MySQL:", err)
    }

    // ⑤ 自动建表/迁移
    // 检查 users 表和 todos 表是否存在，不存在就创建，字段有变化就更新
    db.AutoMigrate(&models.User{}, &models.Todo{})

    // ⑥ 初始化 Gin 路由并启动服务器
    r := routes.SetupRouter(db)  // db 传进去，让 controller 能操作数据库
    r.Run(":" + serverPort)      // 阻塞在这里，监听端口，等待请求
}
```

**关键洞察**: `main.go` 的职责很简单——读配置、连数据库、启动服务器。所有业务逻辑都分散在 models/controllers/routes 中。这种分层架构让代码清晰可维护。

---

## 11. 数据模型设计 (models/)

### 11.1 User 模型 (models/user.go)

> 打开 `backend/models/user.go`

```go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model                                        // ID, CreatedAt, UpdatedAt, DeletedAt
    Username string `json:"username" gorm:"uniqueIndex;size:50"`
    Password string `json:"-"`
}
```

**逐字段解析：**

| 字段 | 类型 | Tag 说明 |
|------|------|----------|
| `gorm.Model` | 嵌入结构体 | 自动提供 ID, 时间戳字段 |
| `Username` | string | `json:"username"` — JSON 返回时 key 叫 `username`<br>`gorm:"uniqueIndex;size:50"` — 数据库唯一索引，最大 50 字符 |
| `Password` | string | **`json:"-"` — 这个 `-` 非常关键！** 它的意思是：当这个 struct 被转换成 JSON 时，**跳过这个字段**。这样密码哈希永远不会返回给前端。 |

### 11.2 Todo 模型 (models/todo.go)

> 打开 `backend/models/todo.go`

```go
package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Title     string `json:"title" binding:"required"`
    Completed bool   `json:"completed"`
    UserID    uint   `json:"user_id"`
}
```

**逐字段解析：**

| 字段 | Tag 说明 |
|------|----------|
| `Title` | `binding:"required"` — Gin 绑定 JSON 时会检查：如果请求里没有 `title`，直接返回 400 错误 |
| `Completed` | 布尔值，默认 `false` (Go 零值) |
| `UserID` | **外键！** 表示这条 Todo 属于哪个用户。GORM 会在数据库中创建 `user_id` 列 |

**两张表的关系 (一对多)**:
```
users 表                    todos 表
┌────┬──────────┐           ┌────┬──────────┬─────────┐
│ id │ username │           │ id │ title    │ user_id │
├────┼──────────┤           ├────┼──────────┼─────────┤
│ 1  │ alice    │  ◄──┐     │ 1  │ 买菜     │ 1       │
│ 2  │ bob      │     ├──── │ 2  │ 写代码   │ 1       │
└────┴──────────┘     │     │ 3  │ 看电影   │ 2       │
                      └──── └────┴──────────┴─────────┘
```
Alice (user_id=1) 有两个 Todo，Bob (user_id=2) 有一个 Todo。

---

## 12. 用户认证系统 — 理论篇

在写代码之前，先理解几个核心概念。

### 12.1 为什么需要认证？

原来的 Todo App，所有人共享同一份数据。张三可以删李四的 Todo。加了认证后：
1. **身份验证**: 你是谁？（登录）
2. **权限控制**: 你能做什么？（只能操作自己的 Todo）

### 12.2 密码安全 —— bcrypt 哈希

**绝对不能明文存储密码！** 如果数据库被黑，所有用户密码都泄露了。

bcrypt 是一种**单向哈希算法**：
```
明文密码 "alice123"  →  bcrypt  →  "$2a$10$K7L1OJ45/4Y2nIvhRVpCe..."
```

- **不可逆**: 从哈希值**无法**还原出原始密码
- **加盐**: 即使两个用户密码相同，哈希值也不同（bcrypt 内部自动加随机盐）
- **验证**: 登录时，把用户输入的密码再 hash 一遍，和数据库中的 hash 比较

```go
// 注册时：明文 → 哈希
hash, _ := bcrypt.GenerateFromPassword([]byte("alice123"), bcrypt.DefaultCost)

// 登录时：比对哈希
err := bcrypt.CompareHashAndPassword(hash, []byte("alice123"))
// err == nil → 密码正确
// err != nil → 密码错误
```

### 12.3 JWT (JSON Web Token) —— 无状态的身份令牌

#### 传统方式 (Session):
```
用户登录 → 服务器创建 Session，存到内存/数据库 → 返回 Session ID (Cookie)
每次请求 → 浏览器自动带 Cookie → 服务器查找 Session → 确认身份
```
问题：服务器要存储所有用户的 Session，多台服务器还得共享。

#### JWT 方式:
```
用户登录 → 服务器生成一个 Token (包含用户信息) → 返回给前端
每次请求 → 前端在 Header 中带上 Token → 服务器验证 Token 签名 → 确认身份
```
优势：**服务器不需要存任何东西**。Token 本身就包含了所有信息。

#### JWT 长什么样？

```
eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFsaWNlIiwiZXhwIjoxNzExMDAxMjEwfQ.bqW8vHd2YayEDnVr...
```

用 `.` 分成三部分：
```
Header.Payload.Signature
```

1. **Header** (头部): `{"alg": "HS256", "typ": "JWT"}` — 使用什么算法签名
2. **Payload** (负载): `{"user_id": 1, "username": "alice", "exp": 1711001210}` — 用户信息 + 过期时间
3. **Signature** (签名): 用**服务器的密钥** (JWT_SECRET) 对前两部分签名

```
签名 = HMAC-SHA256(Header + "." + Payload, JWT_SECRET)
```

**为什么安全？** 因为没有 JWT_SECRET 就无法伪造签名。如果有人篡改了 Payload（比如把 user_id 从 1 改成 2），签名验证就会失败。

#### 完整的认证流程:

```
1. 用户注册: POST /api/register {username, password}
   → 服务器: bcrypt(password) → 存入 users 表

2. 用户登录: POST /api/login {username, password}
   → 服务器: 查用户 → bcrypt.Compare(hash, password) → 生成 JWT → 返回 token

3. 后续请求: GET /api/todos  (Header: Authorization: Bearer <token>)
   → JWT中间件: 提取 token → 验证签名 → 取出 user_id → 存入 Context
   → Controller: 从 Context 取 user_id → 只查该用户的数据
```

---

## 13. 用户认证系统 — 代码篇 (controllers/auth.go)

> 打开 `backend/controllers/auth.go`，对照下面逐行阅读。

### 13.1 结构体和构造函数

```go
type AuthController struct {
    DB *gorm.DB  // 持有数据库连接，和 TodoController 一样
}

func NewAuthController(db *gorm.DB) *AuthController {
    return &AuthController{DB: db}
}
```

### 13.2 请求输入结构

```go
type authInput struct {
    Username string `json:"username" binding:"required,min=2,max=50"`
    Password string `json:"password" binding:"required,min=6"`
}
```

**`binding` Tag 的作用**: 当 Gin 把 JSON 绑定到这个 struct 时，会自动校验：
- `required` — 不能为空
- `min=2` — 最少 2 个字符
- `max=50` — 最多 50 个字符

如果校验失败，`ShouldBindJSON()` 会返回 error，请求直接被拒绝。

### 13.3 注册逻辑 (Register)

```go
func (ac *AuthController) Register(c *gin.Context) {
    // ① 解析请求 JSON → authInput struct
    var input authInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return  // 数据格式不对，直接返回 400
    }

    // ② 密码哈希
    // bcrypt.DefaultCost = 10，意思是做 2^10 = 1024 轮运算，故意让它慢，防暴力破解
    hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // ③ 创建用户记录
    user := models.User{
        Username: input.Username,
        Password: string(hash),  // 存的是哈希，不是明文！
    }

    // ④ 插入数据库
    // 如果 username 已存在（uniqueIndex），Create 会失败
    if err := ac.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
        return  // 409 Conflict
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}
```

### 13.4 登录逻辑 (Login)

```go
func (ac *AuthController) Login(c *gin.Context) {
    // ① 解析请求
    var input authInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ② 查找用户
    var user models.User
    if err := ac.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        // 用户不存在
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // ③ 验证密码
    // user.Password 是数据库里存的 hash
    // input.Password 是用户刚输入的明文
    if err := bcrypt.CompareHashAndPassword(
        []byte(user.Password),
        []byte(input.Password),
    ); err != nil {
        // 密码不匹配
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
        // 注意：不论是用户名错还是密码错，都返回相同的消息
        // 这是安全最佳实践，防止攻击者通过不同报错信息枚举用户名
    }

    // ④ 生成 JWT Token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  user.ID,       // 放入用户 ID
        "username": user.Username, // 放入用户名
        "exp":      time.Now().Add(72 * time.Hour).Unix(),  // 72 小时后过期
    })

    // ⑤ 用密钥签名 token
    secret := os.Getenv("JWT_SECRET")
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // ⑥ 返回 token 给前端
    c.JSON(http.StatusOK, gin.H{
        "token":    tokenString,
        "username": user.Username,
    })
}
```

---

## 14. JWT 中间件实现 (middleware/auth.go)

> 打开 `backend/middleware/auth.go`

中间件就是一个"安检门"，在请求到达 controller 之前执行。

```go
package middleware

func JWTAuth() gin.HandlerFunc {
    // 返回一个函数 —— 这就是 Gin 中间件的标准写法
    return func(c *gin.Context) {
        // ① 从请求 Header 中取 Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header required"})
            c.Abort()  // 终止请求！不会继续执行后面的 handler
            return
        }

        // ② 解析 "Bearer <token>" 格式
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(401, gin.H{"error": "Invalid authorization format"})
            c.Abort()
            return
        }

        tokenString := parts[1]
        secret := os.Getenv("JWT_SECRET")

        // ③ 验证 token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // 这个回调函数检查签名算法是否是 HMAC (我们用的 HS256)
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return []byte(secret), nil  // 返回密钥用于验证签名
        })

        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // ④ 从 token 中取出用户 ID
        claims := token.Claims.(jwt.MapClaims)
        userID := uint(claims["user_id"].(float64))  // JSON 数字默认是 float64

        // ⑤ 把 userID 存入 Context，供后续的 controller 使用
        c.Set("userID", userID)

        // ⑥ 安检通过，放行
        c.Next()
    }
}
```

**数据传递流程**:
```
前端发请求 (Header: Authorization: Bearer xxx)
    ↓
JWTAuth 中间件：验证 token → c.Set("userID", 1)
    ↓
TodoController: userID := c.MustGet("userID").(uint)  // 取到 1
    ↓
DB 查询: WHERE user_id = 1
```

---

## 15. Todo 业务逻辑 (controllers/todo.go)

> 打开 `backend/controllers/todo.go`

和之前的版本相比，每个函数都多了一步：**从 Context 中取出 userID，查询时过滤用户**。

### 15.1 获取当前用户 ID 的工具函数

```go
func getUserID(c *gin.Context) uint {
    return c.MustGet("userID").(uint)
    // MustGet: 如果 key 不存在，会 panic（但因为有中间件保护，不会发生）
    // .(uint): 类型断言，把 interface{} 转成 uint
}
```

### 15.2 查询 —— 只返回当前用户的 Todo

```go
func (tc *TodoController) GetTodos(c *gin.Context) {
    userID := getUserID(c)  // 从中间件传递的 Context 中取出
    var todos []models.Todo
    tc.DB.Where("user_id = ?", userID).Find(&todos)
    // SQL: SELECT * FROM todos WHERE user_id = 1 AND deleted_at IS NULL
    c.JSON(200, todos)
}
```

### 15.3 创建 —— 自动绑定当前用户

```go
func (tc *TodoController) CreateTodo(c *gin.Context) {
    userID := getUserID(c)
    var todo models.Todo
    c.ShouldBindJSON(&todo)  // 从请求 JSON 中取 title

    todo.UserID = userID     // ← 关键！自动绑定用户 ID，前端不需要传
    tc.DB.Create(&todo)
    c.JSON(201, todo)
}
```

### 15.4 更新/删除 —— 确保只能改自己的

```go
func (tc *TodoController) UpdateTodo(c *gin.Context) {
    userID := getUserID(c)
    id := c.Param("id")

    var todo models.Todo
    // ↓ 关键安全检查：WHERE id = ? AND user_id = ?
    if err := tc.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
        c.JSON(404, gin.H{"error": "Todo not found"})
        // 如果 Bob 试图访问 Alice 的 Todo，会走到这里
        // 返回 404 而不是 403，不让攻击者知道这条数据存在
        return
    }
    // ... 后续更新逻辑
}
```

**安全设计**: 即使 Bob 知道 Alice 有个 id=6 的 Todo，他发 `PUT /api/todos/6`，因为 `user_id` 不匹配，会直接返回 404。

---

## 16. 路由系统 (routes/routes.go)

> 打开 `backend/routes/routes.go`

这个文件是整个后端的"目录"，把 URL 映射到处理函数。

```go
func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // ① CORS (跨域资源共享)
    // 浏览器安全策略：前端 (localhost:5173) 不能直接请求后端 (localhost:8080)
    // 需要后端明确允许
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},  // 允许哪些前端地址
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},  // 允许的 Header
        AllowCredentials: true,
    }))

    // ② 初始化 Controller（把 db 传进去）
    authCtrl := controllers.NewAuthController(db)
    todoCtrl := controllers.NewTodoController(db)

    // ③ API 路由组
    api := r.Group("/api")
    {
        // 公开路由 —— 不需要登录
        api.POST("/register", authCtrl.Register)  // 注册
        api.POST("/login", authCtrl.Login)         // 登录

        // 需要认证的路由 —— 必须带 JWT Token
        authorized := api.Group("/")
        authorized.Use(middleware.JWTAuth())  // ← 加上 JWT 安检门
        {
            authorized.GET("/todos", todoCtrl.GetTodos)
            authorized.POST("/todos", todoCtrl.CreateTodo)
            authorized.PUT("/todos/:id", todoCtrl.UpdateTodo)
            authorized.DELETE("/todos/:id", todoCtrl.DeleteTodo)
        }
    }

    return r
}
```

**路由总览**:

| 方法 | 路径 | 处理函数 | 需要 Token |
|------|------|----------|-----------|
| POST | /api/register | Register | 否 |
| POST | /api/login | Login | 否 |
| GET | /api/todos | GetTodos | 是 |
| POST | /api/todos | CreateTodo | 是 |
| PUT | /api/todos/:id | UpdateTodo | 是 |
| DELETE | /api/todos/:id | DeleteTodo | 是 |

---

# Part III: 前端从零到一

---

## 17. React 核心思想

React 的世界里只有两个核心概念：**组件 (Component)** 和 **状态 (State)**。

### 17.1 什么是组件?

在 React 中，组件就是一个函数，它**接收数据 (Props)**，**返回 UI (JSX)**。

```jsx
// 一个最简单的组件
function Welcome(props) {
    return <h1>Hello, {props.name}!</h1>
}

// 使用时：
<Welcome name="Alice" />
// 渲染出: <h1>Hello, Alice!</h1>
```

**在本项目中**: `App.jsx`、`Login.jsx`、`Todos.jsx` 都是组件。

### 17.2 什么是 Props?

Props 是父组件传给子组件的数据。就像函数的参数。

```jsx
// App.jsx (父组件) 传 token 给 Todos (子组件)
<Todos token={token} username={username} onLogout={handleLogout} />

// Todos.jsx (子组件) 接收
function Todos({ token, username, onLogout }) {
    // 直接使用这些 props
}
```

**解构语法**: `{ token, username, onLogout }` 等同于 `props.token`, `props.username`, `props.onLogout`，只是写法更简洁。

### 17.3 什么是状态 (`useState`)?

普通变量 `let i = 0` 在 React 里是不管用的 —— React 不知道 `i` 变了，所以不会刷新页面。

想要页面动起来，必须用 `useState`:

```jsx
// todos: 当前的数据列表
// setTodos: 修改 todos 的函数。你一旦调用它，React 就会重新渲染页面！
const [todos, setTodos] = useState([])

// 正确 ✅: 通过 setTodos 修改
setTodos([...todos, newTodo])

// 错误 ❌: 直接修改不会触发重新渲染
todos.push(newTodo)
```

### 17.4 副作用 (`useEffect`)

`useEffect` 的意思是：**"当某件事发生时，顺便做点什么"**。

最常见的用法：**"当组件第一次加载时，去后台拉取数据"**。

```jsx
useEffect(() => {
    fetchTodos()
}, [])  // 空数组 [] 表示：只在组件挂载 (Mount) 时执行一次
```

如果依赖数组里有值，比如 `[token]`，那么每次 `token` 变化时都会重新执行。

### 17.5 JSX —— HTML 和 JS 的混合体

JSX 让你在 JavaScript 里写"类似 HTML"的代码：

```jsx
return (
    <div className="app">          {/* className 而不是 class (避免和 JS 关键字冲突) */}
        <h1>{username}</h1>         {/* {} 里可以写任何 JS 表达式 */}
        {loading && <Spinner />}    {/* 条件渲染：loading 为 true 才显示 */}
        {todos.map(todo => (        /* 列表渲染：数组 → JSX 元素 */
            <div key={todo.ID}>{todo.title}</div>
        ))}
    </div>
)
```

---

## 18. Vite 与工程化

### 18.1 Vite 是什么？

你在开发时修改了代码，浏览器立马就变了，谁在干活？是 **Vite**。

- `npm run dev` 启动了一个开发服务器 (localhost:5173)
- 它利用了浏览器原生的 ES Modules 能力，实现了**毫秒级**的热更新 (HMR)
- 相比老牌的 Webpack，Vite 快了 10-100 倍

### 18.2 代理 (Proxy)

打开 `frontend/vite.config.js`，你会看到：

```js
export default defineConfig({
    plugins: [react()],
    server: {
        proxy: {
            '/api': 'http://localhost:8080'
        }
    }
})
```

这配置的意思是：前端 `fetch('/api/todos')` 时，Vite 开发服务器会自动把请求转发到 `http://localhost:8080/api/todos`。

**为什么需要代理？** 前端在 5173 端口，后端在 8080 端口，浏览器的同源策略会阻止跨域请求。通过代理，浏览器以为请求的是 5173（同源），实际被 Vite 转发到了 8080。

---

## 19. Axios 与前后端通信

### 19.1 Axios 基础

Axios 是一个 HTTP 客户端库，用于在前端发送请求给后端。

```jsx
import axios from 'axios'

// 创建一个带配置的实例
const api = axios.create({
    baseURL: '/api',  // 所有请求自动加上 /api 前缀
    headers: { Authorization: `Bearer ${token}` }  // 自动带上 JWT token
})

// 发请求
const res = await api.get('/todos')      // GET /api/todos
const res = await api.post('/todos', { title: '买菜' })  // POST /api/todos
const res = await api.put('/todos/1', { completed: true })
await api.delete('/todos/1')
```

### 19.2 async/await —— 异步操作

网络请求不是瞬间完成的，需要等待。JavaScript 用 `async/await` 来处理：

```jsx
const fetchTodos = async () => {     // async 标记这是个异步函数
    try {
        const res = await api.get('/todos')  // await 等待请求完成
        setTodos(res.data)                   // res.data 是后端返回的 JSON
    } catch (err) {
        // 请求失败（网络错误、401 未授权等）
        if (err.response?.status === 401) {
            onLogout()  // token 过期，自动退出
        }
    }
}
```

### 19.3 带 Token 的请求

登录成功后，前端收到 JWT token。之后每个请求都要在 Header 中携带它：

```
GET /api/todos
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9...
```

在 `Todos.jsx` 中，我们创建 axios 实例时就配置好了：
```jsx
const api = axios.create({
    baseURL: '/api',
    headers: { Authorization: `Bearer ${token}` },
})
```

这样所有通过 `api` 发出的请求都会自动带上 Token。

---

## 20. React Router 前端路由

### 20.1 什么是前端路由？

传统网站：点击链接 → 浏览器请求新页面 → 服务器返回新 HTML → 整个页面刷新。

SPA (单页应用)：只有一个 HTML 文件。点击链接 → JavaScript 切换显示的组件 → URL 变了，但页面没有刷新。

React Router 就是实现 SPA 路由的库。

### 20.2 路由配置

> 打开 `frontend/src/main.jsx`

```jsx
import { BrowserRouter } from 'react-router-dom'

createRoot(document.getElementById('root')).render(
    <StrictMode>
        <BrowserRouter>       {/* 整个应用包裹在 BrowserRouter 里 */}
            <App />
        </BrowserRouter>
    </StrictMode>,
)
```

> 打开 `frontend/src/App.jsx`

```jsx
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom'

function App() {
    const [token, setToken] = useState(localStorage.getItem('token'))
    const navigate = useNavigate()  // 用于程序化跳转

    return (
        <Routes>
            {/* 路由 1: /login */}
            <Route
                path="/login"
                element={
                    token ? <Navigate to="/" /> : <Login onLogin={handleLogin} />
                }
                // 如果已登录 (有 token)，自动跳转到首页
                // 否则显示登录页
            />

            {/* 路由 2: / (首页) */}
            <Route
                path="/"
                element={
                    token ? (
                        <Todos token={token} username={username} onLogout={handleLogout} />
                    ) : (
                        <Navigate to="/login" />
                    )
                }
                // 如果已登录，显示 Todo 页
                // 否则跳转到登录页
            />
        </Routes>
    )
}
```

**核心组件说明**:

| 组件 | 作用 |
|------|------|
| `<BrowserRouter>` | 路由的容器，通常包在最外层 |
| `<Routes>` | 路由的匹配器，只渲染第一个匹配的 Route |
| `<Route path="/login" element={...}>` | 定义路由规则：URL 是 /login 时显示什么 |
| `<Navigate to="/">` | 重定向到另一个路径 |
| `useNavigate()` | Hook，返回 `navigate` 函数，用于代码中跳转页面 |

### 20.3 Token 存储与 localStorage

```jsx
// 登录成功后，把 token 存到 localStorage
const handleLogin = (newToken, newUsername) => {
    localStorage.setItem('token', newToken)
    localStorage.setItem('username', newUsername)
    setToken(newToken)
    navigate('/')
}

// 退出登录，清除 localStorage
const handleLogout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('username')
    setToken(null)
    navigate('/login')
}
```

**为什么用 localStorage？** 因为它会持久存储在浏览器中。即使用户关闭浏览器再打开，token 还在，不需要重新登录（除非 token 过期了）。

---

## 21. 前端代码逐行拆解

### 21.1 登录页 (pages/Login.jsx)

> 打开 `frontend/src/pages/Login.jsx`

```jsx
function Login({ onLogin }) {
    // onLogin 是从父组件 App.jsx 传来的回调函数

    const [isRegister, setIsRegister] = useState(false)  // 当前是注册还是登录模式
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState(null)
    const [loading, setLoading] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()  // 阻止表单默认的页面跳转行为
        setError(null)
        setLoading(true)

        try {
            if (isRegister) {
                // 先注册
                await api.post('/register', { username, password })
                // 注册成功后不需要手动跳转，直接往下走 login
            }
            // 登录
            const res = await api.post('/login', { username, password })
            onLogin(res.data.token, res.data.username)  // 调用父组件传来的回调
        } catch (err) {
            setError(err.response?.data?.error || 'Connection failed')
        } finally {
            setLoading(false)  // 无论成功失败，都停止 loading
        }
    }

    return (
        // ... JSX (登录表单 UI)
        // 表单切换：登录 ↔ 注册
        <button onClick={() => setIsRegister(!isRegister)}>
            {isRegister ? '去登录' : '注册一个'}
        </button>
    )
}
```

**设计巧妙之处**: 登录和注册共用一个表单组件，通过 `isRegister` 状态切换行为。如果是注册模式，先调 `/register`，然后自动调 `/login`，一步到位。

### 21.2 Todo 页面 (pages/Todos.jsx)

> 打开 `frontend/src/pages/Todos.jsx`

和之前的 `App.jsx` 几乎一样，但有两个关键变化：

**变化 1: axios 实例带了 Token**
```jsx
const api = axios.create({
    baseURL: '/api',
    headers: { Authorization: `Bearer ${token}` },  // ← 每个请求都带 token
})
```

**变化 2: 401 自动登出**
```jsx
const fetchTodos = async () => {
    try {
        const res = await api.get('/todos')
        setTodos(res.data || [])
    } catch (err) {
        if (err.response?.status === 401) {
            onLogout()  // ← token 过期或无效，自动退出到登录页
            return
        }
        setError('无法连接到后端服务器...')
    }
}
```

**变化 3: 用户信息栏**
```jsx
<div className="user-bar">
    <span>👤 {username}</span>
    <button className="btn-link" onClick={onLogout}>退出登录</button>
</div>
```

---

# Part IV: 融会贯通

---

## 22. 一个完整请求的生命周期

让我们跟踪一个完整的流程：**用户登录后创建一条 Todo**。

### 第一幕: 登录

```
1. 用户在 Login.jsx 输入 alice / alice123，点击"登录"
2. Login.jsx: await api.post('/login', { username: 'alice', password: 'alice123' })
3. Vite 代理把请求转发到 http://localhost:8080/api/login
4. Gin 路由匹配: POST /api/login → authCtrl.Login()
5. Login(): 查数据库找到 alice → bcrypt 验证密码 → 生成 JWT
6. 返回: { token: "eyJ...", username: "alice" }
7. Login.jsx: onLogin(token, username)
8. App.jsx: 存入 localStorage，setToken(token)，navigate('/')
9. React Router 检测到 token 存在 → 渲染 <Todos>
```

### 第二幕: 创建 Todo

```
1. 用户在 Todos.jsx 输入"买菜"，点击"添加"
2. Todos.jsx: await api.post('/todos', { title: '买菜' })
   Header: Authorization: Bearer eyJ...
3. Gin 路由匹配: POST /api/todos → 先经过 JWTAuth() 中间件
4. JWTAuth(): 验证 token → 取出 user_id=1 → c.Set("userID", 1) → c.Next()
5. CreateTodo(): userID = getUserID(c) → 1
   todo.UserID = 1 → db.Create(&todo)
6. MySQL 执行: INSERT INTO todos (title, completed, user_id) VALUES ('买菜', false, 1)
7. 返回: { ID: 8, title: "买菜", completed: false, user_id: 1 }
8. Todos.jsx: setTodos([...todos, res.data]) → React 重新渲染 → 页面出现新 todo
```

### 第三幕: 数据隔离

```
Bob 登录后，请求 GET /api/todos (Bob 的 token, user_id=2)
→ JWTAuth(): c.Set("userID", 2)
→ GetTodos(): WHERE user_id = 2
→ 只返回 Bob 的 Todo，Alice 的数据完全看不到
```

---

## 23. 从这里启程

恭喜你！你已经从零理解了一个完整的全栈应用，包括：

- **Go + Gin**: HTTP 服务器、路由、中间件
- **GORM + MySQL**: 数据模型、CRUD、外键关系
- **JWT + bcrypt**: 用户认证、密码安全、Token 验证
- **React + Router**: 组件、状态、前端路由、条件渲染
- **Axios**: 前后端 HTTP 通信、Token 管理
- **Docker**: 容器化数据库环境
- **环境变量**: 敏感信息管理

**下一步可以挑战：**

| 方向 | 说明 |
|------|------|
| **Redis 缓存** | 给频繁查询的数据加缓存，学习内存数据库 |
| **消息队列** | 用 RabbitMQ 实现异步通知（如：Todo 到期提醒） |
| **文件上传** | 给 Todo 添加附件功能 |
| **部署上线** | 把应用部署到云服务器，让全世界都能访问 |
| **WebSocket** | 实现多人实时协作编辑 Todo |

保持好奇心，Stay hungry, stay foolish.
