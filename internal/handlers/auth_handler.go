package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/internal/dto"
	"github.com/salawatbro/raxmet/internal/interfaces"
	"github.com/salawatbro/raxmet/pkg/response"
)

type AuthHandler struct {
	service interfaces.AuthServiceInterface
}

func (handler *AuthHandler) Login(ctx *fiber.Ctx) error {
	req := new(dto.LoginDTO)
	if err := ctx.BodyParser(req); err != nil {
		return response.JsonErrorValidation(ctx, []error{err})
	}
	return handler.service.Login(ctx, req)
}

func (handler *AuthHandler) Register(ctx *fiber.Ctx) error {
	req := new(dto.RegisterDTO)
	if err := ctx.BodyParser(req); err != nil {
		return response.JsonErrorValidation(ctx, []error{err})
	}
	return handler.service.Register(ctx, req)
}

func (handler *AuthHandler) Logout(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthHandler(service interfaces.AuthServiceInterface) interfaces.AuthHandlerInterface {
	return &AuthHandler{service: service}
}
