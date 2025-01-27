package controllers

import (
	"net/http"

	"UMS/models"
	"UMS/services"

	"github.com/gin-gonic/gin"
)

type IssueController struct {
	Service *services.IssueService
}

// Создание Issue
func (c *IssueController) CreateIssue(ctx *gin.Context) {
	var issue models.Issue
	if err := ctx.ShouldBindJSON(&issue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := c.Service.CreateIssue(issue)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create issue"})
		return
	}

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
