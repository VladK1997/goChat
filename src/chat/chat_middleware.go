package chat

import "github.com/gofiber/fiber/v2"

func ChatRoute(app *fiber.App, route string) fiber.Router {
	group := app.Group(route)
	group.Get("/")
	return group
}
