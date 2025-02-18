package feedback

type FeedbackService interface {
	SubmitFeedback(userID uint, category string, rating int, comment string) (*Feedback, error)
	GetUserFeedback(userID uint) ([]Feedback, error)
}

type feedbackService struct {
	repo FeedbackRepository
}

func NewFeedbackService(repo FeedbackRepository) FeedbackService {
	return &feedbackService{repo: repo}
}

func (s *feedbackService) SubmitFeedback(userID uint, category string, rating int, comment string) (*Feedback, error) {
	feedback := &Feedback{
		UserID:   userID,
		Category: category,
		Rating:   rating,
		Comment:  comment,
	}

	err := s.repo.Save(feedback)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func (s *feedbackService) GetUserFeedback(userID uint) ([]Feedback, error) {
	return s.repo.GetByUserID(userID)
}
