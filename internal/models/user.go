package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name            string    `json:"name" gorm:"varchar(255);not null"`
	Email           string    `json:"email" gorm:"varchar(255);unique;not null"`
	Password        string    `json:"password" gorm:"varchar(255);not null"`
	ProfilePic      string    `json:"profile_pic" gorm:"varchar(255)'"`
	IsEmailVerified bool      `json:"is_email_verified" gorm:"default:false"`
	Role            string    `json:"role" gorm:"varchar(255);default:'user'"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:now()"`
}
