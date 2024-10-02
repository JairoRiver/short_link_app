package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

func (c *Control) CheckCustomLinkIsFree(ctx context.Context, token string) (bool, error) {
	// Check if exist a custom token
	_, err := c.repo.GetCustomLinkByToken(ctx, model.CustomLinkToken(token))
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return true, nil
	}
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return false, fmt.Errorf("Control CheckCustomLink GetCustomLinkByToken error: %w", err)
	}

	return false, nil
}

type GetLinkResponse struct {
	Url string
}

func (c *Control) GetByCustomToken(ctx context.Context, token string) (*GetLinkResponse, error) {
	if len(token) > util.MaxLenToken {
		customLink, err := c.repo.GetCustomLinkByToken(ctx, model.CustomLinkToken(token))
		if err != nil {
			return nil, fmt.Errorf("Control GetByCustomToken GetCustomLinkByToken error: %w", err)
		}
		rps := GetLinkResponse{
			Url: customLink.Url,
		}
		return &rps, nil
	}
	return nil, fmt.Errorf("Control GetByCustomToken token len must be > 6, error: %w", ErrInvalidCustomToken)
}

func (c *Control) GetByToken(ctx context.Context, token string) (*GetLinkResponse, error) {
	s_k, err := decodingToken(token)
	if err != nil {
		return nil, fmt.Errorf("Control GetByToken decodingToken error: %w", err)
	}

	shortLink, err := c.repo.GetShortLinkBySKey(ctx, model.ShortLinkId(s_k))
	if err != nil {
		return nil, fmt.Errorf("Control GetByToken GetShortLinkBySKey error: %w", err)
	}

	rps := GetLinkResponse{
		Url: shortLink.Url,
	}

	return &rps, nil
}
