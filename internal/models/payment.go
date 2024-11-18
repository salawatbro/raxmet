package models

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
	"time"
)

type Payment struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PaidBy    uuid.UUID       `json:"paid_by" gorm:"type:uuid;not null"`
	PaidTo    uuid.UUID       `json:"paid_to" gorm:"type:uuid;not null"`
	Amount    decimal.Decimal `json:"amount" gorm:"not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"default:now()"`
}

func (table *Payment) TableName() string {
	return "payments"
}
