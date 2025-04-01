package app

import (
	"fmt"

	"github.com/EusRique/pizzaria-backend/internal/domain"
	"github.com/EusRique/pizzaria-backend/internal/infra/repositories"
	"github.com/EusRique/pizzaria-backend/internal/model"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{repo: repositories.NewOrderRepository()}
}

func (s *OrderService) CreateOrder(order model.Order) (uint, error) {
	purchaseOrder, err := domain.NewOrder(order, order.Items)
	if err != nil {
		return 0, err
	}
	err = s.repo.CreateOrder(purchaseOrder)
	return purchaseOrder.ID, err
}

func (s *OrderService) ListOrders() ([]domain.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *OrderService) UpdateStatus(id uint, status string) error {
	_, err := s.repo.GetOrderById(id)
	if err != nil {
		return fmt.Errorf("pedido não encontrado")
	}

	return s.repo.UpdateOrderStatus(id, status)
}

func (s *OrderService) MarkOrderAsPaid(id uint) error {
	return s.repo.MarkOrderAsPaid(id)
}

func (s *OrderService) ListOrdersByStatus(paid bool) ([]domain.Order, error) {
	var orders []domain.Order
	err := s.repo.GetOrderByStatus(&orders, paid)
	return orders, err
}
