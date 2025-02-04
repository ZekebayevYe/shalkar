package main

import (
	"UMS/config"
	"UMS/internal/auth"
	"UMS/internal/document"
	"UMS/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	config.ConnectDB()
	db := config.DB

	// Создание репозиториев, сервисов и хендлеров
	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	fileRepo := document.NewFileRepository(db)
	fileService := document.NewFileService(fileRepo)
	fileHandler := document.NewFileHandler(fileService)

	// Инициализация роутера
	r := gin.Default()

	// Группировка роутов
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware()) // Middleware для проверки JWT
	{
		protectedRoutes.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Доступ разрешен!"})
		})

		// Маршруты для работы с файлами
		protectedRoutes.POST("/upload", middleware.AdminMiddleware(), fileHandler.UploadFile)
		protectedRoutes.GET("/files", fileHandler.ListFiles)
		protectedRoutes.GET("/download/:id", fileHandler.DownloadFile)
	}

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Println("Сервер запущен на порту", port)
	r.Run(":" + port)
}
