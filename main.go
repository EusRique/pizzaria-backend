package main

import (
	"fmt"

	"github.com/EusRique/pizzaria-backend/config"
	"github.com/EusRique/pizzaria-backend/internal/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()
	pizzaHandler := interfaces.NewPizzaHandler()

	r.POST("/pizzas", pizzaHandler.CreatePizza)
	r.GET("/pizzas", pizzaHandler.ListPizzas)

	r.POST("/order", interfaces.NewOrderHandler().CreateOrder)
	r.GET("/orders", interfaces.NewOrderHandler().ListOrders)
	r.PUT("/orders/:id/status", interfaces.NewOrderHandler().UpdateOrderStatus)

	fmt.Println("Server running on port 8080")
	r.Run(":3000")
}
