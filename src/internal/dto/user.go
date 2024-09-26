package dto

import (
	"src/internal/model"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	AccessToken string `json:"access_token"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRes struct {
	AccessToken string `json:"access_token"`
}

type UpdateRoleReq struct {
	ID   int            `json:"id"`
	Role model.UserRole `json:"role"`
}
