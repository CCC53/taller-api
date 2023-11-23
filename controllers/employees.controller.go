package controllers

import (
	"strings"
	"taller-api/models"
	"taller-api/services"

	"github.com/gofiber/fiber/v2"
)

func FindEmployees(ctx *fiber.Ctx) error {
	var tokenString = ctx.Cookies("authorization")
	pageSize := ctx.QueryInt("pageSize", 5)
	page := ctx.QueryInt("page", 1)
	employees, totalCount := services.ListEmployees(pageSize, page, tokenString)
	return ctx.JSON(&fiber.Map{
		"employees":  employees,
		"totalCount": totalCount,
	})
}

func FindEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	employee, err := services.GetEmployeeByID(id)
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"employee": services.MappingEmployee(employee),
	})
}

func CreateEmployee(ctx *fiber.Ctx) error {
	var employee models.Employee
	ctx.BodyParser(&employee)
	employeeRes, err := services.CreateEmployee(employee)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		if strings.Contains(err.Error(), "duplicate") {
			return ctx.JSON(&fiber.Map{
				"error": "Ya existe un usuario con este email",
			})
		}
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"employee": employeeRes,
	})
}

func UpdateEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var employee models.Employee
	ctx.BodyParser(&employee)
	employeeRes, err := services.UpdateEmployee(id, employee)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		if strings.Contains(err.Error(), "duplicate") {
			return ctx.JSON(&fiber.Map{
				"error": "Ya existe un usuario con este email",
			})
		}
		return ctx.JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(&fiber.Map{
		"employee": employeeRes,
	})
}

func DeleteEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	deleted, err := services.DeleteEmployee(id)
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
