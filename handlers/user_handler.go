package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/models/failure"
	"go-flutter-bootcamp/repository/user_repository"
	"log"
	"mime/multipart"
	"strconv"
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

func (r UserHandler) GetToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenExp, err := strconv.Atoi(c.Query("expired", "0"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InvalidInput,
			})
		}
		token, err := helper.SignJwt("", tokenExp)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Message: "Invalid token",
				Error:   failure.InvalidToken,
			})
		}
		return c.JSON(models.GeneralResponse{
			Message: "success",
			Token:   *token,
		})
	}
}

func (r UserHandler) RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {

		auth := c.Get("Authorization")
		claims, err := helper.ParseJwt(auth)
		if err != nil && err.Error() != failure.ExpiredToken {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Message: "Invalid Token",
				Error:   failure.InvalidToken,
			})
		}
		id := claims["id"].(string)
		if id == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Message: "Invalid Token",
				Error:   failure.InvalidToken,
			})
		}
		token, err := helper.SignJwt(id, 5)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Message: "Invalid Token",
				Error:   failure.InvalidToken,
			})
		}
		return c.JSON(models.GeneralResponse{
			Message: "success",
			Token:   *token,
		})
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
		tokenExp, err := strconv.Atoi(c.Query("expired", "0"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InvalidInput,
			})
		}
		token, err := helper.SignJwt(data.Id, tokenExp)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Message: "Login failed",
				Error:   failure.InvalidCredential,
			})
		}
		return c.JSON(models.GeneralResponse{
			Message: "success",
			Data:    data,
			Token:   *token,
		})
	}
}

func (r UserHandler) GetProfile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := helper.GetUserFromLocals(c)
		data, err := r.repo.Detail(context.TODO(), user.Id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InternalServerError,
			})
		}
		return c.JSON(models.GeneralResponse{
			Message: "success",
			Data:    data,
		})
	}
}

func (r UserHandler) UpdateProfile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var file *multipart.File
		form, _ := c.FormFile("photo")
		if form != nil {
			f, err := form.Open()
			defer func(f multipart.File) {
				err := f.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(f)
			file = &f
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
					Message: err.Error(),
					Error:   failure.InvalidInput,
				})
			}
		}
		reqBody, err := helper.ParseAndValidateBody[models.UpdateProfileRequest](c)
		user := helper.GetUserFromLocals(c)
		data, err := r.repo.Update(context.TODO(), user.Id, file, reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InvalidInput,
			})
		}
		return c.JSON(models.GeneralResponse{
			Message: "success",
			Data:    data,
		})
	}
}
