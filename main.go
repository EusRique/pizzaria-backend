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

	fmt.Println("Server running on port 8080")
	r.Run(":3000")
}
