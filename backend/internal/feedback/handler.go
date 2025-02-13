package feedback

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	service ChatService
}

func NewChatHandler(service ChatService) *ChatHandler {
	return &ChatHandler{service: service}
}

func (h *ChatHandler) SendMessageHandler(c *gin.Context) {
	var input struct {
		ReceiverID int    `json:"receiver_id"`
		Message    string `json:"message"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	senderID := c.GetInt("user_id")
	isAdmin := c.GetString("role") == "admin"
	chatRoom := generateChatRoomID(senderID, input.ReceiverID, isAdmin)

	err := h.service.SendMessage(senderID, input.ReceiverID, isAdmin, chatRoom, input.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}

func (h *ChatHandler) GetChatHistoryHandler(c *gin.Context) {
	chatRoom := c.Query("chat_room")
	if chatRoom == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing chat_room parameter"})
		return
	}

	messages, err := h.service.GetChatHistory(chatRoom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch chat history"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func generateChatRoomID(userID, receiverID int, isAdmin bool) string {
	if isAdmin {
		return "admin-" + strconv.Itoa(receiverID)
	}
	return "admin-" + strconv.Itoa(userID)
}
