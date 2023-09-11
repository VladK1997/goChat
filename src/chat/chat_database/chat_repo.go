package chat_database

import (
	"database/sql"
	"fmt"
	"goChat/src/chat/chat_domain"
	"log"
	"time"
)

const tableName = "chats"

type ChatRepoImpl struct {
	db *sql.DB
}

func NewChatRepoImpl(db *sql.DB) *ChatRepoImpl {
	return &ChatRepoImpl{
		db: db,
	}
}

func (c ChatRepoImpl) GetChat(id uint64) (*chat_domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepoImpl) GetChats(filter *chat_domain.ChatFilter) ([]chat_domain.Chat, error) {
	//TODO implement me
	var chats []chat_domain.Chat

	if filter == nil {
		filter = &chat_domain.ChatFilter{}
	}

	rows, _ := c.db.Query("SELECT * FROM chats")
	/*if filter.Limit != nil {
		res = res.Limit(*filter.Limit)
	}

	if filter.Offset != nil {
		res = res.Offset(*filter.Offset)
	}

	if filter.Search != "" {
		res = res.Where("task LIKE ?", filter.Search)
	}

	if filter.Sort != "" {
		res = res.Order(filter.Sort)
	}
	*/
	/*if err := res.Find(&chats).Error; err != nil {
		return nil, err
	}*/
	defer rows.Close()

	for rows.Next() {
		var (
			id           uint64
			name         string
			last_message uint64
			subscribers  []uint8
			created_at   time.Time
			updated_at   time.Time
		)
		if err := rows.Scan(
			&id,
			&name,
			&last_message,
			&subscribers,
			&created_at,
			&updated_at,
		); err != nil {
			log.Fatal(err)
		}
		row := chat_domain.Chat{
			ID:          id,
			Name:        name,
			Subscribers: subscribers,
			LastMessage: last_message,
			CreatedAt:   created_at,
			UpdatedAt:   updated_at,
		}
		chats = append(chats, row)
		fmt.Printf("%v", rows)
	}

	return chats, nil
}

func (c ChatRepoImpl) CreateChat(chat chat_domain.Chat) (*chat_domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepoImpl) UpdateChat(chat chat_domain.Chat) (*chat_domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepoImpl) DeleteChat(id uint64) error {
	//TODO implement me
	panic("implement me")
}
