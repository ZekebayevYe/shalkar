package feedback

import "time"

type Feedback struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Category    string    `json:"category"`
	Message     string    `json:"message"`
	Status      string    `json:"status"`
	IsAnonymous bool      `json:"is_anonymous"`
	CreatedAt   time.Time `json:"created_at"`
}

type FeedbackComment struct {
	ID         int       `json:"id"`
	FeedbackID int       `json:"feedback_id"`
	UserID     int       `json:"user_id"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}
