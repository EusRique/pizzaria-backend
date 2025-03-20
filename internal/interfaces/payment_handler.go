package interfaces

import (
	"net/http"

	"github.com/EusRique/pizzaria-backend/internal/app"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *app.PaymentService
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{service: app.NewPaymentService()}
}

func (h *PaymentHandler) CreatePaymentPix(c *gin.Context) {
	var request struct {
		Value       float64 `json:"value" binding:"required"`
		Description string  `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Verifique os dados enviados"})
		return
	}

	qrCode, err := h.service.CreatePaymentPix(request.Value, request.Description, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar pagamento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"qr_code": qrCode})
}
