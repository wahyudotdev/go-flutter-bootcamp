package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-flutter-bootcamp/config"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/router"
	"log"
)

func main() {
	db := config.Db
	err := db.AutoMigrate(models.NoteEntity{}, models.UserEntity{})
	if err != nil {
		log.Fatalln(err)
	}
	app := fiber.New()
	app.Use(config.Cors)
	app.Use(logger.New())
	app.Static("/public", config.PublicDir)
	api := app.Group("/api")
	{
		router.User(api, db)
		router.Note(api, db)
	}
	log.Fatalln(app.Listen(config.Port))
}
