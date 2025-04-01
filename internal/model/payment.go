package model

type PaymentPix struct {
	Value    float64 `json:"value" binding:"required"`
	PedidoID uint    `json:"pedido_id"`
}

type PaymentCreditCard struct {
	PedidoID      uint    `json:"pedido_id"`
	Value         float64 `json:"value" binding:"required"`
	Token         string  `json:"token" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
	Installments  int     `json:"installments" binding:"required"`
}
