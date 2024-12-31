package utils

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/MoiMoiTan/linh-san-store/internal/models"
)

var jwtSecret = []byte("your-secret-key") // Trong thực tế nên đặt trong env

type Claims struct {
	UserID uint
	Role   models.Role
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	// Tạo claims
	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Tạo token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Ký token
	return token.SignedString(jwtSecret)
}

func VerifyToken(tokenString string) (*Claims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Kiểm tra token hợp lệ
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
