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

// ✅ РЕГИСТРАЦИЯ (Добавлена проверка на пустой пароль)
func (s *AuthService) Register(username, password, role string) error {
	if username == "" || password == "" {
		return errors.New("username and password cannot be empty")
	}

	// Если роль не указана или пустая, устанавливаем "user" по умолчанию
	if role == "" {
		role = "user"
	}

	// Проверяем, что роль корректная
	if role != "admin" && role != "user" {
		return errors.New("invalid role")
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

// ✅ ЛОГИН (Теперь возвращает `user` + `token`)
func (s *AuthService) Login(username, password string) (*User, string, error) {
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

	return user, token, nil // ✅ Теперь возвращаем user + token
}
