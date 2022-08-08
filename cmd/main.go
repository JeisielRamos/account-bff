package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
	})
}

func main() {
	godotenv.Load()

	app := fiber.New()

	app.Get("/helthCheck", healthCheck)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "5000"
	}

	err := app.Listen(":" + httpPort)
	if err != nil {
		panic(err)
	}
}
