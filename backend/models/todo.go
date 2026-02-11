package models

import "gorm.io/gorm"

// Todo represents a todo item in the database
type Todo struct {
	gorm.Model
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"`
}
