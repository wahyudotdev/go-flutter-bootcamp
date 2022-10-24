package helper

import (
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/models"
)

var GetUserFromLocals = func(ctx *fiber.Ctx) *models.UserEntity {
	return ctx.Locals("user").(*models.UserEntity)
}
