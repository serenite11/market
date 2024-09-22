package failure

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func HandleHttpError(log *zap.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var error *LogicError
		if errors.As(err, error) {
			if error.Unwrap() != nil {
				log.Error("wrapped error", zap.Error(error.Unwrap()))
			}
			return ctx.Status(error.HttpCode).JSON(error.Desc)
		}
		log.Error("internal_server_error", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
}
