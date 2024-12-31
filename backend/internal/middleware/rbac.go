package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/MoiMoiTan/linh-san-store/internal/models"
	"net/http"
)

func RBACMiddleware(requiredRole models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy user từ context (đã được set bởi AuthMiddleware)
		userInterface, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, ok := userInterface.(*models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
			c.Abort()
			return
		}

		// Kiểm tra role
		if user.Role != requiredRole && user.Role != models.AdminRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
