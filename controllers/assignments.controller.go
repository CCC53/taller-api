package controllers

import (
	"taller-api/models"
	"taller-api/services"

	"github.com/gofiber/fiber/v2"
)

func FindItemsAviable(ctx *fiber.Ctx) error {
	table := ctx.Params("table")
	return ctx.JSON(&fiber.Map{
		"data": services.ListItemsAviable(table),
	})
}

func AssignEmployeeToService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var employeeDTO models.AssignEmployeeDTO
	ctx.BodyParser(&employeeDTO)
	assigned, err := services.AssignEmployeeToService(id, employeeDTO.EmployeeID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"assigned": assigned,
	})
}

func AssignSparePartToService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var sparePartDTO models.AssignSparePartDTO
	ctx.BodyParser(&sparePartDTO)
	assigned, err := services.AssignSparePartToService(id, sparePartDTO)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": services.Capitalize(err.Error()),
		})
	}
	return ctx.JSON(&fiber.Map{
		"assigned": assigned,
	})
}

func RemoveItemFromService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	table := ctx.Params("table")
	removed, err := services.RemoveItemFromService(table, id)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"removed": removed,
	})
}
