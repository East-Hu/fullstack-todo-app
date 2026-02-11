package main

import (
	"fmt"
	"log"
	"os"

	"fullstack-todo-app/models"
	"fullstack-todo-app/routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Read config from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	// MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Auto migrate models
	db.AutoMigrate(&models.User{}, &models.Todo{})

	fmt.Println("✅ MySQL connected and migrated!")
	fmt.Printf("🚀 Server starting at http://localhost:%s\n", serverPort)
	fmt.Println("📝 API endpoints:")
	fmt.Println("   POST   /api/register    - Register a new user")
	fmt.Println("   POST   /api/login       - Login and get JWT token")
	fmt.Println("   GET    /api/todos       - List your todos (auth required)")
	fmt.Println("   POST   /api/todos       - Create a todo (auth required)")
	fmt.Println("   PUT    /api/todos/:id   - Update a todo (auth required)")
	fmt.Println("   DELETE /api/todos/:id   - Delete a todo (auth required)")

	// Setup router and start server
	r := routes.SetupRouter(db)
	r.Run(":" + serverPort)
}
