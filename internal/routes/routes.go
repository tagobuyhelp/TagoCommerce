package routes

import (
    "github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(r *gin.Engine) {
    // API group
    api := r.Group("/api")
    
    // Health check endpoint
    api.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
            "message": "TagoCommerce API is running",
        })
    })

    // TODO: Add more route groups here
    // setupAuthRoutes(api)
    // setupProductRoutes(api)
    // setupCartRoutes(api)
    // setupOrderRoutes(api)
}

// Uncomment and implement these functions as you develop your API
/*
func setupAuthRoutes(rg *gin.RouterGroup) {
    auth := rg.Group("/auth")
    
    auth.POST("/register", handlers.RegisterUser)
    auth.POST("/login", handlers.LoginUser)
}

func setupProductRoutes(rg *gin.RouterGroup) {
    products := rg.Group("/products")
    
    products.GET("/", handlers.GetProducts)
    products.GET("/:id", handlers.GetProduct)
    products.POST("/", middleware.AuthMiddleware(), handlers.CreateProduct)
    products.PUT("/:id", middleware.AuthMiddleware(), handlers.UpdateProduct)
    products.DELETE("/:id", middleware.AuthMiddleware(), handlers.DeleteProduct)
}
*/