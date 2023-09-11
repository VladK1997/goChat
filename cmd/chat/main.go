package main

import (
	"github.com/gofiber/fiber/v2"
	"goChat/postgres"
	"goChat/src/chat"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := postgres.NewPostgresClient(dsn)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	chat.ChatRoute("/chat", app, db)
	log.Fatal(app.Listen(":8080"))
}
