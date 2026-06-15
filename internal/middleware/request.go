package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		// generate request ID
		reqID := uuid.New().String()
		c.Set("X-Request-ID", reqID)

		// continue request
		err := c.Next()

		// calculate response time
		duration := time.Since(start)

		c.Set("X-Response-Time", duration.String())

		return err
	}
}
