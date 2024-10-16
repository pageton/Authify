package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger(c *fiber.Ctx) error {
	log.Printf("Request: %s %s", c.Method(), c.Path())

	return c.Next()
}
