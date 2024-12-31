package models

import (
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

type Role string

const (
    AdminRole Role = "admin"
    UserRole  Role = "user"
    GuestRole Role = "guest"
)

type User struct {
    gorm.Model
    Username     string `gorm:"unique;not null"`
    Email        string `gorm:"unique;not null"`
    Password     string `gorm:"not null"`
    Role         Role   `gorm:"type:enum('admin','user','guest');default:'user'"`
    FullName     string
    PhoneNumber  string
    Address      string
    Orders       []Order
}

func (u *User) BeforeSave(tx *gorm.DB) error {
    if u.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        u.Password = string(hashedPassword)
    }
    return nil
}
