package feedback

import (
	"time"
)

type ChatService interface {
	SendMessage(senderID, receiverID int, isAdmin bool, chatRoom, message string) error
	GetChatHistory(chatRoom string) ([]ChatMessage, error)
}

type chatService struct {
	repo ChatRepository
}

func NewChatService(repo ChatRepository) ChatService {
	return &chatService{repo: repo}
}

func (s *chatService) SendMessage(senderID, receiverID int, isAdmin bool, chatRoom, message string) error {
	chat := ChatMessage{
		SenderID:   senderID,
		ReceiverID: receiverID,
		IsAdmin:    isAdmin,
		ChatRoom:   chatRoom,
		Message:    message,
		Timestamp:  time.Now(),
	}

	return s.repo.SendMessage(&chat)
}

func (s *chatService) GetChatHistory(chatRoom string) ([]ChatMessage, error) {
	return s.repo.GetChatHistory(chatRoom)
}
