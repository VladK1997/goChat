package chat_domain

type MessageService interface {
	GetMessage(id uint64) (*Message, error)
	GetMessages() ([]Message, error)
	CreateMessage(Message) (*Message, error)
	UpdateMessage(Message) (*Message, error)
	DeleteMessage(id uint64) error
}
