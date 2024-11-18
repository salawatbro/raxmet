package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/salawatbro/raxmet/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id              uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name            string    `json:"name" gorm:"varchar(255);not null"`
	Email           string    `json:"email" gorm:"varchar(255);unique;not null"`
	Password        string    `json:"password" gorm:"varchar(255);not null"`
	ProfilePic      string    `json:"profile_pic" gorm:"varchar(255)'"`
	IsEmailVerified bool      `json:"is_email_verified" gorm:"default:false"`
	Role            string    `json:"role" gorm:"varchar(255);default:'user'"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:now()"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (user *User) GenerateToken() string {
	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * config.Cfg.JWT.Exp).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token and return
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}
	return tokenString
}
