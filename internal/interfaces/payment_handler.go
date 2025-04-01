package interfaces

import (
	"net/http"

	"github.com/EusRique/pizzaria-backend/internal/app"
	"github.com/EusRique/pizzaria-backend/internal/model"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *app.PaymentService
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{service: app.NewPaymentService()}
}

func (h *PaymentHandler) CreatePaymentPix(c *gin.Context) {
	var payments model.PaymentPix
	if err := c.ShouldBindJSON(&payments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Verifique os dados enviados"})
		return
	}

	qrCode, err := h.service.CreatePaymentPix(payments.Value, payments.PedidoID, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar pagamento", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": qrCode})
}

func (h *PaymentHandler) CreatePaymentCreditCard(c *gin.Context) {
	var payments model.PaymentCreditCard
	if err := c.ShouldBindJSON(&payments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Verifique os dados enviados"})
		return
	}

	status, err := h.service.CreatePaymentCreditCard(payments.Value, payments.PedidoID, payments.Token, payments.Email, payments.PaymentMethod, payments.Installments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao processar pagamento", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": status})
}
