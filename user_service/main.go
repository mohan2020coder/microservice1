package main

import (
    "github.com/gin-gonic/gin"
    "user_service/routes"
)

func main() {
    r := gin.Default()
    routes.RegisterUserRoutes(r)
    r.Run(":8080") // Run on port 8080
}
