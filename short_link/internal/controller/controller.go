package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
)

type Controller interface {
	CreateShortLink(ctx context.Context, url string, userId repository.HasUserID) (*ShortLinkResponse, error)
	CreateCustomLink(ctx context.Context, url string, userId repository.HasUserID, token string) (*ShortLinkResponse, error)
	DeleteCustomLink(ctx context.Context, token string) (*DeleteLinkResponse, error)
	DeleteLink(ctx context.Context, token string) (*DeleteLinkResponse, error)
	CheckCustomLinkIsFree(ctx context.Context, token string) (bool, error)
	GetByCustomToken(ctx context.Context, token string) (*GetLinkResponse, error)
	GetByToken(ctx context.Context, token string) (*GetLinkResponse, error)
}

// Control defines a short link service controller.
type Control struct {
	repo repository.Storer
}

// New creates a short link service controller.
func New(repo repository.Storer) *Control {
	return &Control{repo}
}

var ErrInvalidCustomToken = errors.New("error invalid custom token")

func decodingToken(token string) (uint64, error) {
	if len(token) > util.MaxLenToken {
		return 0, fmt.Errorf("Control GetByToken token len must be lower than 7, error: %w", util.ErrInvalidToken)
	}

	s_k, err := util.FromBase62(token)
	if err != nil {
		return 0, fmt.Errorf("Control GetByToken util.FromBase62 error: %w", err)
	}

	return s_k, nil
}
