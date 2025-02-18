package issue

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IssueController struct {
	Service IssueService
}

// Создание Issue
func (c *IssueController) CreateIssue(ctx *gin.Context) {
	var issue Issue

	// Привязываем JSON-данные к структуре
	if err := ctx.ShouldBindJSON(&issue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Проверяем, передан ли user_id в контексте (из middleware)
	userID, exists := ctx.Get("user_id")
	if exists {
		// Если user_id есть в контексте, используем его
		log.Println("✅ user_id найден в контексте:", userID)
		issue.UserID = userID.(int)
	} else {
		// Если в контексте нет user_id, проверяем, есть ли он в JSON
		if issue.UserID == 0 {
			log.Println("❌ Ошибка: user_id отсутствует в контексте и в запросе")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			return
		}
		log.Println("⚠️ Используем user_id из JSON:", issue.UserID)
	}

	// 🚀 Лог перед сохранением
	log.Println("📌 Сохраняем issue:", issue)

	// Сохраняем в сервисе
	id, err := c.Service.CreateIssue(issue)
	if err != nil {
		log.Println("❌ Ошибка сохранения issue:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create issue"})
		return
	}

	log.Println("✅ Issue создан с ID:", id)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Issue created successfully", "id": id})
}

// Получение всех Issue
func (c *IssueController) GetIssues(ctx *gin.Context) {
	issues, err := c.Service.GetAllIssues()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch issues"})
		return
	}

	ctx.JSON(http.StatusOK, issues)
}
