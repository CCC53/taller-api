package controllers

import (
	"taller-api/models"
	"taller-api/services"

	"github.com/gofiber/fiber/v2"
)

func FindServices(ctx *fiber.Ctx) error {
	pageSize := ctx.QueryInt("pageSize", 5)
	page := ctx.QueryInt("page", 1)
	services, totalCount := services.ListServices(pageSize, page)
	return ctx.JSON(&fiber.Map{
		"services":   services,
		"totalCount": totalCount,
	})
}

func FindService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	serviceDB, err := services.GetServicePopulatedByID(id)
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"service": serviceDB,
	})
}

func CreateService(ctx *fiber.Ctx) error {
	var service models.Service
	ctx.BodyParser(&service)
	serviceDB, err := services.CreateService(service)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"service": serviceDB,
	})
}

func UpdateService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var service models.Service
	ctx.BodyParser(&service)
	serviceDB, err := services.UpdateService(id, service)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"service": serviceDB,
	})
}

func DeleteService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	deleted, err := services.DeleteService(id)
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
