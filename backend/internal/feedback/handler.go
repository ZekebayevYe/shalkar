package feedback

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	service FeedbackService
}

func NewFeedbackHandler(service FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{service: service}
}

func (h *FeedbackHandler) SubmitFeedback(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}

	var input struct {
		Category string `json:"category"`
		Rating   int    `json:"rating"`
		Comment  string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	feedback, err := h.service.SubmitFeedback(uint(userIDInt), input.Category, input.Rating, input.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feedback", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback submitted successfully", "feedback": feedback})
}

func (h *FeedbackHandler) GetUserFeedback(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	feedbacks, err := h.service.GetUserFeedback(uint(userIDInt)) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch feedback", "details": err.Error()})
		return
	}

	if len(feedbacks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No feedback found", "feedback": []interface{}{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"feedback": feedbacks})
}
