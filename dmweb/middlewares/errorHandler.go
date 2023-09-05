package middlewares

import (
	"github.com/git-download-manager/gitd-website/dmweb/loggers"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	loggers.Plain.Error("internal server error", zap.String("err", err.Error()))

	return ctx.Status(fiber.StatusInternalServerError).Render("502", fiber.Map{
		"Title":   "502 - Server Error",
		"Message": "Oppps Something wrong!",
	}, "layouts/error")
}
