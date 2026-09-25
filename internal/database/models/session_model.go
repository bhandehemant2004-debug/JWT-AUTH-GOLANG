package models

import "time"

type SessionModel struct {
	ID               int       `json:"id" db:"id"`
	UserID           int       `json:"user_id" db:"user_id" validate:"required"`
	IsRevoked        bool      `json:"is_revoked" db:"is_revoked"`
	RefreshTokenHash string    `json:"refresh_token_hash" db:"refresh_token_hash" validate:"required"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
