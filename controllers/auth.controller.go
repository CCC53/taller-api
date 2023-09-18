package controllers

import (
	"taller-api/models"
	"taller-api/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	var loginDTO models.LoginDTO
	ctx.BodyParser(&loginDTO)
	employee, err := services.FindByCredentials(loginDTO.Email, loginDTO.Password)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(&fiber.Map{
			"error": "Email o contraseña incorrectos",
		})
	}
	token, _ := services.GenerateToken(*employee)
	ctx.Cookie(&fiber.Cookie{
		Name:     "authorization",
		Value:    token,
		HTTPOnly: true,
		Expires:  time.Now().Add((time.Hour) * 1),
		Secure:   true,
		SameSite: "none",
	})
	return ctx.JSON(&fiber.Map{
		"routes": services.LoadMenu(employee.Role),
		"token":  token,
	})
}

func ValidateToken(ctx *fiber.Ctx) error {
	var tokenString = ctx.Cookies("authorization")
	valid := services.ValidateToken(tokenString)
	if !valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"valid": valid,
		})
	}
	return ctx.JSON(&fiber.Map{
		"valid": valid,
	})
}

func ValidateRole(ctx *fiber.Ctx) error {
	var tokenString = ctx.Cookies("authorization")
	valid := services.ValidateRole(tokenString)
	if !valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"valid": valid,
		})
	}
	return ctx.JSON(&fiber.Map{
		"valid": valid,
	})
}

func GetMe(ctx *fiber.Ctx) error {
	var tokenString = ctx.Cookies("authorization")
	me := services.GetMe(tokenString)
	return ctx.JSON(&fiber.Map{
		"me": services.MappingEmployee(*me),
	})
}

func UpdateProfile(ctx *fiber.Ctx) error {
	var tokenString = ctx.Cookies("authorization")
	var profileFormData models.Employee
	ctx.BodyParser(&profileFormData)
	updated := services.UpdateProfile(tokenString, profileFormData)
	return ctx.JSON(&fiber.Map{
		"me": updated,
	})
}

func Logout(ctx *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 1)
	ctx.Cookie(&fiber.Cookie{
		Name:    "authorization",
		Value:   "",
		Expires: expired,
	})
	return ctx.JSON(&fiber.Map{
		"logout": true,
	})
}
