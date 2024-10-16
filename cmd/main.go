package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pageton/authify/config"
	db "github.com/pageton/authify/db/model"
	"github.com/pageton/authify/handler"
	"github.com/pageton/authify/middleware"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting IP addresses:", err)
		os.Exit(1)
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "localhost"
}

func main() {
	app := fiber.New()

	app.Use(middleware.CORSMiddleware)

	app.Use(middleware.ErrorHandlingMiddleware)

	app.Use(middleware.RequestLogger)

	if err := middleware.Init(); err != nil {
		log.Fatalf("Could not initialize middleware: %v", err)
	}

	app.Use(middleware.RateLimitMiddleware)

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	database, err := sql.Open("sqlite3", cfg.DatabasePath)

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	defer database.Close()

	queries := db.New(database)

	app.Post("/register", func(c *fiber.Ctx) error {
		return handler.RegisterUser(c, queries)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return handler.LoginUser(c, queries)
	})

	app.Post("/logout", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(string)

		if err := queries.DeleteAuthToken(c.Context(), userID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not log out"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successfully logged out"})
	})

	app.Get("/protected", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "This is a protected route!"})
	})

	ip := getLocalIP()

	fmt.Printf("🚀 Server is running at: http://%s%s\n", ip, cfg.Port)
	log.Fatal(app.Listen(cfg.Port))
}
