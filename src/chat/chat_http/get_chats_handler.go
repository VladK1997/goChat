package chat_http

import (
	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"goChat/src/chat/chat_domain"
)

func GetChatsHandler(chatService chat_domain.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		query := &ChatQuery{}
		if err := ctx.Bind(structs.Map(query)); err != nil {
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
