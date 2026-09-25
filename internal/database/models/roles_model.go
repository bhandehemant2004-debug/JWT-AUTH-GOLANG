package models

import "time"

type RoleModel struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required"`
	Description string    `json:"description" db:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
