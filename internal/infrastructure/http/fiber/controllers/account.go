package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/application/services"
	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
)

type AccountController struct {
	accountService *services.AccountServices
}

func NewAccountController(accountService *services.AccountServices) *AccountController {
	return &AccountController{
		accountService,
	}
}

func (controller *AccountController) Create(ctx *fiber.Ctx) error {
	account := new(entities.Account)

	err := ctx.BodyParser(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	accountCreated, errors := controller.accountService.Create(account)
	if errors != nil {
		return ctx.Status(errors.StatusCode).JSON(errors.Message)
	}

	return ctx.JSON(accountCreated)
}

func (controller *AccountController) GetAll(ctx *fiber.Ctx) error {
	var account entities.Account

	err := ctx.BodyParser(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// err = controller.producer.SendToUpdateAssignmentQueue(account)
	// if err != nil {
	// 	return err
	// }
	return ctx.JSON(nil)
}

func (controller *AccountController) GetAccountBalance(ctx *fiber.Ctx) error {
	var account_id string

	println(account_id)
	// err := ctx.BodyParser(&account)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	// }

	// err = controller.producer.SendToUpdateAssignmentQueue(account)
	// if err != nil {
	// 	return err
	// }

	return ctx.JSON(nil)
}
