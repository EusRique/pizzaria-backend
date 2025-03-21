package model

type Payment struct {
	Value    float64 `json:"value" binding:"required"`
	PedidoID uint    `json:"pedido_id"`
}
