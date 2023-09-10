package chat_http

import (
	"goChat/src/chat/chat_domain"
)

func ChatToDto(input chat_domain.Chat) ChatDto {
	return ChatDto{
		ID:          input.ID,
		Name:        input.Name,
		LastMessage: input.LastMessage,
		Subscribers: input.Subscribers,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
	}
}

func ChatFromCreateDto(input ChatDto) chat_domain.Chat {
	return chat_domain.Chat{
		Name:        input.Name,
		Subscribers: input.Subscribers,
	}
}

func ChatFromUpdateDto(input ChatDto) chat_domain.Chat {
	return chat_domain.Chat{
		ID:          input.ID,
		LastMessage: input.LastMessage,
	}
}
