package repository

import (
	"context"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
)

type Storer interface {
	//Custom Link
	PutCustomLink(ctx context.Context, customLink model.CustomLink) error
	GetCustomLinkByID(ctx context.Context, customLinkID model.CustomLinkId) (*model.CustomLink, error)
	GetCustomLinkByToken(ctx context.Context, customLinkToken model.CustomLinkToken) (*model.CustomLink, error)
	ListAllCustomLink(ctx context.Context) ([]model.CustomLink, error)
	ListActiveCustomLink(ctx context.Context) ([]model.CustomLink, error)
	ListCustomLinkByUser(ctx context.Context, userId uuid.UUID) ([]model.CustomLink, error)
	DeleteCustomLink(ctx context.Context, customLinkID model.CustomLinkId) (*model.CustomLink, error)
	// Recycle Link
	PutRecycleLink(ctx context.Context, recycleLink model.RecycleLink) error
	GetRecycleLink(ctx context.Context, recycleLinkID model.RecycleLinkId) (*model.RecycleLink, error)
	DeleteRecycleLink(ctx context.Context, recycleLinkID model.RecycleLinkId) error
	// Short Link
	PutShortLink(ctx context.Context, shotLinkParams CreateShortLinkParams) (model.ShortLink, error)
	GetShortLinkByID(ctx context.Context, shotLinkID model.ShortLinkId) (*model.ShortLink, error)
	GetShortLinkBySKey(ctx context.Context, sKeyID model.ShortLinkId) (*model.ShortLink, error)
	ListAllShortLink(ctx context.Context) ([]model.ShortLink, error)
	ListActiveShortLink(ctx context.Context) ([]model.ShortLink, error)
	ListShortLinkByUser(ctx context.Context, userId uuid.UUID) ([]model.ShortLink, error)
	DeleteShortLink(ctx context.Context, shotLinkID model.ShortLinkId) (*model.ShortLink, error)
}

type HasUserID struct {
	ID    uuid.UUID
	Valid bool
}

// CreateCustomLinkParams struct use as param to create a new short link
type CreateShortLinkParams struct {
	UserId    HasUserID
	Url       string
	Token     string
	SKey      model.ShortLinkId
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateCustomLinkParams struct use as param to create a new custom link
type CreateCustomLinkParams struct {
	UserId       HasUserID
	Url          string
	Token        model.CustomLinkToken
	IsSuggestion bool
	SuggestionId model.SuggestionId
	Deleted      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
