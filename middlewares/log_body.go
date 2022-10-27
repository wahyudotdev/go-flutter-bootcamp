package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func LogBody() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("====Headers====")
		for _, r := range c.GetReqHeaders() {
			log.Println(r)
		}
		log.Println("====BODY=====")
		log.Println(string(c.Body()))
		return c.Next()
	}
}
