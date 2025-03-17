package app

import (
	"github.com/EusRique/pizzaria-backend/internal/domain"
	"github.com/EusRique/pizzaria-backend/internal/infra/repositories"
)

type PizzaService struct {
	repo *repositories.PizzaRepository
}

func NewPizzaService() *PizzaService {
	return &PizzaService{repo: repositories.NewPizzaRepository()}
}

func (s *PizzaService) CreatePizza(name, description string, price float64, imageURL string) error {
	pizza := domain.Pizza{
		Name:        name,
		Description: description,
		Price:       price,
		ImageURL:    imageURL,
	}

	return s.repo.Create(&pizza)
}

func (s *PizzaService) ListPizzas() ([]domain.Pizza, error) {
	return s.repo.GetAll()
}
