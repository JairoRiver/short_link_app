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

func (c *Control) DeleteCustomLink(ctx context.Context, token string) (*DeleteLinkResponse, error) {
	if len(token) > util.MaxLenToken {
		customLink, err := c.repo.DeleteCustomLinkByToken(ctx, model.CustomLinkToken(token))
		if err != nil {
			return nil, fmt.Errorf("Control DeleteCustomLink DeleteCustomLinkByToken error: %w", err)
		}
		rps := DeleteLinkResponse{
			url: customLink.Url,
		}
		return &rps, nil
	}
	return nil, fmt.Errorf("Control DeleteCustomLink token len must be > 6, error: %w", ErrInvalidCustomToken)
}

func (c *Control) DeleteLink(ctx context.Context, token string) (*DeleteLinkResponse, error) {
	s_k, err := decodingToken(token)
	if err != nil {
		return nil, fmt.Errorf("Control DeleteLink decodingToken error: %w", err)
	}
	shortLink, err := c.repo.DeleteShortLinkBySK(ctx, model.ShortLinkId(s_k))
	if err != nil {
		return nil, fmt.Errorf("Control DeleteLink DeleteShortLink error: %w", err)
	}

	// Insert the sK on the table recycle link
	err = c.repo.PutRecycleLink(ctx, model.RecycleLink{SKey: model.RecycleLinkId(s_k)})
	if err != nil {
		return nil, fmt.Errorf("Control DeleteLink PutRecycleLink error: %w", err)
	}

	rps := DeleteLinkResponse{
		url: shortLink.Url,
	}
	return &rps, nil
}
