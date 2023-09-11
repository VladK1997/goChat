package chat

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"goChat/src/chat/chat_database"
	"goChat/src/chat/chat_domain"
	"goChat/src/chat/chat_http"
)

func ChatRoute(route string, app *fiber.App, db *sql.DB) fiber.Router {
	chatRepo := chat_database.NewChatRepoImpl(db)
	chatService := chat_domain.NewChatServiceImpl(chatRepo)
	group := app.Group(route)
	group.Get("", chat_http.GetChatsHandler(chatService))
	group.Post("", chat_http.CreateChatHandler(chatService))
	group.Put("", chat_http.UpdateChatHandler(chatService))
	group.Delete("/:id", chat_http.DeleteChatHandler(chatService))
	return group
}
