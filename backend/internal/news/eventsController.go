package news

import (
	"UMS/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получение всех событий
func GetEvents(c *gin.Context) {
	var events []Events
	config.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

// Добавление события
func CreateEvent(c *gin.Context) {
	var event Events
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	config.DB.Create(&event)
	c.JSON(http.StatusOK, event)
}

// Лайк/дизлайк события
func ReactEvent(c *gin.Context) {
	id := c.Param("id")
	var event Events

	if err := config.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Событие не найдено"})
		return
	}

	var req struct {
		Action string `json:"action"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if req.Action == "like" {
		event.Likes++
	} else if req.Action == "dislike" {
		event.Dislikes++
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверное действие"})
		return
	}
	config.DB.Save(&event)
	c.JSON(http.StatusOK, event)
}
