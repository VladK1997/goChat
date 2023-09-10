package main

import (
	"github.com/gofiber/fiber/v2"
	"goChat/src/chat"
)

func main() {
	app := fiber.New()
	chat.ChatRoute(app, "/chat")
}
