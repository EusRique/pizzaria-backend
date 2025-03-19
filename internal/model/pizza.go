package model

type Pizza struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Image       string  `json:"image" binding:"required"`
}
