package repositories

import (
	"github.com/EusRique/pizzaria-backend/config"
	"github.com/EusRique/pizzaria-backend/internal/domain"

	"gorm.io/gorm"
)

type PizzaRepository struct {
	db *gorm.DB
}

func NewPizzaRepository() *PizzaRepository {
	return &PizzaRepository{db: config.DB}
}

func (r *PizzaRepository) Create(pizza *domain.Pizza) error {
	return r.db.Create(pizza).Error
}

func (r *PizzaRepository) GetAll() ([]domain.Pizza, error) {
	var pizzas []domain.Pizza
	err := r.db.Find(&pizzas).Error
	return pizzas, err
}

func (r *PizzaRepository) GetByID(id uint) (*domain.Pizza, error) {
	var pizza domain.Pizza
	err := r.db.First(&pizza, id).Error
	return &pizza, err
}
