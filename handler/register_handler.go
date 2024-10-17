package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/pageton/authify/db/model"
	"github.com/pageton/authify/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx, queries *db.Queries) error {
	var user models.UserModel
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "error": "invalid input"})
	}

	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "error": err.Error()})
	}

	existingUser, err := queries.GetUser(c.Context(), db.GetUserParams{
		Username: user.Username,
	})

	if err == nil && existingUser.Username != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "error": "username already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ok": false, "error": "authentication failed"})
	}

	userID := uuid.New().String()

	_, err = queries.CreateUser(c.Context(), db.CreateUserParams{
		ID:       userID,
		Username: user.Username,
		Password: string(hashedPassword),
	})

	if err != nil {
		log.Println("Could not register user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ok": false, "error": "could not register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ok":       true,
		"id":       userID,
		"username": user.Username,
	})
}
