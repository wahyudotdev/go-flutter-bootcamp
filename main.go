package main

import (
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/config"
	"go-flutter-bootcamp/router"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(config.Cors)

	api := app.Group("/api")
	{
		router.User(api)
	}

	log.Fatalln(app.Listen(config.Port))
}
