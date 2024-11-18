package dto

import "github.com/salawatbro/raxmet/pkg/validators"

type GroupDTO struct {
	Name string `json:"name" validate:"required"`
}

func (dto *GroupDTO) Validate() []error {
	return validators.ExtractValidationError(dto)
}
