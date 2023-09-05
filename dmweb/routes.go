package main

import (
	"github.com/git-download-manager/gitd-website/dmweb/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	// Routes with Handlers

	// Robot Txt
	app.Get("/robots.txt", handlers.RobotsTxt()).Name("gitdm.robots.txt")

	// Home
	app.Get("/", handlers.Home()).Name("gitdm.home")

	// Page
	app.Get("/p/:slug", handlers.PageDetail()).Name("gitdm.page.detail")
}
