package http

import (
	"online-food/delivery/http/order"
	"online-food/delivery/http/product"

	"github.com/gin-gonic/gin"
)

func setupRouter(app *gin.Engine, config Config) {
	// func setupRouter(app *gin.Engine) {

	app.POST("/ping", ping)

	api := app.Group("/api")

	orderRouter := api.Group("/order")
	productRouter := api.Group("/product")
	order.New(orderRouter, config.OrderRepo, config.ProductRepo)
	product.New(productRouter, config.ProductRepo)

}
