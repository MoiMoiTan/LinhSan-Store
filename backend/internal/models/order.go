package models

import "gorm.io/gorm"

type OrderStatus string

const (
    Pending   OrderStatus = "pending"
    Confirmed OrderStatus = "confirmed"
    Shipping  OrderStatus = "shipping"
    Delivered OrderStatus = "delivered"
    Cancelled OrderStatus = "cancelled"
)

type Order struct {
    gorm.Model
    UserID      uint        `gorm:"not null"`
    User        User
    Status      OrderStatus `gorm:"type:enum('pending','confirmed','shipping','delivered','cancelled');default:'pending'"`
    TotalAmount float64
    OrderItems  []OrderItem
    Address     string
    PhoneNumber string
}

type OrderItem struct {
    gorm.Model
    OrderID   uint    `gorm:"not null"`
    ProductID uint    `gorm:"not null"`
    Product   Product
    Quantity  int     `gorm:"not null"`
    Price     float64 `gorm:"not null"`
}
