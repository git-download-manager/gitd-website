package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Home() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title": "Gitd Download Manager - gitd export",
		})
	}
}
