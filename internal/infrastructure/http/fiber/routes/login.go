package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/http/fiber/controllers"
)

const LOGIN_POST = "/login"

func SetupLoginRouter(router fiber.Router, controller *controllers.LoginController) {
	router.Post(LOGIN_POST, controller.AuthenticateUser)
}
