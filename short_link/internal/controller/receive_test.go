package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/memory"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestCheckCustomLinkIsFree(t *testing.T) {
	repo := memory.New()
	control := New(repo)

	//valid token is free
	token := util.RandomString(10)
	isFree, err := control.CheckCustomLinkIsFree(context.Background(), token)
	assert.NoError(t, err)
	assert.True(t, isFree)

	//token are in use
	token_2 := util.RandomString(10)
	customLinkPrms := repository.CreateCustomLinkParams{
		Token: model.CustomLinkToken(token_2),
	}
	cLink, err := repo.PutCustomLink(context.Background(), customLinkPrms)
	assert.NoError(t, err)
	assert.NotEmpty(t, cLink)

	inUse, err := control.CheckCustomLinkIsFree(context.Background(), token_2)
	assert.NoError(t, err)
	assert.False(t, inUse)
}

func TestGetByCustomToken(t *testing.T) {
	repo := memory.New()
	control := New(repo)

	//Token invalid len
	token := util.RandomString(6)
	customLinkPrms := repository.CreateCustomLinkParams{
		Token: model.CustomLinkToken(token),
	}
	cLink, err := repo.PutCustomLink(context.Background(), customLinkPrms)
	assert.NoError(t, err)
	assert.NotEmpty(t, cLink)

	customLink, err := control.GetByCustomToken(context.Background(), token)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrInvalidCustomToken))
	assert.Nil(t, customLink)

	//Token valid
	token_2 := util.RandomString(10)
	url := util.RandomURL(10)
	customLinkPrms = repository.CreateCustomLinkParams{
		Token: model.CustomLinkToken(token_2),
		Url:   url,
	}
	cLink, err = repo.PutCustomLink(context.Background(), customLinkPrms)
	assert.NoError(t, err)
	assert.NotEmpty(t, cLink)

	customLink, err = control.GetByCustomToken(context.Background(), token_2)
	assert.NoError(t, err)
	assert.NotNil(t, customLink)
	assert.Equal(t, customLink.url, url)
}
