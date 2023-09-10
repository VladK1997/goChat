package message_domain

import "time"

type Message struct {
	ID        uint64
	ChatID    uint64
	Text      string
	CreatedBy uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
