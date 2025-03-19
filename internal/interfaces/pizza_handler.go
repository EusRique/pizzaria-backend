package interfaces

import (
	"log"
	"net/http"

	"github.com/EusRique/pizzaria-backend/internal/app"
	"github.com/EusRique/pizzaria-backend/internal/model"
	"github.com/gin-gonic/gin"
)

type PizzaHandler struct {
	service *app.PizzaService
}

func NewPizzaHandler() *PizzaHandler {
	return &PizzaHandler{service: app.NewPizzaService()}
}

func (h *PizzaHandler) CreatePizza(c *gin.Context) {
	var pizza model.Pizza

	if err := c.ShouldBindJSON(&pizza); err != nil {
		log.Println("Invalid data:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Verifique os dados enviados"})
		return
	}

	err := h.service.CreatePizza(pizza.Name, pizza.Description, pizza.Price, pizza.Image)
	if err != nil {
		log.Println("Error creating pizza:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating pizza", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pizza criada com sucesso!"})
}

func (h *PizzaHandler) ListPizzas(c *gin.Context) {
	pizzas, err := h.service.ListPizzas()
	if err != nil {
		log.Println("Error listing pizzas:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Erro ao buscar pizzas"})
		return
	}

	c.JSON(http.StatusOK, pizzas)
}
