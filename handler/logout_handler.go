package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	db "github.com/pageton/authify/db/model"
)

func LogOutUser(c *fiber.Ctx, queries *db.Queries) error {
	userID := c.FormValue("user_id")

	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "error": "User ID is required"})
	}

	if err := queries.DeleteAuthToken(c.Context(), userID); err != nil {
		log.Println("Error deleting auth token for user:", userID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ok": false, "error": "Could not log out"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true, "message": "Successfully logged out"})
}
