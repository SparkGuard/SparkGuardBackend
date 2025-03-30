package users

import (
	"SparkGuardBackend/internal/db"
)

type GetUserRequest struct {
	ID uint `uri:"id"`
}

type GetUserResponse struct {
	db.User
}

type CreateUserRequest struct {
	db.User
}

type EditUserRequest struct {
	ID uint `uri:"id"`
	db.User
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string   `json:"token"`
	User  *db.User `json:"user"`
}
