package dto

import (
	"github.com/google/uuid"
	"github.com/salawatbro/raxmet/pkg/validators"
)

type PaymentDTO struct {
	PaidBy uuid.UUID `json:"paid_by" validate:"required,uuid"`
	PaidTo uuid.UUID `json:"paid_to" validate:"required,uuid"`
	Amount string    `json:"amount" validate:"required"`
}

func (dto *PaymentDTO) Validate() []error {
	return validators.ExtractValidationError(dto)
}
