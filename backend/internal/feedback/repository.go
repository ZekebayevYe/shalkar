package feedback

import (
	"log"

	"gorm.io/gorm"
)

type ChatRepository interface {
	SendMessage(chat *ChatMessage) error
	GetChatHistory(chatRoom string) ([]ChatMessage, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) SendMessage(chat *ChatMessage) error {
	log.Println("📌 Попытка сохранить сообщение:", chat)

	err := r.db.Create(chat).Error
	if err != nil {
		log.Println("❌ Ошибка сохранения сообщения в БД:", err)
	}
	return err
}

func (r *chatRepository) GetChatHistory(chatRoom string) ([]ChatMessage, error) {
	var messages []ChatMessage
	err := r.db.Where("chat_room = ?", chatRoom).Order("timestamp ASC").Find(&messages).Error
	return messages, err
}
