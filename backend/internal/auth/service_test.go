package auth

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"UMS/internal/auth/mocks" 

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

