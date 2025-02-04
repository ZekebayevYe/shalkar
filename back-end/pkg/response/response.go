package response

import "github.com/gin-gonic/gin"

func ErrorResponse(cont *gin.Context, code int, message string) {
	cont.JSON(code, gin.H{"error": message})
}

func SuccessResponse(cont *gin.Context, data interface{}) {
	cont.JSON(200, data)
}
