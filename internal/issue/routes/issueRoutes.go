package routes

import (
	"database/sql"

	"UMS/controllers"
	"UMS/services"

	"github.com/gin-gonic/gin"
)

func RegisterIssueRoutes(router *gin.RouterGroup, db *sql.DB) {
	issueService := &services.IssueService{DB: db}
	issueController := &controllers.IssueController{Service: issueService}

	router.POST("/issues", issueController.CreateIssue)
	router.GET("/issues", issueController.GetIssues)
}
