package user

import (
	"github.com/smartbot/account/database"
)

// validation models
type OnboardRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type CreateUserRequest struct {
	Username       string `json:"username" validate:"required,email"`
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	PrimaryAddress string `json:"primary_address"`
	Mobile         string `json:"mobile" validate:"mobileNo"`
	Role           string `json:"role" validate:"required,oneof=USER SUPERVISOR"`
}

type UpdateUserRequest struct {
	Username       string `json:"username,omitempty" validate:"omitempty,email"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	PrimaryAddress string `json:"primary_address,omitempty"`
	Mobile         string `json:"mobile,omitempty" validate:"omitempty,mobileNo"`
	Role           string `json:"role,omitempty" validate:"omitempty,oneof=USER SUPERVISOR"`
}

type UserResponse struct {
	ID             string              `json:"id"`
	Username       string              `json:"username"`
	FirstName      string              `json:"first_name"`
	FullName       string              `json:"full_name"`
	LastName       string              `json:"last_name"`
	PrimaryAddress string              `json:"primary_address"`
	Mobile         string              `json:"mobile"`
	Role           database.UserRole   `json:"role"`
	Status         database.UserStatus `json:"status"`
	Avatar         string              `json:"avatar"`
	CreatedAt      string              `json:"created_at"`
}

type UsersRequest struct {
	PageNo   int `form:"page_no" validate:"required,gte=1"`
	PageSize int `form:"page_size" validate:"required,gte=1"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
	Total int64          `json:"total"`
}
