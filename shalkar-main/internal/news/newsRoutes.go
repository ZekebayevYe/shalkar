package news

import (
	"github.com/gin-gonic/gin"
)

func NewsRegisterRoutes(router *gin.Engine) {
	router.POST("/login", Login)

	router.GET("/news", GetNews)
	router.POST("/news/:id/react", ReactNews)

	router.GET("/events", GetEvents)
	router.POST("/events/:id/react", ReactEvent)

	admin := router.Group("/admin")
	admin.Use(AuthMiddleware("admin"))
	{
		admin.POST("/news", CreateNews)
		admin.POST("/events", CreateEvent)
	}
}
