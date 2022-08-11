package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/application/services"
	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
)

type TransferController struct {
	transferServices *services.TransferServices
}

func NewTransferController(transferServices *services.TransferServices) *TransferController {
	return &TransferController{
		transferServices,
	}
}

func (controller *TransferController) AccountToAccountTransfer(ctx *fiber.Ctx) error {
	transfer := new(entities.Transfer)

	err := ctx.BodyParser(&transfer)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	cpfUser := ctx.Locals("user").(string)
	transferRsp, errors := controller.transferServices.TransferAccountToAccount(cpfUser, transfer)
	if errors != nil {
		return ctx.Status(errors.StatusCode).JSON(errors)
	}

	return ctx.JSON(transferRsp)
}

func (controller *TransferController) GetAccountTransfer(ctx *fiber.Ctx) error {

	cpfUser := ctx.Locals("user").(string)

	transferRsp, errors := controller.transferServices.GetAccountTransfer(cpfUser)
	if errors != nil {
		return ctx.Status(errors.StatusCode).JSON(errors)
	}

	return ctx.JSON(transferRsp)
}
