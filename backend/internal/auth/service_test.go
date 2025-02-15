package auth

import (
	"testing"

	"UMS/internal/auth/mocks"
	"UMS/internal/models"
	"UMS/utils"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthRepositoryInterface(ctrl)
	authService := NewAuthService(mockRepo)

	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)

	err := authService.Register("testuser", "password123", "user")
	assert.NoError(t, err)
}
func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthRepositoryInterface(ctrl)
	authService := NewAuthService(mockRepo)

	hashedPassword, _ := utils.HashPassword("password123")

	expectedUser := &models.User{
		ID:       1,
		Username: "testuser",
		Password: hashedPassword, 
		Role:     "user",
	}

	mockRepo.EXPECT().FindByUsername("testuser").Return(expectedUser, nil)

	user, token, err := authService.Login("testuser", "password123")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, token) 

	claims, err := utils.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, claims.UserID)
	assert.Equal(t, expectedUser.Role, claims.Role)
}




