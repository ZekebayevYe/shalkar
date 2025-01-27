package services

import (
	"database/sql"
	"log"

	"github.com/NameSurname/Assignment1/GOLANGFINAL/models"
)

type IssueService struct {
	DB *sql.DB
}

func (s *IssueService) CreateIssue(issue models.Issue) (int, error) {
	query := `INSERT INTO Issue (user_id, title, description, status) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := s.DB.QueryRow(query, issue.UserID, issue.Title, issue.Description, issue.Status).Scan(&id)
	if err != nil {
		log.Printf("Failed to insert issue: %v", err) // Лог ошибки
		return 0, err
	}
	return id, nil
}

// Получение всех Issue
func (s *IssueService) GetAllIssues() ([]models.Issue, error) {
	query := `SELECT * FROM issue`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []models.Issue
	for rows.Next() {
		var issue models.Issue
		if err := rows.Scan(&issue.ID, &issue.UserID, &issue.Title, &issue.Description, &issue.Status); err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}
	return issues, nil
}
