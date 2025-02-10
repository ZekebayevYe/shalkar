package feedback

import (
	"encoding/json"
	"net/http"
)

type ChatHandler struct {
	Service *ChatService
}

func NewChatHandler(service *ChatService) *ChatHandler {
	return &ChatHandler{Service: service}
}

func (h *ChatHandler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var chat ChatMessage
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	if err := h.Service.SendMessage(chat); err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Message sent successfully"})
}

func (h *ChatHandler) GetChatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	chatRoom := r.URL.Query().Get("chat_room")
	if chatRoom == "" {
		http.Error(w, "Missing chat_room parameter", http.StatusBadRequest)
		return
	}

	messages, err := h.Service.GetChatHistory(chatRoom)
	if err != nil {
		http.Error(w, "Failed to fetch chat history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
