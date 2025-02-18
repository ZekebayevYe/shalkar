package feedback

import (
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Save(feedback *Feedback) error
	GetByUserID(userID uint) ([]Feedback, error)
}

type feedbackRepo struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepo{db: db}
}

func (r *feedbackRepo) Save(feedback *Feedback) error {
	return r.db.Create(feedback).Error
}

func (r *feedbackRepo) GetByUserID(userID uint) ([]Feedback, error) {
	var feedback []Feedback
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&feedback).Error
	return feedback, err
}
