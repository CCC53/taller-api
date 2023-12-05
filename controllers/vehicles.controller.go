package controllers

import (
	"taller-api/models"
	"taller-api/services"

	"github.com/gofiber/fiber/v2"
)

func FindVehicles(ctx *fiber.Ctx) error {
	pageSize := ctx.QueryInt("pageSize", 5)
	page := ctx.QueryInt("page", 1)
	vehicles, totalCount := services.ListVehicles(pageSize, page)
	return ctx.JSON(&fiber.Map{
		"data":       vehicles,
		"totalCount": totalCount,
	})
}

func FindVehiclesSelect(ctx *fiber.Ctx) error {
	return ctx.JSON(&fiber.Map{
		"vehiclesSelect": services.ListVehiclesSelect(),
	})
}

func FindVehicle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	vehicleDB, err := services.GetVehicleByID(id)
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"vehicle": vehicleDB,
	})
}

func CreateVehicle(ctx *fiber.Ctx) error {
	var vechile models.Vehicle
	ctx.BodyParser(&vechile)
	vehicleDB, err := services.CreateVehicle(vechile)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"vehicle": vehicleDB,
	})
}

func UpdateVehicle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var vehicle models.Vehicle
	ctx.BodyParser(&vehicle)
	vehicleDB, err := services.UpdateVehicle(id, vehicle)
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"vehicle": vehicleDB,
	})
}

func DeleteVehicle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	deleted, err := services.DeleteVehicle(id)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"deleted": deleted,
	})
}
