package users

import (
	"SparkGuardBackend/db"
)

type GetUserRequest struct {
	ID uint `uri:"id"`
}

type GetUserResponse struct {
	db.User
}

type CreateUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	AccessLevel int    `json:"access_level" binding:"required"`

	Salt string `json:"salt" binding:"required"`
	Hash string `json:"hash" binding:"required"`
}

type EditUserRequest struct {
	ID uint `uri:"id"`
	db.User
}
