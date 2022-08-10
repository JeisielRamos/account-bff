package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/http/fiber/controllers"
)

const ACCOUNT_GET = "/accounts"
const ACCOUNT_POST = "/accounts"
const ACCOUNT_GET_BALANCE = "/accounts/:account_id/balance"

func SetupAccountRouter(router fiber.Router, controller *controllers.AccountController) {
	router.Get(ACCOUNT_GET, controller.GetAllAccount)
	router.Get(ACCOUNT_GET_BALANCE, controller.GetAccountBalance)
	router.Post(ACCOUNT_POST, controller.CreateAccount)
}
