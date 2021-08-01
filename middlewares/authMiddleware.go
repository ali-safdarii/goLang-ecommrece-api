package middlewares

import (
	"alisafdarirepo/util"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func AuthRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized",
			})
		},
		SigningKey: []byte(util.SecretKey),
	})
}