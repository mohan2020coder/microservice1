package routes

import (
    "github.com/gin-gonic/gin"
    "order_service/controllers"
)

func RegisterOrderRoutes(r *gin.Engine) {
    orderRoutes := r.Group("/orders")
    {
        orderRoutes.GET("/", controllers.GetAllOrders)       // Get all orders
        orderRoutes.GET("/:id", controllers.GetOrderByID)    // Get an order by ID
        orderRoutes.POST("/", controllers.CreateOrder)       // Create a new order
        orderRoutes.PUT("/:id", controllers.UpdateOrder)     // Update an order by ID
        orderRoutes.DELETE("/:id", controllers.DeleteOrder)  // Delete an order by ID
    }
}