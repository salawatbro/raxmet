package dto

import (
	"github.com/salawatbro/raxmet/internal/models"
	"github.com/salawatbro/raxmet/pkg/validators"
	"golang.org/x/crypto/bcrypt"
)

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (dto *LoginDTO) Validate() []error {
	return validators.ExtractValidationError(dto)
}

type RegisterDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (dto *RegisterDTO) Validate() []error {
	return validators.ExtractValidationError(dto)
}

func (dto *RegisterDTO) ToModel() (models.User, error) {
	if err := dto.HashPassword(); err != nil {
		return models.User{}, err
	}
	return models.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: dto.Password,
	}, nil
}

func (dto *RegisterDTO) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	dto.Password = string(hashedPassword)
	return nil
}
