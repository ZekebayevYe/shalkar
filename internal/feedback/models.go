package feedback

import (
	"time"
)

type ChatMessage struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	IsAdmin    bool      `json:"is_admin"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
	ChatRoom   string    `json:"chat_room"`
}
