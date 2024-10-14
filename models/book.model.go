package models

import (
	"gorm.io/gorm"
)

type Book struct {
    gorm.Model
    Count       int        `json:"count"`
    Name        string     `json:"name"`
    Category    string     `json:"category"`
    Description string     `json:"description"`
}
