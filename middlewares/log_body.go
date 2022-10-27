package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func LogBody() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("====Headers====")
		for k, r := range c.GetReqHeaders() {
			log.Printf("%s : %s", k, r)
		}
		log.Println("====BODY=====")
		log.Println(string(c.Body()))
		return c.Next()
	}
}
