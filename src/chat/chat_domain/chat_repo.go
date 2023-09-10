package chat_domain

type ChatRepo interface {
	GetChat(id uint64) (*Chat, error)
	GetChats() ([]Chat, error)
	CreateChat(Chat) (*Chat, error)
	UpdateChat(Chat) (*Chat, error)
	DeleteChat(id uint64) error
}
