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

func (s *PaymentService) CreatePaymentPix(value float64, pedidoID uint, email string) (*payment.Response, error) {
	req := payment.Request{
		TransactionAmount: value,
		Description:       fmt.Sprintf("Pagamento do pedido: #%d", pedidoID),
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
	fmt.Println("#CODE PIX", resp)
	return resp, nil
}

func (s *PaymentService) CreatePaymentCreditCard(value float64, pedidoID uint, token, email, paymentMethod string, installments int) (*payment.Response, error) {
	req := payment.Request{
		TransactionAmount: value,
		Token:             token,
		Description:       fmt.Sprintf("Pagamento do pedido: #%d", pedidoID),
		Installments:      installments,
		PaymentMethodID:   paymentMethod, // Pode ser "master", "amex", etc.
		Payer: &payment.PayerRequest{
			Email: "pagante@gmail.com",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pagamento com Cartão de Crédito: %w", err)
	}
	fmt.Println("#CODE CREDIT CARD", resp)
	return resp, nil
}
