package main

import (
	"UMS/config"
	"UMS/internal/auth"
	"UMS/internal/document"
	"UMS/internal/expenses"
	"UMS/internal/feedback"
	"UMS/middleware"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 🔄 Подключение к базе данных
	log.Println("🔄 Подключение к базе данных...")
	config.ConnectDB()

	// Получаем подключение к БД
	db := config.DB

	// Инициализация сервисов
	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	fileRepo := document.NewFileRepository(db)
	fileService := document.NewFileService(fileRepo)
	fileHandler := document.NewFileHandler(fileService)

	// Инициализация модулей
	expRepo := expenses.NewExpenseRepository(db)
	expService := expenses.NewExpenseService(expRepo)
	expHandler := expenses.NewExpenseHandler(expService)

	chatRepo := feedback.NewChatRepository(db)
	chatService := feedback.NewChatService(chatRepo)
	chatHandler := feedback.NewChatHandler(chatService)

	// Создание маршрутизатора
	r := gin.Default()

	// 🔥 CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 🔑 Роуты авторизации
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	// 🔒 Защищённые роуты (JWT)
	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Доступ разрешен!"})
		})
		protectedRoutes.GET("/files", fileHandler.ListFiles)
		protectedRoutes.GET("/download/:id", fileHandler.DownloadFile)

		// 🔐 Маршруты коммунальных расходов
		protectedRoutes.POST("/expenses/calculate", expHandler.CalculateExpense)

		// 🔐 Маршруты чата
		protectedRoutes.POST("/chat/send", chatHandler.SendMessageHandler)
		protectedRoutes.GET("/chat/history", chatHandler.GetChatHistoryHandler)

		// 🔐 Админские маршруты
		adminRoutes := protectedRoutes.Group("/")
		adminRoutes.Use(middleware.AdminMiddleware())
		{
			adminRoutes.POST("/upload", fileHandler.UploadFile)
			adminRoutes.DELETE("/files/:id", fileHandler.DeleteFile)
		}
	}

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("🚀 Сервер запущен на http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
