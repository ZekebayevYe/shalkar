package auth

import (
	"database/sql"
)

type AuthRepository interface {
	ValidateCredentials(userID, password string) (bool, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (rep *authRepository) ValidateCredentials(userID, password string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE user_id = $1 AND password = $2`
	var count int
	err := rep.db.QueryRow(query, userID, password).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
