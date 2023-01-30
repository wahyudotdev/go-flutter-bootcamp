package router

import (
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/config"
	"go-flutter-bootcamp/handlers"
	"go-flutter-bootcamp/middlewares"
	"go-flutter-bootcamp/repository/user_repository"
	"gorm.io/gorm"
)

func User(router fiber.Router, db *gorm.DB) {
	repo := user_repository.New(db)
	handler := handlers.NewUserHandler(repo)
	router.Get("/user/refresh-token", handler.RefreshToken())
	router.Get("/user/profile", middlewares.Auth(), handler.GetProfile())
	router.Patch("/user/profile", middlewares.Auth(), handler.UpdateProfile())
	group := router.Group("/user")
	if config.ApiSecret != "" {
		group.Use(middlewares.ApiToken(config.ApiSecret))
		group.Get("/get-token", handler.GetToken())
	}
	{
		group.Post("/", handler.Create())
		group.Post("/login", handler.Login())
	}
}
