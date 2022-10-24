package router

import (
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/config"
	"go-flutter-bootcamp/handlers"
	"go-flutter-bootcamp/repository/user_repository"
)

func User(router fiber.Router) {
	repo := user_repository.New(config.Db)
	handler := handlers.NewUserHandler(repo)
	group := router.Group("/user")
	{
		group.Post("/", handler.Create())
		group.Post("/login", handler.Login())
	}
}
