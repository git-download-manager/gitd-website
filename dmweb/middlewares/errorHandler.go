package middlewares

import (
	"github.com/git-download-manager/gitd-website/dmweb/loggers"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	loggers.Plain.Error("502", zap.String("err", err.Error()))

	return ctx.Status(fiber.StatusInternalServerError).Render("502", fiber.Map{
		"Title":   "502 - Server Error",
		"Message": "Oppps Something wrong!",
	}, "templates/layouts/error")
}
