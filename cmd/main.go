package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.com/desafio-stone/account-bff/internal/application/services"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/http/fiber/controllers"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/http/fiber/middleware"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/http/fiber/routes"
)

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
	})
}

func main() {
	godotenv.Load()
	mysql.InitDBRepository()

	app := fiber.New()
	app.Get("/healthCheck", HealthCheck)

	api := app.Group("/api", middleware.Authenticate)

	accountService := services.NewAccountServices()
	accountController := controllers.NewAccountController(accountService)
	routes.SetupAccountRouter(api, accountController)

	transferService := services.NewTransferServices()
	transferController := controllers.NewTransferController(transferService)
	routes.SetupTransferRouter(api, transferController)

	loginService := services.NewLoginServices()
	loginController := controllers.NewLoginController(loginService)
	routes.SetupLoginRouter(app, loginController)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "5000"
	}

	err := app.Listen(":" + httpPort)
	if err != nil {
		panic(err)
	}
}
