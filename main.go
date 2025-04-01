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
	orderHandler := interfaces.NewOrderHandler()
	paymentsHandler := interfaces.NewPaymentHandler()
	webhooksHandler := interfaces.NewWebhookHandler()

	r.POST("/pizzas", pizzaHandler.CreatePizza)
	r.GET("/pizzas", pizzaHandler.ListPizzas)

	r.POST("/order", orderHandler.CreateOrder)
	r.GET("/orders", orderHandler.ListOrders)
	r.GET("/ordersstatus", orderHandler.ListOrdersByStatus)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	r.POST("/payments/pix", paymentsHandler.CreatePaymentPix)
	r.POST("/payments/creditcard", paymentsHandler.CreatePaymentCreditCard)
	r.POST("/webhook/payments", webhooksHandler.ProcessPayment)

	fmt.Println("Server running on port 8080")
	r.Run(":3000")
}
