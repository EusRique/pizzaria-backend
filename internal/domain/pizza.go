package domain

import (
	"fmt"
	"strings"
)

type Pizza struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(255); not null"`
	Description string  `gorm:"type:text; not null"`
	Price       float64 `gorm:"type:decimal; not null"`
	Image       string  `gorm:"type:varchar(255)"`
}

func (pizza *Pizza) IsValid() error {
	if strings.TrimSpace(pizza.Name) == "" {
		return fmt.Errorf("nome é obrigatório")
	}

	if strings.TrimSpace(pizza.Description) == "" {
		return fmt.Errorf("descrição é obrigatória")
	}

	if (pizza.Price) == 0 {
		return fmt.Errorf("preço é obrigatório")
	}

	return nil
}

func NewPizza(name, description string, price float64, image string) (*Pizza, error) {
	pizza := Pizza{
		Name:        name,
		Description: description,
		Price:       price,
		Image:       image,
	}

	if err := pizza.IsValid(); err != nil {
		return nil, err
	}

	return &pizza, nil
}
