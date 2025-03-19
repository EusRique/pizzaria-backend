package interfaces

import (
	"net/http"
	"strconv"

	"github.com/EusRique/pizzaria-backend/internal/app"
	"github.com/EusRique/pizzaria-backend/internal/model"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *app.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{service: app.NewOrderService()}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Dados inválidos"})
		return
	}

	id, err := h.service.CreateOrder(order.Customer, order.Address, order.Phone, order.Items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating order", "message": err.Error()})
		return
	}

	c.JSON((http.StatusCreated), gin.H{"message": "Pedido criado com sucesso!", "pedido_id": id})
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.service.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Erro ao buscar pedidos"})
	}

	c.JSON(http.StatusCreated, orders)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var status model.OrderStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err = h.service.UpdateStatus(uint(id), status.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Erro ao atualizar status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso!"})
}
