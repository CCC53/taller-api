package controllers

import (
	"taller-api/models"
	"taller-api/services"

	"github.com/gofiber/fiber/v2"
)

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
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"assigned": assigned,
	})
}

func RemoveEmployeeFromService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	removed, err := services.RemoveEmployeeFromService(id)
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

func RemoveSparePartFromService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	removed, err := services.RemoveSparePartFromService(id)
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

func FindMechanicsAviable(ctx *fiber.Ctx) error {
	mechanicsAviable := services.ListMechanicsAviable()
	return ctx.JSON(&fiber.Map{
		"employeesAviable": mechanicsAviable,
	})
}

func FindSaprePartsAviable(ctx *fiber.Ctx) error {
	sparePartsAviable := services.ListSaparePartsAviable()
	return ctx.JSON(&fiber.Map{
		"sparePartsAviable": sparePartsAviable,
	})
}
