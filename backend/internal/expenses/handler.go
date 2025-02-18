package expenses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	service ExpenseService
}

func NewExpenseHandler(service ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

func (h *ExpenseHandler) ShowPaymentPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Use this QR code to make a payment",
		"qr_code": "https://web-development.kz/images/detailed/8/11.png", // Замени на реальную ссылку
	})
}

func (h *ExpenseHandler) GetExpenseHistory(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}

	userIDInt := userID.(int)
	expenses, err := h.service.GetUserExpenses(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"history": expenses})
}

func (h *ExpenseHandler) CalculateExpense(c *gin.Context) {
	var input Expense

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("❌ Ошибка чтения тела запроса:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	fmt.Println("Полученный JSON:", string(body))

	if err := json.Unmarshal(body, &input); err != nil {
		fmt.Println("❌ Ошибка JSON Unmarshal:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	fmt.Println("✅ Успешно распарсенные данные:", input)

	userID, exists := c.Get("user_id")
	if !exists {
		fmt.Println("❌ Ошибка: Отсутствует user_id")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		fmt.Println("❌ Ошибка: Неверный формат user_id")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid User ID"})
		return
	}

	if input.ColdWater < 0 || input.HotWater < 0 || input.Heating < 0 || input.Gas < 0 || input.Electricity < 0 {
		fmt.Println("❌ Ошибка: Значения не могут быть отрицательными")
		c.JSON(http.StatusBadRequest, gin.H{"error": "All values must be positive numbers"})
		return
	}

	result, err := h.service.CalculateAndSave(userIDInt, input)
	if err != nil {
		fmt.Println("❌ Ошибка сохранения расходов:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save expenses"})
		return
	}

	fmt.Println(" Расходы успешно сохранены:", result)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Expense calculated successfully",
		"total":     result.TotalCost,
		"breakdown": result,
	})
}

func (h *ExpenseHandler) GetExpenseDetails(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		return
	}

	userIDInt := userID.(int)
	expenseIDStr := c.Param("expense_id")
	expenseID, err := strconv.Atoi(expenseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	expense, err := h.service.GetExpenseDetails(expenseID, userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expense details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"details": expense})
}
	