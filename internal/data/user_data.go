package data

import (
	"github.com/sashabaranov/go-openai"
	"sync"
)

// Pusher for assertion
type Pusher interface {
	Set(userID, msg string)
	Get(userID int64) []openai.ChatCompletionMessage
	Clear()
}

// UserData save users in memory
type UserData struct {
	ID  map[int64][]openai.ChatCompletionMessage
	mux sync.Mutex
}

// New simple constructor
func New() *UserData {
	return &UserData{
		ID: make(map[int64][]openai.ChatCompletionMessage),
	}
}

// Set added data to map
func (data *UserData) Set(userID int64, msg []openai.ChatCompletionMessage) {
	data.mux.Lock()
	mp := data.ID[userID]
	mp = append(mp, msg...)
	data.ID[userID] = mp
	data.mux.Unlock()
}

// Get data from map
func (data *UserData) Get(userID int64) []openai.ChatCompletionMessage {
	data.mux.Lock()
	mp := data.ID[userID]
	data.mux.Unlock()
	return mp
}

// Clear map
func (data *UserData) Clear() {
	data.mux.Lock()
	if len(data.ID) > 0 {
		data.ID = make(map[int64][]openai.ChatCompletionMessage)
	}
	data.mux.Unlock()
}
