package issue

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		RegisterIssueRoutes(api, db)
	}
}
