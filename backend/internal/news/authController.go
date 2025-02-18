package news

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("supersecretkey")

// Генерация JWT
func generateToken(username, role string) (string, error) {
	claims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Логин
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Простая проверка пользователей (можно заменить на БД)
	users := map[string]string{
		"admin": "admin123",
		"user":  "user123",
	}

	role := "user"
	if req.Username == "admin" && req.Password == users["admin"] {
		role = "admin"
	} else if req.Username == "user" && req.Password == users["user"] {
		role = "user"
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	token, _ := generateToken(req.Username, role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
