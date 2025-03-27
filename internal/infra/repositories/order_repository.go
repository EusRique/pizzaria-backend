package repositories

import (
	"github.com/EusRique/pizzaria-backend/config"
	"github.com/EusRique/pizzaria-backend/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{db: config.DB}
}

func (r *OrderRepository) CreateOrder(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetAllOrders() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetOrderById(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) UpdateOrderStatus(id uint, status string) error {
	return r.db.Model(&domain.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *OrderRepository) MarkOrderAsPaid(id uint) error {
	return r.db.Model(&domain.Order{}).Where("id = ?", id).Update("paid", true).Error
}
