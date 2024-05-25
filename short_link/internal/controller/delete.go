package controller

import (
	"context"
	"fmt"

	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

type DeleteLinkResponse struct {
	url string
}

func (c *Controller) DeleteCustomLink(ctx context.Context, token string) (*DeleteLinkResponse, error) {
	if len(token) > util.MaxLenToken {
		customLink, err := c.repo.DeleteCustomLinkByToken(ctx, model.CustomLinkToken(token))
		if err != nil {
			return nil, fmt.Errorf("Controller DeleteCustomLink DeleteCustomLinkByToken error: %w", err)
		}
		rps := DeleteLinkResponse{
			url: customLink.Url,
		}
		return &rps, nil
	}
	return nil, fmt.Errorf("Controller DeleteCustomLink token len must be > 6, error: %w", ErrInvalidCustomToken)
}
