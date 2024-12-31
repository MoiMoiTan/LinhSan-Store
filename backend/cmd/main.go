package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/MoiMoiTan/linh-san-store/pkg/database"
    "github.com/MoiMoiTan/linh-san-store/internal/handlers"
    "github.com/MoiMoiTan/linh-san-store/internal/middleware"
    "github.com/MoiMoiTan/linh-san-store/internal/models"
)

func main() {
    // Kết nối database
    db, err := database.ConnectDatabase()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Kiểm tra kết nối
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get database instance: %v", err)
    }

    // Test the connection
    err = sqlDB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Println("Successfully connected to database!")

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(db)
    productHandler := handlers.NewProductHandler(db)

    // Setup router
    r := gin.Default()

    // Public routes
    r.POST("/api/auth/login", authHandler.Login)
    r.GET("/api/products", productHandler.GetProducts)

    // Protected routes
    authorized := r.Group("/api")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.POST("/products", middleware.RBACMiddleware(models.AdminRole), productHandler.CreateProduct)
    }

    r.Run(":8080")
}
