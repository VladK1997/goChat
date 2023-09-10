package chat_http

import (
	"github.com/gofiber/fiber/v2"
	"goChat/src/chat/chat_domain"
)

func UpdateChatHandler(chatService chat_domain.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		chat := &chat_domain.Chat{}

		if err := ctx.Bind(chat); err != nil {
			return err
		}

		chat, err := chatService.UpdateChat(*chat)
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(chat)
	}
}
