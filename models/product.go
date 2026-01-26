package models

type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name" binding:"required,min=2"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Quantity int     `json:"quantity" binding:"required,gte=0"`
}
