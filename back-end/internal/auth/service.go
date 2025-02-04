package auth

import (
	"errors"
	"UMS/utils"
)

type AuthService struct {
	repo *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(username, password, role string) error {
	if role != "admin" && role != "user" {
		return errors.New("неверная роль")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("пользователь не найден")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("неверный пароль")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
