package chat_domain

type ChatService interface {
	GetChat(id uint64) (*Chat, error)
	GetChats(filter *ChatFilter) ([]Chat, error)
	CreateChat(Chat) (*Chat, error)
	UpdateChat(Chat) (*Chat, error)
	DeleteChat(id uint64) error
}

type ChatServiceImpl struct {
	chatRepo ChatRepo
}

func NewChatServiceImpl(chatRepo ChatRepo) *ChatServiceImpl {
	return &ChatServiceImpl{
		chatRepo: chatRepo,
	}
}

func (c ChatServiceImpl) GetChat(id uint64) (*Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatServiceImpl) GetChats(filter *ChatFilter) ([]Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatServiceImpl) CreateChat(chat Chat) (*Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatServiceImpl) UpdateChat(chat Chat) (*Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatServiceImpl) DeleteChat(id uint64) error {
	//TODO implement me
	panic("implement me")
}
