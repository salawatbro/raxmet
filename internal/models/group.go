package models

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID           uuid.UUID     `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name         string        `json:"name" gorm:"varchar(255);not null"`
	OwnerID      uuid.UUID     `json:"owner_id" gorm:"type:uuid;not null"`
	CreatedAt    time.Time     `json:"created_at" gorm:"default:now()"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"default:now()"`
	Owner        User          `json:"created_by_user" gorm:"foreignKey:owner_id"`
	GroupMembers []GroupMember `json:"group_members" gorm:"foreignKey:group_id"`
}

func (table *Group) TableName() string {
	return "groups"
}

type GroupMember struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	GroupID  uuid.UUID `json:"group_id" gorm:"type:uuid;not null"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	JoinedAt time.Time `json:"joined_at" gorm:"default:now()"`
}

func (table *GroupMember) TableName() string {
	return "group_members"
}
