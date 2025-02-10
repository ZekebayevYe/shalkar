package news

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretnews = []byte("supersecretkey")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// Middleware для проверки JWT
func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		// Проверка роли, если требуется
		if requiredRole != "" && claims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав"})
			c.Abort()
			return
		}

		// Сохраняем пользователя в контексте
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
