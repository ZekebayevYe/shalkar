package auth

import (
	"UMS/internal/models"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock_auth_repository.go -package=mocks

type AuthRepositoryInterface interface {
	CreateUser(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *AuthRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
