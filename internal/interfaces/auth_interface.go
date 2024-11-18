package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/internal/dto"
)

type AuthHandlerInterface interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}

type AuthServiceInterface interface {
	Login(ctx *fiber.Ctx, dto *dto.LoginDTO) error
	Register(ctx *fiber.Ctx, dto *dto.RegisterDTO) error
	Logout(ctx *fiber.Ctx) error
}
