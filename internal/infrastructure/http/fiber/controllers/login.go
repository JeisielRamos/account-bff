package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/application/services"
	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
)

type LoginController struct {
	loginService *services.LoginServices
}

func NewLoginController(loginService *services.LoginServices) *LoginController {
	return &LoginController{
		loginService,
	}
}

func (controller *LoginController) AuthenticateUser(ctx *fiber.Ctx) error {
	login := new(entities.Login)

	err := ctx.BodyParser(&login)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tokenRsp, errors := controller.loginService.AuthenticateUser(login)
	if errors != nil {
		return ctx.Status(errors.StatusCode).JSON(errors.Message)
	}

	return ctx.JSON(tokenRsp)
}
