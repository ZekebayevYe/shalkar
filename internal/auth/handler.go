package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handl *AuthHandler) Login(cont *gin.Context) {
	var loginRequest LoginRequest
	if err := cont.ShouldBindJSON(&loginRequest); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := handl.service.Login(loginRequest.UserID, loginRequest.Password)
	if err != nil {
		cont.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cont.JSON(http.StatusOK, gin.H{"token": token})
}
