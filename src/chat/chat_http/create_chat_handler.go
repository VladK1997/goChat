package chat_http

import (
	"github.com/gofiber/fiber/v2"
	"goChat/src/chat/chat_domain"
)

func CreateTodoHandler(chatService chat_domain.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		chatDto := ChatDto{}

		if err := ctx.Bind(&chatDto); err != nil {
			return err
		}

		chat, err := chatService.CreateChat(ChatFromCreateDto(chatDto))
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(ChatToDto(*chat))
	}
}
