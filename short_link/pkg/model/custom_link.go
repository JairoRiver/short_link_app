package model

import (
	"time"

	"github.com/google/uuid"
)

type CustomLinkId int
type SuggestionId int
type CustomLinkToken string

type CustomLink struct {
	Id           CustomLinkId    `json:"id"`
	UserId       uuid.UUID       `json:"user_id"`
	Url          string          `json:"url"`
	Token        CustomLinkToken `json:"token"`
	IsSuggestion bool            `json:"is_suggestion"`
	SuggestionId SuggestionId    `json:"suggestion_id"`
	Deleted      bool            `json:"deleted"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}
