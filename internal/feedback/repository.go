package feedback

import (
	"database/sql"
)

type FeedbackRepository struct {
	DB *sql.DB
}

func (r *FeedbackRepository) SubmitFeedback(feedback Feedback) error {
	query := `INSERT INTO feedback (user_id, category, message, status, is_anonymous) 
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, feedback.UserID, feedback.Category, feedback.Message, feedback.Status, feedback.IsAnonymous)
	return err
}

func (r *FeedbackRepository) GetAllFeedback() ([]Feedback, error) {
	query := `SELECT id, user_id, category, message, status, is_anonymous, created_at FROM feedback`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []Feedback
	for rows.Next() {
		var f Feedback
		if err := rows.Scan(&f.ID, &f.UserID, &f.Category, &f.Message, &f.Status, &f.IsAnonymous, &f.CreatedAt); err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, f)
	}
	return feedbacks, nil
}

func (r *FeedbackRepository) UpdateFeedbackStatus(id int, status string) error {
	query := `UPDATE feedback SET status = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, status, id)
	return err
}

func (r *FeedbackRepository) AddComment(comment FeedbackComment) error {
	query := `INSERT INTO feedback_comments (feedback_id, user_id, comment) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, comment.FeedbackID, comment.UserID, comment.Comment)
	return err
}

func (r *FeedbackRepository) GetComments(feedbackID int) ([]FeedbackComment, error) {
	query := `SELECT id, feedback_id, user_id, comment, created_at FROM feedback_comments WHERE feedback_id = $1`
	rows, err := r.DB.Query(query, feedbackID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []FeedbackComment
	for rows.Next() {
		var c FeedbackComment
		if err := rows.Scan(&c.ID, &c.FeedbackID, &c.UserID, &c.Comment, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
