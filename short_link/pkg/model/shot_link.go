package model

import (
	"time"

	"github.com/google/uuid"
)

type ShortLink struct {
	Id        int       `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Url       string    `json:"url"`
	Token     string    `json:"token"`
	SKey      int       `json:"s_key"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
