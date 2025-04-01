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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Verifique os dados enviados"})
		return
	}

	id, err := h.service.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao criar o pedido", "error": err.Error()})
		return
	}

	c.JSON((http.StatusCreated), gin.H{"message": "Pedido criado com sucesso!", "pedido_id": id})
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.service.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar pedidos", "error": err.Error()})
	}

	c.JSON(http.StatusCreated, orders)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var status model.OrderStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Verifique os dados enviados"})
		return
	}

	err = h.service.UpdateStatus(uint(id), status.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Pedido não encontrado ou erro ao atualizar status", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso!"})
}

func (h *OrderHandler) MarkOrderAsPaid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	err = h.service.MarkOrderAsPaid(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar status do pagamento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pagamento atualizado com sucesso!"})
}

func (h *OrderHandler) ListOrdersByStatus(c *gin.Context) {
	status := c.Query("paid")

	paid := false
	if status == "true" {
		paid = true
	}

	orders, err := h.service.ListOrdersByStatus(paid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar pedidos"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
