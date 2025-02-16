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
	log.Println("🔄 Connecting to DB")
	config.ConnectDB()

	db := config.DB

	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	fileRepo := document.NewFileRepository(db)
	fileService := document.NewFileService(fileRepo)
	fileHandler := document.NewFileHandler(fileService)

	expRepo := expenses.NewExpenseRepository(db)
	expService := expenses.NewExpenseService(expRepo)
	expHandler := expenses.NewExpenseHandler(expService)

	chatRepo := feedback.NewChatRepository(db)
	chatService := feedback.NewChatService(chatRepo)
	chatHandler := feedback.NewChatHandler(chatService)

	r := gin.Default()

	// Разрешаем фронту (localhost:3000) делать запросы
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 📌 Маршруты API
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Доступ разрешен!"})
		})
		protectedRoutes.GET("/files", fileHandler.ListFiles)
		protectedRoutes.GET("/download/:id", fileHandler.DownloadFile)

		protectedRoutes.POST("/expenses/calculate", expHandler.CalculateExpense)

		protectedRoutes.POST("/chat/send", chatHandler.SendMessageHandler)
		protectedRoutes.GET("/chat/history", chatHandler.GetChatHistoryHandler)

		adminRoutes := protectedRoutes.Group("/admin")
		adminRoutes.Use(middleware.AdminMiddleware())
		{
			adminRoutes.POST("/upload", fileHandler.UploadFile)
			adminRoutes.DELETE("/files/:id", fileHandler.DeleteFile)
		}
	}

	r.Static("/static", "./public")

	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("🚀 Сервер запущен на http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
