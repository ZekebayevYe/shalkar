package expenses

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	service ExpenseService
}

func NewExpenseHandler(service ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

func (h *ExpenseHandler) CalculateExpense(c *gin.Context) {
	var input Expense
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetInt("user_id") // Теперь получаем user_id как int
	if userID == 0 {
		log.Println("❌ Ошибка: user_id не передан или равен 0")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id"})
		return
	}

	log.Println("📌 user_id:", userID)
	log.Println("📌 Входные данные:", input)

	result, err := h.service.CalculateAndSave(userID, input)
	if err != nil {
		log.Println("❌ Ошибка при сохранении:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save expenses"})
		return
	}

	c.JSON(http.StatusOK, result)
}
