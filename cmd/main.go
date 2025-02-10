package main

import (
	"UMS/config"
	"UMS/internal/auth"
	"UMS/internal/document"
	"UMS/middleware"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// Логирование процесса подключения к БД
	log.Println("🔄 Подключение к базе данных...")
	config.ConnectDB()
	db := config.DB

	// Инициализация сервисов
	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	fileRepo := document.NewFileRepository(db)
	fileService := document.NewFileService(fileRepo)
	fileHandler := document.NewFileHandler(fileService)

	// Создание маршрутизатора
	r := gin.Default()

	// 🔥 CORS Middleware (разрешаем все домены, если нужно только фронту - укажи его)
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Разрешаем запросы с любых доменов
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())  // Логирование запросов
	r.Use(gin.Recovery()) // Восстановление после ошибок

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
		// ✅ Доступ для всех пользователей
		protectedRoutes.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Доступ разрешен!"})
		})
		protectedRoutes.GET("/files", fileHandler.ListFiles)
		protectedRoutes.GET("/download/:id", fileHandler.DownloadFile)

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
