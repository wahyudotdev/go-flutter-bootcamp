package middlewares

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/models/failure"
	"strings"
)

func ApiToken(apiSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Error:   failure.InvalidToken,
				Message: "Invalid Token",
			})
		}
		_, err := helper.ParseJwt(token)
		if err == nil {
			return c.Next()
		}
		decByte, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			return err
		}
		decString := string(decByte)
		data := strings.Split(decString, "|")
		//date, err := time.Parse("2006-01-02", data[0])
		//if err != nil {
		//	return c.Status(fiber.StatusBadRequest).JSON(models.GeneralError{
		//		Error:   failure.InvalidInput,
		//		Message: err.Error(),
		//	})
		//}
		//d1 := date.Format("2006-01-02")
		//d2 := time.Now().Format("2006-01-02")
		//if d1 != d2 {
		//	return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
		//		Error:   failure.InvalidInput,
		//		Message: "Expired token",
		//	})
		//}

		if data[1] != apiSecret {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralError{
				Error:   failure.InvalidToken,
				Message: "Invalid Token",
			})
		}
		return c.Next()
	}
}
