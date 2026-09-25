package models

import "time"

type UserModel struct {
	ID         int       `json:"id" db:"id"`
	Username   string    `json:"username" db:"username" validate:"required"`
	Email      string    `json:"email" db:"email" validate:"required"`
	Password   string    `json:"-" db:"password" validate:"required"`
	IsVerified bool      `json:"is_verified" db:"is_verified"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
