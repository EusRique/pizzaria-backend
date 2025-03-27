package interfaces

import (
	"net/http"
	"strconv"

	"github.com/EusRique/pizzaria-backend/internal/app"
	"github.com/EusRique/pizzaria-backend/internal/model"
	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	orderService *app.OrderService
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{orderService: app.NewOrderService()}
}

func (h *WebhookHandler) ProcessPayment(c *gin.Context) {
	var processPayment model.ProcessPayment
	if err := c.ShouldBindJSON(&processPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Verifique os dados enviados"})
	}

	if processPayment.Action == "payment.updated" {
		paymentID := processPayment.Data.ID
		// Verificar se chama API do mercado livre
		id, err := strconv.Atoi(paymentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
			return
		}
		h.orderService.MarkOrderAsPaid(uint(id))
		c.JSON(http.StatusOK, gin.H{"message": "Webhook processado"})
	}
}
