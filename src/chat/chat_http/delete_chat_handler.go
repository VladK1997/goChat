package chat_http

import (
	"github.com/gofiber/fiber/v2"
	"goChat/src/chat/chat_domain"
	"strconv"
)

func DeleteTodoHandler(chatService chat_domain.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		idStr := ctx.Params("id")

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return err
		}

		if err := chatService.DeleteChat(id); err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(nil)
	}
}
