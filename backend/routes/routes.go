package routes

import (
	"fullstack-todo-app/controllers"
	"fullstack-todo-app/middleware"

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
	authCtrl := controllers.NewAuthController(db)
	todoCtrl := controllers.NewTodoController(db)

	// API routes group
	api := r.Group("/api")
	{
		// Public routes (no auth required)
		api.POST("/register", authCtrl.Register)
		api.POST("/login", authCtrl.Login)

		// Protected routes (auth required)
		authorized := api.Group("/")
		authorized.Use(middleware.JWTAuth())
		{
			authorized.GET("/todos", todoCtrl.GetTodos)
			authorized.POST("/todos", todoCtrl.CreateTodo)
			authorized.PUT("/todos/:id", todoCtrl.UpdateTodo)
			authorized.DELETE("/todos/:id", todoCtrl.DeleteTodo)
		}
	}

	return r
}
