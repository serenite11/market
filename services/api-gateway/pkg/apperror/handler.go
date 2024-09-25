package apperror

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func HandleHTTPError(logger *zap.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var customErr Error
		ok := errors.As(err, &customErr)
		if !ok {
			customErr = Error{
				Description:    "unknown error",
				HTTPStatusCode: fiber.StatusInternalServerError,
			}
		}
		logger.Error("HTTP request failed", zap.String("error", customErr.Error()))
		return ctx.Status(customErr.HTTPStatusCode).JSON(customErr)
	}
}
