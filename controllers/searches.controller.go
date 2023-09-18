package controllers

import (
	"taller-api/services"

	"github.com/gofiber/fiber/v2"
)

func SearchByTable(ctx *fiber.Ctx) error {
	table := ctx.Params("table")
	search := ctx.Query("search")
	pageSize := ctx.QueryInt("pageSize", 5)
	page := ctx.QueryInt("page", 1)
	var data = services.Search(table, search, page, pageSize)
	return ctx.JSON(&fiber.Map{
		"data":       data["data"],
		"totalCount": data["total"],
	})
}

func ResotreData(ctx *fiber.Ctx) error {
	var tokenString = ctx.Cookies("authorization")
	table := ctx.Params("table")
	pageSize := ctx.QueryInt("pageSize", 5)
	page := ctx.QueryInt("page", 1)
	var data = services.ResotreData(table, pageSize, page, tokenString)
	return ctx.JSON(&fiber.Map{
		"data":       data["data"],
		"totalCount": data["total"],
	})
}
