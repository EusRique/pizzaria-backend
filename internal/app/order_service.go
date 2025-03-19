package app

import (
	"github.com/EusRique/pizzaria-backend/internal/domain"
	"github.com/EusRique/pizzaria-backend/internal/infra/repositories"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{repo: repositories.NewOrderRepository()}
}

func (s *OrderService) CreateOrder(customer, address, phone string, items []domain.OrderItem) (uint, error) {
	order, err := domain.NewOrder(customer, address, phone, items)

	return order.ID, err
}

func (s *OrderService) ListOrders() ([]domain.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *OrderService) UpdateStatus(id uint, status string) error {
	return s.repo.UpdateOrderStatus(id, status)
}
