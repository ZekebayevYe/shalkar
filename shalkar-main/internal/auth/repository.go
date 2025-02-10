package auth

import (
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *AuthRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
