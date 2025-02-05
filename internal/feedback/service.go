package feedback

type Service struct {
	Repo *FeedbackRepository
}

func (s *Service) SubmitFeedbackService(feedback Feedback) error {
	return s.Repo.SubmitFeedback(feedback)
}

func (s *Service) GetAllFeedbackService() ([]Feedback, error) {
	return s.Repo.GetAllFeedback()
}

func (s *Service) UpdateFeedbackStatusService(id int, status string) error {
	return s.Repo.UpdateFeedbackStatus(id, status)
}

func (s *Service) AddCommentService(comment FeedbackComment) error {
	return s.Repo.AddComment(comment)
}

func (s *Service) GetCommentsService(feedbackID int) ([]FeedbackComment, error) {
	return s.Repo.GetComments(feedbackID)
}
