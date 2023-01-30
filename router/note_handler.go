package router

import (
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/handlers"
	"go-flutter-bootcamp/middlewares"
	"go-flutter-bootcamp/repository/note_repository"
	"gorm.io/gorm"
)

func Note(router fiber.Router, db *gorm.DB) {

	repo := note_repository.New(db)

	handler := handlers.NewNoteHandler(repo)

	group := router.Group("/note")
	{
		group.Post("/", middlewares.Auth(), handler.Create())
		group.Get("/", middlewares.Auth(), handler.GetAll())
		group.Delete("/:id", middlewares.Auth(), handler.Delete())
		group.Patch("/:id", middlewares.Auth(), handler.Update())
	}
}
