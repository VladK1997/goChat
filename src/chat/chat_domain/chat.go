package chat_domain

import "time"

type Chat struct {
	ID          uint64
	Name        string
	LastMessage uint64
	Subscribers []uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Message struct {
	ID        uint64
	ChatID    uint64
	Text      string
	CreatedBy uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        uint64
	Username  string
	Avatar    string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCredentials struct {
	ID       uint64 // foreign key on user
	Username string
	Password string // bcrypt
}

type UserChat struct {
	UserID uint64
	ChatID uint64
	// Primary key by both fields
}

type UserMessage struct {
	UserID      uint64
	MessageID   uint64
	IsRead      bool
	IsDelivered bool
}

type Session struct {
	ID        string // TODO use UUID instead of string
	UserID    uint64
	Username  string
	FirstName string
	LastName  string
}
