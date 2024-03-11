package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	VerifyEmail bool      `json:"verify_email"`
	Deleted     bool      `json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
