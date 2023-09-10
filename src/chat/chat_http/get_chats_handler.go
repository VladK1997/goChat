package chat_http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"goChat/src/chat/chat_domain"
)

func GetTodosHandler(chatService chat_domain.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		query := &TodoQuery{}
		if err := ctx.Bind(query); err != nil {
			return err
		}
		chats, err := chatService.GetChats(ChatFilterFromQuery(query))
		if err != nil {
			return err
		}
		return ctx.Status(fiber.StatusOK).JSON(lo.Map(chats, func(chat chat_domain.Chat, _ int) ChatDto {
			return ChatToDto(chat)
		}))
	}
}
