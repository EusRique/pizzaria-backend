package model

type Order struct {
	Customer string      `json:"customer" binding:"required"`
	Address  string      `json:"address" binding:"required"`
	Phone    string      `json:"phone" binding:"required"`
	Items    []OrderItem `json:"items" binding:"required"`
}

type OrderItem struct {
	PizzaID   uint    `json:"pizza_id" binding:"required"`
	Quantity  uint    `json:"quantity" binding:"required"`
	UnitPrice float64 `json:"unit_price" binding:"required"`
}
type OrderStatus struct {
	Status string `json:"status" binding:"required"`
}
