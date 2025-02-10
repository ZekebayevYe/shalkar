package feedback

import (
	"time"
)

type ChatMessage struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	SenderID   int       `json:"sender_id" gorm:"index"`
	ReceiverID int       `json:"receiver_id" gorm:"index"`
	IsAdmin    bool      `json:"is_admin"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
	ChatRoom   string    `json:"chat_room"`
}
