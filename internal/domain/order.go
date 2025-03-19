package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/EusRique/pizzaria-backend/internal/model"
)

type Order struct {
	ID        uint        `gorm:"primaryKey"`
	Customer  string      `gorm:"type:varchar(255); not null"`
	Address   string      `gorm:"type:varchar(255); not null"`
	Phone     string      `gorm:"type:varchar(20); not null"`
	Total     float64     `gorm:"type:decimal; not null"`
	CreatedAt time.Time   `gorm:"type:timestamp; not null"`
	UpdatedAt time.Time   `gorm:"type:timestamp"`
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

func (order *Order) IsValid() error {
	if strings.TrimSpace(order.Customer) == "" {
		return fmt.Errorf("nome é obrigatório")
	}

	if strings.TrimSpace(order.Address) == "" {
		return fmt.Errorf("endereço é obrigatório")
	}

	if len(order.Phone) < 8 {
		return fmt.Errorf("telefone é obrigatório")
	}

	if len(order.Items) == 0 {
		return fmt.Errorf("o pedido deve ter pelo menos um item")
	}

	for _, item := range order.Items {
		if item.Quantity == 0 {
			return fmt.Errorf("quantidade inválida")
		}

		if item.UnitPrice == 0 {
			return fmt.Errorf("preço inválido")
		}
	}

	return nil
}

func NewOrder(order model.Order, items []model.OrderItem) (*Order, error) {
	total := 0.0
	orderItems := make([]OrderItem, len(items))
	for i, item := range items {
		total += item.UnitPrice * float64(item.Quantity)
		orderItems[i] = OrderItem{
			PizzaID:   item.PizzaID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
	}

	orderPurchase := Order{
		Customer: order.Customer,
		Address:  order.Address,
		Phone:    order.Phone,
		Total:    total,
		Items:    orderItems,
	}

	if err := orderPurchase.IsValid(); err != nil {
		return nil, err
	}

	return &orderPurchase, nil
}
