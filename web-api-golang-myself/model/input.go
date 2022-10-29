package model

type Input struct {
	Name   string `json:"name" binding:"required"`
	Price  int    `json:"price" binding:"required,number"`
	Rating int    `json:"rating" binding:"required,number"`
	Stock  int    `json:"stock" binding:"required,number"`
}
