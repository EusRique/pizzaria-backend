package interfaces

import (
	"log"
	"net/http"

	"github.com/EusRique/pizzaria-backend/internal/app"
	"github.com/gin-gonic/gin"
)

type PizzaHandler struct {
	service *app.PizzaService
}

func NewPizzaHandler() *PizzaHandler {
	return &PizzaHandler{service: app.NewPizzaService()}
}

func (h *PizzaHandler) CreatePizza(c *gin.Context) {
	var request struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
		ImageURL    string  `json:"image_url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error creating pizza:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating pizza"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pizza created successfully!"})
}

func (h *PizzaHandler) ListPizzas(c *gin.Context) {
	pizzas, err := h.service.ListPizzas()
	if err != nil {
		log.Println("Error listing pizzas:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing pizzas"})
		return
	}

	c.JSON(http.StatusOK, pizzas)
}
