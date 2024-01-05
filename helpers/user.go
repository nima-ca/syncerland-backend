package helpers

import (
	"syncerland/app/jwt"
	"syncerland/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) *models.User {
	return ctx.Locals(jwt.UserLocalKey).(*models.User)
}
