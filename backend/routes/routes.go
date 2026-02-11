package routes

import (
	"fullstack-todo-app/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter configures all routes and middleware
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// CORS middleware - allow React dev server
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize controllers
	todoCtrl := controllers.NewTodoController(db)

	// API routes group
	api := r.Group("/api")
	{
		api.GET("/todos", todoCtrl.GetTodos)
		api.POST("/todos", todoCtrl.CreateTodo)
		api.PUT("/todos/:id", todoCtrl.UpdateTodo)
		api.DELETE("/todos/:id", todoCtrl.DeleteTodo)
	}

	return r
}
