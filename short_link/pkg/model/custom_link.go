package model

import (
	"time"

	"github.com/google/uuid"
)

type CustomLink struct {
	Id           int       `json:"id"`
	UserId       uuid.UUID `json:"user_id"`
	Url          string    `json:"url"`
	Token        string    `json:"token"`
	IsSuggestion bool      `json:"is_suggestion"`
	SuggestionId int       `json:"suggestion_id"`
	Deleted      bool      `json:"deleted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
