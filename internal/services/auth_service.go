package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/internal/dto"
	"github.com/salawatbro/raxmet/internal/interfaces"
	"github.com/salawatbro/raxmet/internal/transformer"
	"github.com/salawatbro/raxmet/pkg/constants"
	"github.com/salawatbro/raxmet/pkg/response"
)

type AuthService struct {
	repo interfaces.UserRepositoryInterface
}

func NewAuthService(repo interfaces.UserRepositoryInterface) interfaces.AuthServiceInterface {
	return &AuthService{repo: repo}
}

func (service *AuthService) Login(ctx *fiber.Ctx, dto *dto.LoginDTO) error {
	if err := dto.Validate(); len(err) > 0 {
		return response.JsonErrorValidation(ctx, err)
	}
	user, err := service.repo.FindByEmail(dto.Email)
	if err != nil {
		return response.JsonError(ctx, []error{constants.ErrInvalidEmailOrPassword}, "invalid_credentials")
	}
	if !user.ComparePassword(dto.Password) {
		return response.JsonError(ctx, []error{constants.ErrInvalidEmailOrPassword}, "invalid_credentials")
	}
	token := user.GenerateToken()
	if token == "" {
		return response.JsonError(ctx, []error{constants.ErrSomethingWentWrong}, "token_generation_failed")
	}
	return response.JsonSuccess(ctx, transformer.NewLoginResponse(token, user))
}

func (service *AuthService) Register(ctx *fiber.Ctx, dto *dto.RegisterDTO) error {
	if err := dto.Validate(); len(err) > 0 {
		return response.JsonErrorValidation(ctx, err)
	}
	if service.repo.ExistsByEmail(dto.Email) {
		return response.JsonError(ctx, []error{constants.ErrEmailExists}, "email_exists")
	}
	user, err := dto.ToModel()
	if err != nil {
		return response.JsonError(ctx, []error{constants.ErrSomethingWentWrong}, "user_creation_failed")
	}
	user, err = service.repo.Create(user)
	if err != nil {
		return response.JsonError(ctx, []error{constants.ErrSomethingWentWrong}, "user_creation_failed")
	}
	return response.JsonSuccess(ctx, user)
}

func (service *AuthService) Logout(ctx *fiber.Ctx) error {
	return response.JsonSuccess(ctx, nil)
}
