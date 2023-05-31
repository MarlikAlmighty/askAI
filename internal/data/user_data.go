package data

import (
	"github.com/sashabaranov/go-openai"
	"sync"
)

// Pusher for assertion
type Pusher interface {
	Set(userID, msg string)
	Clear()
}

// UserData save users in memory
type UserData struct {
	ID  map[int64][]openai.ChatCompletionMessage
	mux sync.Mutex
}

// NewData simple constructor
func NewData() *UserData {
	return &UserData{
		ID: make(map[int64][]openai.ChatCompletionMessage),
	}
}

// Set added data to map
func (r *UserData) Set(userID int64, msg string) {
	r.mux.Lock()
	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})
	r.ID[userID] = messages
	r.mux.Unlock()
}

// Clear map
func (r *UserData) Clear() {
	r.mux.Lock()
	if len(r.ID) > 0 {
		r.ID = make(map[int64][]openai.ChatCompletionMessage)
	}
	r.mux.Unlock()
}
