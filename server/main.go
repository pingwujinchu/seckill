package main

import (
	"log"
	"server/pkg/cache"
	models "server/pkg/model"
	RabbitMQ "server/pkg/rabitmq"
	routers "server/pkg/router"
)

func main() {
	models.Init()
	cache.InitClient()
	RabbitMQ.InitRabbitMQ()
	server := routers.InitRouter()
	err := server.Engine.Run(":8080")
	if err != nil {
		log.Fatal("start failed")
	}
}
