package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name      string `json:"name" binding:"required"`
	Price     string `json:"price" binding:"required,number"`
	Rating    string `json:"rating" binding:"required,number"`
	Stock     string `json:"stock" binding:"required,number"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  time.Time
}
