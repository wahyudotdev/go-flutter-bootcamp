package router

import (
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/config"
	"go-flutter-bootcamp/handlers"
	"go-flutter-bootcamp/middlewares"
	"go-flutter-bootcamp/repository/note_repository"
)

func Note(router fiber.Router) {

	repo := note_repository.New(config.Db)

	handler := handlers.NewNoteHandler(repo)

	group := router.Group("/note")
	{
		group.Post("/", middlewares.Auth(), handler.Create())
		group.Get("/", middlewares.Auth(), handler.GetAll())
		group.Delete("/:id", middlewares.Auth(), handler.Delete())
		group.Patch("/:id", middlewares.Auth(), handler.Update())
	}
}
