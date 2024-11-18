package transformer

import (
	"github.com/salawatbro/raxmet/internal/models"
	"time"
)

type UserResponse struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	ProfilePic      string `json:"profile_pic"`
	Role            string `json:"role"`
	IsEmailVerified bool   `json:"is_email_verified"`
	CreatedAt       string `json:"created_at"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:              user.Id.String(),
		Email:           user.Email,
		Name:            user.Name,
		ProfilePic:      user.ProfilePic,
		Role:            user.Role,
		IsEmailVerified: user.IsEmailVerified,
		CreatedAt:       user.CreatedAt.Format(time.RFC3339),
	}
}
