package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/crypto/jwt"
)

func Authenticate(c *fiber.Ctx) error {

	token := string(c.Request().Header.Peek("Authentication"))

	err := jwt.ValidadeToken(token)
	if err != nil {
		return err
	}

	return c.Next()
}
