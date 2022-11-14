package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/models/failure"
	"go-flutter-bootcamp/repository/note_repository"
)

type NoteHandler struct {
	repo note_repository.Repository
}

func NewNoteHandler(repo note_repository.Repository) NoteHandler {
	return NoteHandler{
		repo: repo,
	}
}

func (r NoteHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqBody, err := helper.ParseAndValidateBody[models.CreateNoteRequest](c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InvalidInput,
			})
		}
		user := helper.GetUserFromLocals(c)
		err = r.repo.Create(context.TODO(), user.Id, reqBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InternalServerError,
			})
		}
		return c.JSON(models.GeneralResponse{Message: "success"})
	}
}

func (r NoteHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := helper.GetUserFromLocals(c)
		data, err := r.repo.GetAll(context.TODO(), user.Id)
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

func (r NoteHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := helper.GetUserFromLocals(c)
		err := r.repo.Delete(context.TODO(), user.Id, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InternalServerError,
			})
		}
		return c.JSON(models.GeneralResponse{Message: "note has been deleted"})
	}
}

func (r NoteHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqBody, err := helper.ParseAndValidateBody[models.UpdateNotes](c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
				Message: "Invalid data",
				Error:   failure.InvalidInput,
			})
		}
		id := c.Params("id")
		user := helper.GetUserFromLocals(c)
		err = r.repo.Update(context.TODO(), user.Id, id, reqBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralError{
				Message: err.Error(),
				Error:   failure.InternalServerError,
			})
		}
		return c.JSON(models.GeneralResponse{Message: "note has been updated"})
	}
}
