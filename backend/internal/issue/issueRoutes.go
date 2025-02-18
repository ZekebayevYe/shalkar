package issue

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterIssueRoutes(router *gin.RouterGroup, db *gorm.DB) {
	issueService := IssueService{DB: db} // Теперь db — *gorm.DB
	issueController := IssueController{Service: issueService}

	router.POST("/issues", issueController.CreateIssue)
	router.GET("/issues", issueController.GetIssues)
}
