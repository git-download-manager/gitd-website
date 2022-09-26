package handlers

import "github.com/gofiber/fiber/v2"

// RobotTxt
// Via: https://developers.google.com/search/docs/advanced/robots/create-robots-txt
func RobotTxt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Allowing all web crawlers access to all content

		return c.SendString("User-agent: * \rAllow: / \r")
	}
}
