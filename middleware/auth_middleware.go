package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pageton/authify/services"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"ok": false, "error": "No authorization token provided"})
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"ok": false, "error": "Invalid token format"})
	}

	_, err := services.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"ok": false, "error": "Invalid or expired token"})
	}

	return c.Next()
}
