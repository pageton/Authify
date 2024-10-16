package handler

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/pageton/authify/db/model"
	"github.com/pageton/authify/models"
	"github.com/pageton/authify/services"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *fiber.Ctx, queries *db.Queries) error {
	var user models.UserModel

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	userDB, err := queries.GetUser(c.Context(), db.GetUserParams{
		Username: user.Username,
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid username or password"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid username or password"})
	}

	token, err := services.CreateToken(userDB.ID, userDB.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create token"})
	}

	ipAddress := c.IP()
	userAgent := c.Get("User-Agent")

	authID := uuid.New().String()

	authTokenParams := db.CreateAuthTokenParams{
		ID:        authID,
		Userid:    userDB.ID,
		Token:     token,
		Expiresat: sql.NullTime{Time: time.Now().Add(24 * time.Hour), Valid: true},
		Ipaddress: sql.NullString{String: ipAddress, Valid: ipAddress != ""},
		Useragent: sql.NullString{String: userAgent, Valid: userAgent != ""},
	}

	if err := queries.CreateAuthToken(c.Context(), authTokenParams); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
