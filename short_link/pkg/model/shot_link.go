package model

import (
	"time"

	"github.com/google/uuid"
)

type ShortLinkId int

type ShortLink struct {
	Id        ShortLinkId `json:"id"`
	UserId    uuid.UUID   `json:"user_id"`
	Url       string      `json:"url"`
	Token     string      `json:"token"`
	SKey      ShortLinkId `json:"s_key"`
	Deleted   bool        `json:"deleted"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
