package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var jwtSecret = []byte("supersecretkey")
var db *gorm.DB

// Модель новости
type News struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
}

// Модель события
type Events struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
}

// Модель для JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// Подключение к базе данных

func connectDB() {
	dsn := "host=localhost port=5432 user=postgres password=;tybc123 dbname=UMS sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("⛔ Ошибка подключения к БД:", err)
	}
	db = database
	fmt.Println("✅ База данных подключена")
}

// Генерация JWT-токена
func generateToken(username, role string) (string, error) {
	claims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Middleware для аутентификации
func authMiddleware(requiredRole string) gin.HandlerFunc {
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

		if claims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

// Логин для получения JWT
func login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if req.Username == "admin" && req.Password == "admin123" {
		token, _ := generateToken(req.Username, "admin")
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
	}
}

// Получение всех событий
func getEvents(c *gin.Context) {
	var events []Events
	db.Find(&events)
	c.JSON(http.StatusOK, events)
}

// Добавление события
func createEvent(c *gin.Context) {
	var event Events
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	db.Create(&event)
	c.JSON(http.StatusOK, event)
}

// Лайк/дизлайк события
func reactEvent(c *gin.Context) {
	id := c.Param("id")
	var event Events

	if err := db.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Событие не найдено"})
		return
	}

	var req struct {
		Action string `json:"action"` // "like" или "dislike"
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if req.Action == "like" {
		event.Likes++
	} else if req.Action == "dislike" {
		event.Dislikes++
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверное действие"})
		return
	}

	db.Save(&event)
	c.JSON(http.StatusOK, event)
}

// Получение всех новостей
func getNews(c *gin.Context) {
	var news []News
	db.Find(&news)
	c.JSON(http.StatusOK, news)
}

// Добавление новости
func createNews(c *gin.Context) {
	var news News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	db.Create(&news)
	c.JSON(http.StatusOK, news)
}

// Лайк/дизлайк новости
func reactNews(c *gin.Context) {
	id := c.Param("id")
	var news News

	if err := db.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Новость не найдена"})
		return
	}

	var req struct {
		Action string `json:"action"` // "like" или "dislike"
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if req.Action == "like" {
		news.Likes++
	} else if req.Action == "dislike" {
		news.Dislikes++
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверное действие"})
		return
	}
	db.Save(&news)
	c.JSON(http.StatusOK, news)
}

// Обновление новости
func updateNews(c *gin.Context) {
	id := c.Param("id")
	var news News

	if err := db.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Новость не найдена"})
		return
	}

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	db.Save(&news)
	c.JSON(http.StatusOK, news)
}

// Удаление новости
func deleteNews(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&News{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Новость не найдена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Новость удалена"})
}

func main() {
	connectDB()
	r := gin.Default()

	r.POST("/login", login)
	r.GET("/news", getNews)
	r.POST("/news/:id/react", reactNews)

	r.GET("/events", getEvents)
	r.POST("/events/:id/react", reactEvent)

	admin := r.Group("/admin")
	admin.Use(authMiddleware("admin"))
	{
		admin.POST("/news", createNews)
		admin.PUT("/news/:id", updateNews)
		admin.DELETE("/news/:id", deleteNews)

		admin.POST("/events", createEvent)
	}

	r.Run(":8080")
}
