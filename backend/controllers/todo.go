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

// GetTodos returns all todos
func (tc *TodoController) GetTodos(c *gin.Context) {
	var todos []models.Todo
	result := tc.DB.Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateTodo creates a new todo
func (tc *TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := tc.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// UpdateTodo updates a todo by ID
func (tc *TodoController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	if err := tc.DB.First(&todo, id).Error; err != nil {
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

// DeleteTodo deletes a todo by ID
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	if err := tc.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	tc.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
