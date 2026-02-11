package controllers

import (
	"net/http"

	"fullstack-todo-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TodoController holds the database connection
type TodoController struct {
	DB *gorm.DB
}

// NewTodoController creates a new TodoController
func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{DB: db}
}

// getUserID extracts the current user's ID from the gin context
func getUserID(c *gin.Context) uint {
	return c.MustGet("userID").(uint)
}

// GetTodos returns all todos for the current user
func (tc *TodoController) GetTodos(c *gin.Context) {
	userID := getUserID(c)
	var todos []models.Todo
	result := tc.DB.Where("user_id = ?", userID).Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateTodo creates a new todo for the current user
func (tc *TodoController) CreateTodo(c *gin.Context) {
	userID := getUserID(c)
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = userID
	result := tc.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// UpdateTodo updates a todo by ID (only if it belongs to the current user)
func (tc *TodoController) UpdateTodo(c *gin.Context) {
	userID := getUserID(c)
	id := c.Param("id")

	var todo models.Todo
	if err := tc.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != nil {
		todo.Title = *input.Title
	}
	if input.Completed != nil {
		todo.Completed = *input.Completed
	}

	tc.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo deletes a todo by ID (only if it belongs to the current user)
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	userID := getUserID(c)
	id := c.Param("id")

	var todo models.Todo
	if err := tc.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	tc.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
