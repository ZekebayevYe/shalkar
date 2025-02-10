package issue

import (
	"log"

	"gorm.io/gorm"
)

type IssueService struct {
	DB *gorm.DB
}

func (s *IssueService) CreateIssue(issue Issue) (int, error) {
	err := s.DB.Create(&issue).Error
	if err != nil {
		log.Printf("Failed to insert issue: %v", err)
		return 0, err
	}
	return issue.ID, nil
}

func (s *IssueService) GetAllIssues() ([]Issue, error) {
	var issues []Issue
	err := s.DB.Find(&issues).Error
	if err != nil {
		return nil, err
	}
	return issues, nil
}
