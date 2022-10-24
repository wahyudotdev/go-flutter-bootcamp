package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-flutter-bootcamp/config"
	"go-flutter-bootcamp/router"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(config.Cors)
	app.Use(logger.New())

	api := app.Group("/api")
	{
		router.User(api)
		router.Note(api)
	}
	log.Fatalln(app.Listen(config.Port))
}
