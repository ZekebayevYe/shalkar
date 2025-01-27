package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api")
	{
		RegisterIssueRoutes(api, db)
	}
}
