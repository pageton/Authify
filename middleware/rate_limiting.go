package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pageton/authify/config"
)

var (
	cfg       *config.Config
	rateLimit int
	interval  = time.Second
	visitors  = make(map[string]int)
	mu        sync.Mutex
)

func Init() error {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		return err
	}
	rateLimit = cfg.LIMIT
	return nil
}

func RateLimitMiddleware(c *fiber.Ctx) error {
	ip := c.IP()

	mu.Lock()

	defer mu.Unlock()

	visitors[ip]++

	if visitors[ip] > rateLimit {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Too many requests. Please try again later."})
	}

	go resetVisitorCount(ip)

	return c.Next()
}

func resetVisitorCount(ip string) {
	time.Sleep(interval)
	mu.Lock()
	defer mu.Unlock()
	visitors[ip]--
}
