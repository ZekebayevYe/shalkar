package news

import (
	"UMS/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получение всех новостей
func GetNews(c *gin.Context) {
	var news []News
	config.DB.Find(&news)
	c.JSON(http.StatusOK, news)
}

// Добавление новости
func CreateNews(c *gin.Context) {
	var news News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	config.DB.Create(&news)
	c.JSON(http.StatusOK, news)
}

// Лайк/дизлайк новости
func ReactNews(c *gin.Context) {
	id := c.Param("id")
	var news News

	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Новость не найдена"})
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
		news.Likes++
	} else if req.Action == "dislike" {
		news.Dislikes++
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверное действие"})
		return
	}
	config.DB.Save(&news)
	c.JSON(http.StatusOK, news)
}
