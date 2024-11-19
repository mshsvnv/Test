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

type LoginVerifyReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ResetPasswordReq struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
}

type VerifyResetPasswordReq struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
	Code        string `json:"code"`
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

type UpdateReq struct {
	Email string         `json:"email"`
	Role  model.UserRole `json:"role"`
}

type UpdatePasswordReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
