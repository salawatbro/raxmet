package dto

import "github.com/salawatbro/raxmet/pkg/validators"

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (dto *LoginDTO) Validate() []string {
	return validators.ExtractValidationError(dto)
}

type RegisterDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (dto *RegisterDTO) Validate() []string {
	return validators.ExtractValidationError(dto)
}
