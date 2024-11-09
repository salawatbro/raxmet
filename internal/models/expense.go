package models

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
	"time"
)

type Expense struct {
	ID          uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	GroupID     uuid.UUID       `json:"group_id" gorm:"type:uuid"`
	UserID      uuid.UUID       `json:"user_id" gorm:"type:uuid;not null"`
	Amount      decimal.Decimal `json:"amount" gorm:"not null"`
	Title       string          `json:"title" gorm:"varchar(255);not null"`
	Description string          `json:"description" gorm:"text"`
	CreatedAt   time.Time       `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"default:now()"`

	User         User           `json:"user" gorm:"foreignKey:user_id"`
	Group        Group          `json:"group" gorm:"foreignKey:group_id"`
	ExpenseShare []ExpenseShare `json:"expense_share" gorm:"foreignKey:expense_id"`
}

type ExpenseShare struct {
	ID          uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ExpenseID   uuid.UUID       `json:"expense_id" gorm:"type:uuid;not null"`
	UserID      uuid.UUID       `json:"user_id" gorm:"type:uuid;not null"`
	ShareAmount decimal.Decimal `json:"share_amount" gorm:"not null"`
}
