package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func UserAgent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// for google pagespeed insight high score
		c.Bind(fiber.Map{
			"ua": c.Get("User-Agent", ""),
		})

		return c.Next()
	}
}
