# Full-Stack Todo App (Go + React + Docker + JWT Auth)

> A modern, full-stack application with **user authentication**, built to demonstrate the integration of **Go (Gin)** backend, **React** frontend, **MySQL** database, and **JWT-based auth** using **Docker Compose**.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg?style=flat&logo=go)
![React](https://img.shields.io/badge/React-19-61DAFB.svg?style=flat&logo=react)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED.svg?style=flat&logo=docker)

[English](./README.md) | [简体中文](./README.zh-CN.md)

---

## Introduction

This project is a perfect starting point for developers transitioning from algorithms/basics to **Full-Stack Development**. It implements a complete CRUD workflow with **JWT user authentication**, ensuring each user can only access their own data.

### Features
- **User Auth**: Register/Login with JWT tokens, bcrypt password hashing
- **Data Isolation**: Each user only sees their own Todos
- **Frontend**: Built with **React 19**, **React Router**, and **Vite**
- **Backend**: RESTful API powered by **Go** and **Gin** framework
- **Database**: **MySQL 8.0** managed via **GORM** (ORM library)
- **Containerization**: Database runs in **Docker**, keeping your local machine clean
- **Environment Variables**: Sensitive config managed via `.env` files

---

## Tech Stack

| Component | Technology | Description |
|-----------|------------|-------------|
| **Frontend** | React 19, React Router, Vite, Axios | Modern UI with client-side routing |
| **Backend** | Go, Gin, GORM, JWT | High-performance API with auth middleware |
| **Auth** | bcrypt, golang-jwt | Password hashing + JSON Web Tokens |
| **Database** | MySQL 8.0 | Industrial-standard relational database |
| **Infrastructure** | Docker Compose | Container orchestration for development |

---

## API Endpoints

| Method | Path | Description | Auth Required |
|--------|------|-------------|:------------:|
| POST | /api/register | Register a new user | No |
| POST | /api/login | Login and get JWT token | No |
| GET | /api/todos | List your todos | Yes |
| POST | /api/todos | Create a todo | Yes |
| PUT | /api/todos/:id | Update a todo | Yes |
| DELETE | /api/todos/:id | Delete a todo | Yes |

---

## Getting Started

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Go](https://go.dev/) (v1.21+)
- [Node.js](https://nodejs.org/) (v18+)

### 1. Clone the Repository
```bash
git clone https://github.com/YOUR_USERNAME/fullstack-todo-app.git
cd fullstack-todo-app
```

### 2. Start the Database (Docker)
```bash
docker compose up -d
```

### 3. Start the Backend
```bash
cd backend
cp .env.example .env   # Create env file (or create .env manually)
go mod tidy             # Install Go dependencies
go run main.go          # Start API server at http://localhost:8080
```

### 4. Start the Frontend
Open a **new terminal**:
```bash
cd frontend
npm install             # Install Node dependencies
npm run dev             # Start UI at http://localhost:5173
```

**Visit [http://localhost:5173](http://localhost:5173)** — register an account and start managing your Todos!

---

## Project Structure

```
fullstack-todo-app/
├── docker-compose.yml          # MySQL database config
├── backend/
│   ├── .env                    # Environment variables (DB, JWT secret)
│   ├── main.go                 # Entry point (DB connection, server startup)
│   ├── models/
│   │   ├── user.go             # User model (username, password hash)
│   │   └── todo.go             # Todo model (title, completed, user_id)
│   ├── controllers/
│   │   ├── auth.go             # Register & Login handlers
│   │   └── todo.go             # CRUD handlers (user-scoped)
│   ├── middleware/
│   │   └── auth.go             # JWT validation middleware
│   └── routes/
│       └── routes.go           # Route definitions
└── frontend/
    ├── src/
    │   ├── main.jsx            # React entry (BrowserRouter)
    │   ├── App.jsx             # Route container (auth state)
    │   └── pages/
    │       ├── Login.jsx       # Login/Register page
    │       └── Todos.jsx       # Todo management page
    └── vite.config.js          # Build config (API proxy)
```

---

## Learning Resources

This repository includes a comprehensive tutorial:
- **[TUTORIAL.md](./TUTORIAL.md)**: A complete guide (23 chapters) covering every line of code — from Go basics and Gin framework, through JWT authentication theory, to React Router and the full request lifecycle.

---

## Contributing

Feel free to fork this project and submit Pull Requests.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## License

Distributed under the MIT License. See `LICENSE` for more information.
