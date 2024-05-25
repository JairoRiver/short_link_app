package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

func (c *Controller) CheckCustomLinkIsFree(ctx context.Context, token string) (bool, error) {
	// Check if exist a custom token
	_, err := c.repo.GetCustomLinkByToken(ctx, model.CustomLinkToken(token))
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return true, nil
	}
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return false, fmt.Errorf("Controller CheckCustomLink GetCustomLinkByToken error: %w", err)
	}

	return false, nil
}

type GetLinkResponse struct {
	url string
}

func (c *Controller) GetByCustomToken(ctx context.Context, token string) (*GetLinkResponse, error) {
	if len(token) > util.MaxLenToken {
		customLink, err := c.repo.GetCustomLinkByToken(ctx, model.CustomLinkToken(token))
		if err != nil {
			return nil, fmt.Errorf("Controller GetByCustomToken GetCustomLinkByToken error: %w", err)
		}
		rps := GetLinkResponse{
			url: customLink.Url,
		}
		return &rps, nil
	}
	return nil, fmt.Errorf("Controller GetByCustomToken token len must be > 6, error: %w", ErrInvalidCustomToken)
}

func (c *Controller) GetByToken(ctx context.Context, token string) (*GetLinkResponse, error) {
	s_k, err := decodingToken(token)
	if err != nil {
		return nil, fmt.Errorf("Controller GetByToken decodingToken error: %w", err)
	}

	shortLink, err := c.repo.GetShortLinkBySKey(ctx, model.ShortLinkId(s_k))
	if err != nil {
		return nil, fmt.Errorf("Controller GetByToken GetShortLinkBySKey error: %w", err)
	}

	rps := GetLinkResponse{
		url: shortLink.Url,
	}

	return &rps, nil
}
