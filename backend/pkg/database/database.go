package database

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/MoiMoiTan/linh-san-store/internal/models"
)

func ConnectDatabase() (*gorm.DB, error) {
    username := "root"         // Mặc định của XAMPP
    password := ""            // Mặc định để trống
    host := "localhost"
    port := "3306"
    dbname := "linh_san_store"

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        username,
        password,
        host,
        port,
        dbname,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    // Auto Migrate các models
    err = db.AutoMigrate(
        &models.User{},
        &models.Product{},
        &models.Order{},
        &models.OrderItem{},
    )
    if err != nil {
        return nil, fmt.Errorf("failed to migrate database: %v", err)
    }

    return db, nil
}
