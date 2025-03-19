package model

import "github.com/EusRique/pizzaria-backend/internal/domain"

type Order struct {
	Customer string             `json:"customer"`
	Address  string             `json:"address"`
	Phone    string             `json:"phone"`
	Items    []domain.OrderItem `json:"items"`
}

type OrderStatus struct {
	Status string `json:"status"`
}
