package transformer

import "github.com/salawatbro/raxmet/internal/models"

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func NewLoginResponse(token string, user models.User) LoginResponse {
	return LoginResponse{
		Token: token,
		User:  NewUserResponse(user),
	}
}
