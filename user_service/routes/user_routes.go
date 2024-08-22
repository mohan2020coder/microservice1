package routes

import (
    "github.com/gin-gonic/gin"
    "user_service/controllers"
)


// RegisterUserRoutes registers routes related to users
func RegisterUserRoutes(r *gin.Engine) {
    userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", controllers.GetAllUsers)       // Get all users
        userRoutes.GET("/:id", controllers.GetUser)        // Get a user by ID
        userRoutes.POST("/", controllers.CreateUser)       // Create a new user
        userRoutes.PUT("/:id", controllers.UpdateUser)     // Update a user by ID
        userRoutes.DELETE("/:id", controllers.DeleteUser)  // Delete a user by ID
    }
}