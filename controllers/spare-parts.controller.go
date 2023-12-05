package controllers

import (
	"taller-api/models"
	"taller-api/services"

	"github.com/gofiber/fiber/v2"
)

func FindSpareParts(ctx *fiber.Ctx) error {
	pageSize := ctx.QueryInt("pageSize", 5)
	page := ctx.QueryInt("page", 1)
	spareParts, totalCount := services.ListSpareParts(pageSize, page)
	return ctx.JSON(&fiber.Map{
		"data":       spareParts,
		"totalCount": totalCount,
	})
}

func FindSparePart(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	sparePartDB, err := services.GetSparePartByID(id)
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"sparePart": sparePartDB,
	})
}

func CreateSparePart(ctx *fiber.Ctx) error {
	var sparePart models.SparePart
	ctx.BodyParser(&sparePart)
	sparePartDB, err := services.CreateSparePart(sparePart)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"sparePart": sparePartDB,
	})
}

func UpdateSparePart(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var sparePart models.SparePart
	ctx.BodyParser(&sparePart)
	sparePartDB, err := services.UpdateSparePart(id, sparePart)
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"sparePart": sparePartDB,
	})
}

func DeleteSparePart(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	deleted, err := services.DeleteSparePart(id)
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
