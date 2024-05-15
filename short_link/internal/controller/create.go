package controller

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

// CreateShortLink generate an automatic new short link
type ShortLinkResponse struct {
	url   string
	token string
}

func (c *Controller) CreateShortLink(ctx context.Context, url string, userId repository.HasUserID) (*ShortLinkResponse, error) {
	var s_key int64
	// first step find if exist a recycle key
	recicleLink, err := c.repo.GetRecycleLink(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			// Here we generated a new s_key
			aux_s_key, err := c.repo.GetAuxSKey(ctx)
			if err != nil {
				return nil, fmt.Errorf("Controller CreateShortLink GetAuxSKey error: %w", err)
			}

			s_key = int64(aux_s_key.A0) + (int64(aux_s_key.N0) * int64(aux_s_key.Step))

			updateSKParams := repository.AuxSKeyParams{}
			if aux_s_key.N0 < aux_s_key.N {
				updateSKParams.N0 = repository.IsIntValid{Valid: true, Value: uint(aux_s_key.N0) + 1}
			}
			if aux_s_key.N0 == aux_s_key.N {
				updateSKParams.N0 = repository.IsIntValid{Valid: true, Value: 0}
				updateSKParams.A0 = repository.IsIntValid{Valid: true, Value: aux_s_key.A0 + 1}
			}

			// Update sKeyParams
			_, err = c.repo.UpdateAuxSKey(ctx, updateSKParams)
			if err != nil {
				return nil, fmt.Errorf("Controller CreateShortLink UpdateAuxSKey error: %w", err)
			}
		} else {
			return nil, fmt.Errorf("Controller CreateShortLink GetRecycleLink error: %w", err)
		}

	} else {
		// If we have a recycle link we will use this link
		s_key = int64(recicleLink.SKey)
		err = c.repo.DeleteRecycleLink(ctx, recicleLink.SKey)
		if err != nil {
			return nil, fmt.Errorf("Controller CreateShortLink DeleteRecycleLink error: %w", err)
		}
	}

	// Second step genate the token, s_key to base62 encoding
	token := big.NewInt(s_key).Text(62)

	//Create short link
	customLinkParams := repository.CreateShortLinkParams{
		UserId:    userId,
		Url:       url,
		Token:     token,
		SKey:      model.ShortLinkId(s_key),
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	shortLink, err := c.repo.PutShortLink(ctx, customLinkParams)
	if err != nil {
		return nil, fmt.Errorf("Controller CreateShortLink PutShortLink error: %w", err)
	}

	shortLinkRsp := ShortLinkResponse{
		url:   shortLink.Url,
		token: shortLink.Token,
	}
	return &shortLinkRsp, nil
}
