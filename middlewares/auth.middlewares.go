package middlewares

import (
	"taller-api/config"
	"taller-api/enums"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthenticated(ctx *fiber.Ctx) error {
	var tokenString string
	if ctx.Cookies("authorization") != "" {
		tokenString = ctx.Cookies("authorization")
	}
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "No estas autenticado",
		})
	}
	return ctx.Next()
}

func IsAdmin(ctx *fiber.Ctx) error {
	var tokenString = ctx.Cookies("authorization")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "No estas autenticado",
		})
	}
	if claims["role"] == string(enums.Mechanic) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "No eres usuario administrador",
		})
	}
	return ctx.Next()
}
