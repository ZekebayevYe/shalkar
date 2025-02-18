package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"UMS/config"
	"UMS/internal/auth"
	"UMS/internal/document"
	"UMS/internal/expenses"
	"UMS/internal/feedback"
	"UMS/internal/issue"
	"UMS/internal/news"
	"UMS/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🔄 Подключение к базе данных...")
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

	feedRepo := feedback.NewFeedbackRepository(db)
	feedService := feedback.NewFeedbackService(feedRepo)
	feedHandler := feedback.NewFeedbackHandler(feedService)

	r := gin.Default()

	issue.RegisterRoutes(r, db)
	news.NewsRegisterRoutes(r)

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

		protectedRoutes.POST("/expenses", expHandler.CalculateExpense)
		protectedRoutes.GET("/expenses/pay", expHandler.ShowPaymentPage)
		protectedRoutes.GET("/expenses/history", expHandler.GetExpenseHistory)
		protectedRoutes.GET("/expenses/:expense_id", expHandler.GetExpenseDetails)

		protectedRoutes.POST("/feedback", feedHandler.SubmitFeedback)
		protectedRoutes.GET("/feedback", feedHandler.GetUserFeedback)

		adminRoutes := protectedRoutes.Group("/")
		adminRoutes.Use(middleware.AdminMiddleware())
		{
			adminRoutes.POST("/upload", fileHandler.UploadFile)
			adminRoutes.DELETE("/files/:id", fileHandler.DeleteFile)
		}

	}

	r.Static("/frontend", "./frontend")

	r.GET("/register", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/frontend/pages/index.html")
	})
	r.POST("/expenses", func(c *gin.Context) {
		c.File("./frontend/pages/expenses.html")
	})

	r.POST("/feedback", func(c *gin.Context) {
		c.File("./frontend/pages/feedback.html")
	})
	r.GET("/api/me", middleware.AdminMiddleware(), func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user_id": userID})
	})

	r.Use(middleware.AuthMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf(" Сервер запущен на http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf(" Ошибка запуска сервера: %v", err)
	}
}
