package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/models/failure"
	"log"
	"net/http"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt := c.Get("Authorization")
		claims, err := helper.ParseJwt(jwt)
		if err != nil {
			if err.Error() == failure.ExpiredToken {
				return c.Status(http.StatusUnauthorized).JSON(models.GeneralError{
					Message: "Expired token",
					Error:   failure.ExpiredToken,
				})
			}
			return c.Status(http.StatusUnauthorized).JSON(models.GeneralError{
				Message: "Invalid token",
				Error:   failure.InvalidToken,
			})
		}
		user, err := helper.TypeConverter[models.UserEntity](claims)
		if err != nil {
			log.Println(err)
		}
		c.Locals("user", user)
		return c.Next()
	}
}
