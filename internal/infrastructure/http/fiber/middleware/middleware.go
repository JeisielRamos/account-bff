package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/crypto/jwt"
)

func Authenticate(c *fiber.Ctx) error {

	token := string(c.Request().Header.Peek("Authentication"))

	cpf, err := jwt.ValidadeToken(token)
	if err != nil {
		return err
	}
	fmt.Println("Auth cpf ", cpf)

	c.Locals("user", cpf)

	return c.Next()
}
