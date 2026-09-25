package models

import "time"

type OTPModel struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id" validate:"required"`
	UserEmail string    `json:"user_email" db:"user_email" validate:"required,email"`
	OTP       string    `json:"otp" db:"otp" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
