package feedback

import (
	"database/sql"
	"fmt"
	"time"
)

type ChatRepository struct {
	DB *sql.DB
}

func (r *ChatRepository) SendMessage(chat ChatMessage) error {
	query := `
		INSERT INTO chat_messages (sender_id, receiver_id, is_admin, message, timestamp, chat_room)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	err := r.DB.QueryRow(query, chat.SenderID, chat.ReceiverID, chat.IsAdmin, chat.Message, time.Now(), chat.ChatRoom).Scan(&chat.ID)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}

func (r *ChatRepository) GetChatHistory(chatRoom string) ([]ChatMessage, error) {
	query := "SELECT id, sender_id, receiver_id, is_admin, message, timestamp FROM chat_messages WHERE chat_room = $1 ORDER BY timestamp ASC"
	rows, err := r.DB.Query(query, chatRoom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []ChatMessage
	for rows.Next() {
		var msg ChatMessage
		if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.IsAdmin, &msg.Message, &msg.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *ChatRepository) GetUserChat(userID int) ([]ChatMessage, error) {
	query := "SELECT id, sender_id, receiver_id, is_admin, message, timestamp, chat_room FROM chat_messages WHERE sender_id = $1 OR receiver_id = $1 ORDER BY timestamp ASC"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []ChatMessage
	for rows.Next() {
		var msg ChatMessage
		if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.IsAdmin, &msg.Message, &msg.Timestamp, &msg.ChatRoom); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
