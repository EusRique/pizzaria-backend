package domain

import "time"

type Pedido struct {
	ID       uint         `gorm:"primaryKey"`
	Cliente  string       `gorm:"type:varchar(255); not null"`
	Endereco string       `gorm:"type:varchar(255); not null"`
	Total    float64      `gorm:"type:decimal; not null"`
	CriadoEm time.Time    `gorm:"type:timestamp; not null"`
	Status   string       `gorm:"default:'pendente'"`
	Itens    []ItemPedido `gorm:"foreignKey:PedidoID"`
}

type ItemPedido struct {
	ID            uint    `gorm:"primaryKey"`
	PedidoID      uint    `gorm:"not null"`
	PizzaID       uint    `gorm:"not null"`
	Quantidade    uint    `gorm:"not null"`
	PrecoUnitario float64 `gorm:"type:decimal; not null"`
}
