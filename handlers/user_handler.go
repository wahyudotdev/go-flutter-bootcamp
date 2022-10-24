package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/models/failure"
	"go-flutter-bootcamp/repository/user_repository"
)

type UserHandler struct {
	repo user_repository.Repository
}

func NewUserHandler(repo user_repository.Repository) UserHandler {
	return UserHandler{
		repo: repo,
	}
}

func (r UserHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqBody, err := helper.ParseAndValidateBody[models.CreateUserRequest](c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InvalidInput,
			})
		}

		err = r.repo.Create(context.TODO(), reqBody)
		if err != nil {
			if err.Error() == failure.AlreadyExists {
				return c.Status(fiber.StatusConflict).JSON(models.GeneralError{
					Message: err.Error(),
					Error:   failure.AlreadyExists,
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InternalServerError,
			})
		}
		return c.JSON(models.GeneralResponse{Message: "user created"})
	}
}

func (r UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqBody, err := helper.ParseAndValidateBody[models.LoginRequest](c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InvalidInput,
			})
		}
		data, err := r.repo.Login(context.TODO(), reqBody)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Message: "Login failed",
				Error:   failure.InvalidCredential,
			})
		}
		return c.JSON(models.GeneralResponse{
			Message: "success",
			Data:    data,
		})
	}
}
