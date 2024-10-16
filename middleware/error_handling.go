package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandlingMiddleware(c *fiber.Ctx) error {
	err := c.Next()

	if err != nil {
		log.Println("An error occurred:", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "An internal server error occurred. Please try again later."})

	}
	return nil
}
