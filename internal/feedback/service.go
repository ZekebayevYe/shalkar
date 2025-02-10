package feedback

import (
	"fmt"
	"time"
)

type ChatService struct {
	Repo *ChatRepository
}

func NewChatService(repo *ChatRepository) *ChatService {
	return &ChatService{Repo: repo}
}

func (s *ChatService) SendMessage(chat ChatMessage) error {
	chat.Timestamp = time.Now()
	err := s.Repo.SendMessage(chat)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}

func (s *ChatService) GetChatHistory(chatRoom string) ([]ChatMessage, error) {
	messages, err := s.Repo.GetChatHistory(chatRoom)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch chat history: %v", err)
	}
	return messages, nil
}
