package dto

import (
	"github.com/google/uuid"
	"github.com/salawatbro/raxmet/pkg/validators"
)

type ExpenseDTO struct {
	GroupID     uuid.UUID `json:"group_id" validate:"omitempty,uuid"`
	UserID      uuid.UUID `json:"user_id" validate:"required,uuid"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"omitempty"`
	Amount      string    `json:"amount" validate:"required"`

	ExpenseShare []ExpenseShareDTO `json:"expense_share" validate:"required,dive"`
}

type ExpenseShareDTO struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid"`
	Amount string    `json:"amount" validate:"required"`
}

func (dto *ExpenseDTO) Validate() []error {
	return validators.ExtractValidationError(dto)
}
