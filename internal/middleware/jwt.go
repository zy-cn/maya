package middleware

import (
	"maya/global"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtProtected(whitePaths []string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(global.Config.Jwt.Secret),
		ErrorHandler: jwtError,
		Filter: func(c *fiber.Ctx) bool {
			return true
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
