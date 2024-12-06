package main

import (
	port "online-food/delivery/http"
	"online-food/pkg/logger"
	"online-food/repository/mongodb"
)

func main() {
	l, err := logger.InitLogger()

	if err != nil {
		panic(err)
	}

	l.Info("Starting App")

	repo := mongodb.New(l)

	httpConfig := port.Config{
		ProductRepo: repo.ProductRepo,
		OrderRepo:   repo.OrderRepo,
	}

	app := port.RunHttpServer(httpConfig)

	app.Run()
}
