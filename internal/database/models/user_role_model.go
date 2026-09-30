package models

import "time"

type UserRoleModel struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id" validate:"required"`
	RoleID    int       `json:"role_id" db:"role_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
