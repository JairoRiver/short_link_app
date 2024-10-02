package controller

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

// CreateShortLink generate an automatic new short link
type ShortLinkResponse struct {
	Url   string
	Token string
}

func (c *Control) CreateShortLink(ctx context.Context, url string, userId repository.HasUserID) (*ShortLinkResponse, error) {
	var s_key uint64
	// first step find if exist a recycle key
	recicleLink, err := c.repo.GetRecycleLink(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			// Here we generated a new s_key
			aux_s_key, err := c.repo.GetAuxSKey(ctx)
			if err != nil {
				return nil, fmt.Errorf("Control CreateShortLink GetAuxSKey error: %w", err)
			}

			s_key = uint64(aux_s_key.A0) + (uint64(aux_s_key.N0) * uint64(aux_s_key.Step))

			updateSKParams := repository.AuxSKeyParams{}
			if aux_s_key.N0 < aux_s_key.N {
				updateSKParams.N0 = repository.IsIntValid{Valid: true, Value: uint(aux_s_key.N0) + 1}
			}
			if aux_s_key.N0 == aux_s_key.N {
				updateSKParams.N0 = repository.IsIntValid{Valid: true, Value: 1}
				updateSKParams.A0 = repository.IsIntValid{Valid: true, Value: aux_s_key.A0 + 1}
			}

			// Update sKeyParams
			_, err = c.repo.UpdateAuxSKey(ctx, updateSKParams)
			if err != nil {
				return nil, fmt.Errorf("Control CreateShortLink UpdateAuxSKey error: %w", err)
			}
		} else {
			return nil, fmt.Errorf("Control CreateShortLink GetRecycleLink error: %w", err)
		}

	} else {
		// If we have a recycle link we will use this link
		s_key = uint64(recicleLink.SKey)
		err = c.repo.DeleteRecycleLink(ctx, recicleLink.SKey)
		if err != nil {
			return nil, fmt.Errorf("Control CreateShortLink DeleteRecycleLink error: %w", err)
		}
	}

	// Second step genate the token, s_key to base62 encoding
	token := util.ToBase62(s_key)

	//Create short link
	shortLinkParams := repository.CreateShortLinkParams{
		UserId:    userId,
		Url:       url,
		Token:     token,
		SKey:      model.ShortLinkId(s_key),
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	shortLink, err := c.repo.PutShortLink(ctx, shortLinkParams)
	if err != nil {
		return nil, fmt.Errorf("Control CreateShortLink PutShortLink error: %w", err)
	}

	shortLinkRsp := ShortLinkResponse{
		Url:   shortLink.Url,
		Token: shortLink.Token,
	}
	return &shortLinkRsp, nil
}

func (c *Control) CreateCustomLink(ctx context.Context, url string, userId repository.HasUserID, token string) (*ShortLinkResponse, error) {
	// custom token len must be longer than 6
	if len(token) <= util.MaxLenToken {
		return nil, fmt.Errorf("Control CreateCustomLink len token <= %d error: %w", util.MaxLenToken, ErrInvalidCustomToken)
	}

	if !util.IsValidURLPath(token) {
		return nil, fmt.Errorf("Control CreateCustomLink check tokens characteres dont accept special characters error: %w", ErrInvalidCustomToken)
	}

	customLinkParams := repository.CreateCustomLinkParams{
		UserId:    userId,
		Url:       url,
		Token:     model.CustomLinkToken(token),
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	customLink, err := c.repo.PutCustomLink(ctx, customLinkParams)
	if err != nil {
		return nil, fmt.Errorf("Control CreateCustomLink PutCustomLink error: %w", err)
	}

	customLinkRsp := ShortLinkResponse{
		Url:   customLink.Url,
		Token: string(customLink.Token),
	}
	return &customLinkRsp, nil
}
