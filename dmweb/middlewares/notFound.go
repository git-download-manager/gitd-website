package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("templates/404", fiber.Map{
			"Title":   "404 - Not Found",
			"Message": "Oopps nothing is here!",
		}, "templates/layouts/error") // HTTP:404
	}
}
