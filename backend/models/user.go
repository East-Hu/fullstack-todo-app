package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;size:50"`
	Password string `json:"-"` // bcrypt hash, never returned to frontend
}
