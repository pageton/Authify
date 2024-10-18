package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pageton/authify/services"
)

func RedirectHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")

	if token == "" {
		fmt.Println("No token found")
		return c.Redirect("/auth/login", fiber.StatusFound)
	}

	_, err := services.ValidateToken(token)
	if err != nil {
		fmt.Println("Invalid or expired token:", err)
		return c.Redirect("/auth/login", fiber.StatusFound)
	}

	return c.SendFile("./static/index.html")

}
