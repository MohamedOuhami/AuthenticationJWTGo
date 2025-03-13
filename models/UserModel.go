package models

import (
	"gorm.io/gorm"
)

type User struct {

	// Adds a template containing the ID, createdAt and other fields
	gorm.Model
	Email     string `gorm:"unique"` // A pointer to a string, allowing for null values
	Password  string
	FirstName string
	LastName  string
}
