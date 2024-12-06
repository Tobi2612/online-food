package http

import (
	"online-food/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	ProductRepo domain.ProductRepository
	OrderRepo   domain.OrderRepository
}

func RunHttpServer(config Config) *gin.Engine {
	app := gin.Default()

	corsConfig := initCors()

	app.Use(cors.New(corsConfig))
	//csrf ???

	// setup routes
	setupRouter(app, config)

	return app
}
