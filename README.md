# 🚀 Full-Stack Todo App (Go + React + Docker)

> A modern, full-stack application built to demonstrate the integration of **Go (Gin)** backend, **React** frontend, and **MySQL** database using **Docker Compose**.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg?style=flat&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB.svg?style=flat&logo=react)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED.svg?style=flat&logo=docker)

[🇺🇸 English](./README.md) | [🇨🇳 简体中文](./README.zh-CN.md)

---

## 📖 Introduction

This project is a perfect starting point for developers transitioning from algorithms/basics to **Full-Stack Development**. It implements a complete CRUD workflow with a clean architecture.

### ✨ Features
- **Frontend**: Built with **React** and **Vite** for a blazing fast UI experience.
- **Backend**: RESTful API powered by **Go (Golang)** and **Gin** framework.
- **Database**: **MySQL 8.0** managed via GORM (ORM library) for seamless data operations.
- **Containerization**: Entire database environment runs in **Docker**, keeping your local machine clean.
- **Hot Reload**: Instant feedback during development for both frontend and backend.

---

## 🛠 Tech Stack

| Component | Technology | Description |
|-----------|------------|-------------|
| **Frontend** | React, Vite, Axios | Modern UI library with fast build tool |
| **Backend** | Go, Gin, GORM | High-performance compiled language |
| **Database** | MySQL 8.0 | Industrial-standard relational database |
| **Infrastructure** | Docker Compose | Container orchestration for development |

---

## 🚀 Getting Started

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Go](https://go.dev/) (v1.21+)
- [Node.js](https://nodejs.org/) (v18+)

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/YOUR_USERNAME/fullstack-todo-app.git
cd fullstack-todo-app
```

### 2️⃣ Start the Database (Docker)
We use Docker for MySQL so you don't have to install it manually.
```bash
docker compose up -d
```
> This starts MySQL on port `3307` (mapped from container's 3306).

### 3️⃣ Start the Backend
```bash
cd backend
go mod tidy       # Install Go dependencies
go run main.go    # Start API server at http://localhost:8080
```

### 4️⃣ Start the Frontend
Open a **new terminal**:
```bash
cd frontend
npm install       # Install Node dependencies
npm run dev       # Start UI at http://localhost:5173
```

🎉 **That's it! Visit [http://localhost:5173](http://localhost:5173) to see your app.**

---

## 📂 Project Structure

```bash
fullstack-todo-app/
├── docker-compose.yml    # Database configuration
├── TUTORIAL.md           # 📚 Detailed Step-by-Step Tutorial
├── ADVANCED_ROADMAP.md   # 🚀 Future Learning Path (Mini-Amazon)
├── backend/
│   ├── main.go           # Entry point (Server & DB connection)
│   ├── models/           # Database schemas (GORM structs)
│   ├── controllers/      # Business logic & Handlers
│   └── routes/           # API Route definitions
└── frontend/
    ├── src/
    │   ├── App.jsx       # Main UI Component
    │   └── main.jsx      # React Entry point
    └── vite.config.js    # Build configuration
```

---

## 📚 Learning Resources

This repository includes a comprehensive internal tutorial:
- **[TUTORIAL.md](./TUTORIAL.md)**: A complete guide explaining every line of code, Docker setup, and basic Go syntax.
- **[ADVANCED_ROADMAP.md](./ADVANCED_ROADMAP.md)**: A roadmap to take this project to the next level (Authentication, Redis, Microservices).

---

## 🤝 Contributing

Feel free to fork this project and submit Pull Requests.
Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
