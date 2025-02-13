package middleware

import (
	"UMS/utils"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Println("📌 Authorization Header:", authHeader) // Проверяем, что токен передаётся

		if authHeader == "" {
			log.Println("❌ Ошибка: Токен отсутствует")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization not found"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("❌ Ошибка: Неверный формат токена")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid format token"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			log.Println("❌ Ошибка: Токен недействителен -", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		log.Println("✅ Успешная аутентификация: user_id =", claims.UserID, "role =", claims.Role)
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "role must be admin"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,          // Разрешить все домены
		AllowMethods:     []string{"*"}, // Разрешить все методы
		AllowHeaders:     []string{"*"}, // Разрешить все заголовки
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
