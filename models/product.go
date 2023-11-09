package models

type Product struct {
	ID    uint    `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type CreateProductInput struct {
	Name  string  `json:"name" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}
