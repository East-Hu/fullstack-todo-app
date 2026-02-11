package main

import (
	"fmt"
	"log"

	"fullstack-todo-app/models"
	"fullstack-todo-app/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// MySQL DSN (Data Source Name) 格式:
	// 用户名:密码@tcp(主机:端口)/数据库名?参数
	dsn := "todouser:todopass@tcp(127.0.0.1:3307)/todo_app?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Auto migrate the Todo model
	db.AutoMigrate(&models.Todo{})

	fmt.Println("✅ MySQL connected and migrated!")
	fmt.Println("🚀 Server starting at http://localhost:8080")
	fmt.Println("📝 API endpoints:")
	fmt.Println("   GET    /api/todos      - List all todos")
	fmt.Println("   POST   /api/todos      - Create a todo")
	fmt.Println("   PUT    /api/todos/:id   - Update a todo")
	fmt.Println("   DELETE /api/todos/:id   - Delete a todo")

	// Setup router and start server
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
