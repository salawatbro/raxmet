package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func getAuthorizationToken(ctx *fiber.Ctx) string {
	authorizationToken := string(ctx.Request().Header.Peek("Authorization"))
	return strings.Replace(authorizationToken, "Bearer ", "", 1)
}
