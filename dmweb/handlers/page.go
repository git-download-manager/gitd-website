package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func PageDetail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// slug
		slug := c.Params("slug")

		title := "Page"
		template := "page"

		switch slug {
		case "privacy-policy":
			template = slug
			title = "Privacy Policy"
		case "contact-me":
			template = slug
			title = "Contact Me"
		case "faq":
			template = slug
			title = "FAQ"
		default:
			return c.Next()
		}

		return c.Render("templates/"+template, fiber.Map{
			"Title": title,
		})
	}
}
