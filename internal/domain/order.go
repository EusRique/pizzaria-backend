package domain

import (
	"errors"
	"time"
)

type Order struct {
	ID        uint        `gorm:"primaryKey"`
	Customer  string      `gorm:"type:varchar(255); not null"`
	Address   string      `gorm:"type:varchar(255); not null"`
	Phone     string      `gorm:"type:varchar(20); not null"`
	Total     float64     `gorm:"type:decimal; not null"`
	CreatedAt time.Time   `gorm:"type:timestamp; not null"`
	Status    string      `gorm:"default:'pending'"`
	Items     []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	PizzaID   uint    `gorm:"not null"`
	Quantity  uint    `gorm:"not null"`
	UnitPrice float64 `gorm:"type:decimal; not null"`
}

func NewOrder(customer, address, phone string, items []OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("o pedido deve ter pelo menos um item")
	}

	total := 0.0
	for _, item := range items {
		total += item.UnitPrice * float64(item.Quantity)
	}

	pedido := Order{
		Customer: customer,
		Address:  address,
		Phone:    phone,
		Total:    total,
		Items:    items,
	}

	return &pedido, nil
}
