package main

import (
	"order_service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterOrderRoutes(r)
	r.Run(":8080") // Run on port 8080
}
