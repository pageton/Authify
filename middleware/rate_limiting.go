package middleware

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/pageton/authify/config"
	"go.uber.org/ratelimit"
)

var (
	cfg       *config.Config
	rateLimit int
	rlMap     sync.Map
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

	limiter, _ := rlMap.LoadOrStore(ip, ratelimit.New(rateLimit))

	rl := limiter.(ratelimit.Limiter)

	rl.Take()

	if _, ok := rlMap.Load(ip); !ok {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Too many requests. Please try again later.",
		})
	}

	return c.Next()
}
