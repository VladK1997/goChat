package chat_http

import "time"

type ChatDto struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	LastMessage uint64    `json:"lastMessage"`
	Subscribers []uint64  `json:"subscribers"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
