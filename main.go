package main

import (
	"deepak.com/web_rest/db"
	"deepak.com/web_rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
