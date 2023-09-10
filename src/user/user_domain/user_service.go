package user_domain

import "goChat/src/chat/chat_domain"

type UserService interface {
	GetUser(id uint64) (*chat_domain.User, error)
	GetUsers() ([]chat_domain.User, error)
	CreateUser(chat_domain.User) (*chat_domain.User, error)
	UpdateUser(chat_domain.User) (*chat_domain.User, error)
	DeleteUser(id uint64) error
}
