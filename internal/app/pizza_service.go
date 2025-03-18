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

func (s *PizzaService) CreatePizza(name, description string, price float64, image string) error {
	newPizza, err := domain.NewPizza(name, description, price, image)
	if err != nil {
		return err
	}

	return s.repo.Create(newPizza)
}

func (s *PizzaService) ListPizzas() ([]domain.Pizza, error) {
	return s.repo.GetAll()
}
