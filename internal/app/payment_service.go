package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mercadopago/sdk-go/pkg/config"

	"github.com/mercadopago/sdk-go/pkg/payment"
)

type PaymentService struct {
	client payment.Client
}

func NewPaymentService() *PaymentService {
	accessToken := os.Getenv("MP_ACCESS_TOKEN")
	mpConfig, err := config.New(accessToken)
	if err != nil {
		log.Fatalf("Erro ao configurar Mercado Pago: %v", err)
	}

	return &PaymentService{client: payment.NewClient(mpConfig)}
}

func (s *PaymentService) CreatePaymentPix(value float64, PedidoID uint, paymentEmail string) (*payment.Response, error) {
	req := payment.Request{
		TransactionAmount: value,
		Description:       fmt.Sprintf("Pagamento do pedido: #%d", PedidoID),
		PaymentMethodID:   "pix",
		Payer: &payment.PayerRequest{
			Email: "pagante@gmail.com",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pagamento PIX: %w", err)
	}
	fmt.Println("#RESPOSTAAAA", resp)
	return resp, nil
}
