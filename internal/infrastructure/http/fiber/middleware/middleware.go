package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/crypto/jwt"
)

func Authenticate(c *fiber.Ctx) error {

	token := string(c.Request().Header.Peek("Authentication"))

	cpf, err := jwt.ValidadeToken(token)
	if err != nil {
		return err
	}
	c.Locals("user", cpf)

	return c.Next()
}
