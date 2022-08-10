package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/http/fiber/controllers"
)

const TRANSFERS_GET = "/transfers"
const TRANSFERS_POST = "/transfers"

func SetupTransferRouter(router fiber.Router, controller *controllers.TransferController) {
	router.Get(TRANSFERS_GET, controller.GetAccountTransfer)
	router.Post(TRANSFERS_POST, controller.AccountToAccountTransfer)
}
