package auth

import (
	"UMS/utils"
	"errors"
	"UMS/internal/models"
)

type AuthService struct {
    repo AuthRepositoryInterface // ✅ Просто интерфейс
}

func NewAuthService(repo AuthRepositoryInterface) *AuthService { // ✅ Без указателя
    return &AuthService{repo: repo}
}


func (s *AuthService) Register(username, password, role string) error {
	if username == "" || password == "" {
		return errors.New("username and password cannot be empty")
	}

	if role == "" {
		role = "user"
	}

	if role != "admin" && role != "user" {
		return errors.New("invalid role")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (*models.User, string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", errors.New("user not found")
	}

	if !utils.CheckPassword(password, user.Password) {
		return nil, "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
