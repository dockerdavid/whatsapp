package main

import (
	"compi-whatsapp/pkg/client"
	"compi-whatsapp/pkg/middlewares"
	"compi-whatsapp/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	server := gin.Default()

	middlewares.CorsMiddleware(server)

	routes.InitRoutes(server)

	client.InitClient()

	log.Fatal(server.Run(":8080"))
}
