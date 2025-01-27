package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	secretKey string
	repo      AuthRepository
}

func NewAuthService(secretKey string, repo AuthRepository) *AuthService {
	return &AuthService{secretKey: secretKey, repo: repo}
}

func (s *AuthService) Login(userID, password string) (string, error) {
	valid, err := s.repo.ValidateCredentials(userID, password)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", errors.New("Incorrect Data")
	}

	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}
