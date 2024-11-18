package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/salawatbro/raxmet/internal/models"
	"github.com/salawatbro/raxmet/pkg/response"
)

type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
	FindAll(paginate response.Pagination) (*response.Pagination, error)
	ExistsByEmail(email string) bool
	FindByEmail(email string) (models.User, error)
	FindByID(id uuid.UUID) (models.User, error)
	Update(id uuid.UUID, user models.User) (models.User, error)
	UpdatePassword(id uuid.UUID, password string) error
	Delete(id uuid.UUID) error
}

type UserServiceInterface interface {
	FindAll(ctx *fiber.Ctx, paginate response.Pagination) error
	FindByEmail(ctx *fiber.Ctx, email string) error
	FindByID(ctx *fiber.Ctx, id uuid.UUID) error
	Update(ctx *fiber.Ctx, id uuid.UUID, user models.User) error
	UpdatePassword(ctx *fiber.Ctx, id uuid.UUID, password string) error
	Delete(ctx *fiber.Ctx, id uuid.UUID) error
}

type UserHandlerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByEmail(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UpdatePassword(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
