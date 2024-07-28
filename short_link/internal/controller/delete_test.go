package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/memory"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCustomLink(t *testing.T) {
	repo := memory.New()
	control := New(repo)

	//Token invalid len
	token := util.RandomString(util.MaxLenToken)
	customLink, err := control.DeleteCustomLink(context.Background(), token)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrInvalidCustomToken))
	assert.Nil(t, customLink)

	// Error not found
	token_2 := util.RandomString(10)
	customLink, err = control.DeleteCustomLink(context.Background(), token_2)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
	assert.Nil(t, customLink)
	//Token valid
	token_3 := util.RandomString(10)
	url := util.RandomURL(10)
	customLinkPrms := repository.CreateCustomLinkParams{
		Token: model.CustomLinkToken(token_3),
		Url:   url,
	}
	cLink, err := repo.PutCustomLink(context.Background(), customLinkPrms)
	assert.NoError(t, err)
	assert.NotEmpty(t, cLink)

	customLink, err = control.DeleteCustomLink(context.Background(), token_3)
	assert.NoError(t, err)
	assert.NotNil(t, customLink)
	assert.Equal(t, customLink.url, repository.DeleteStringValue)
}

func TestDeleteLink(t *testing.T) {
	repo := memory.New()
	control := New(repo)

	//Token invalid len
	token := util.RandomString(7)
	shortLink, err := control.DeleteLink(context.Background(), token)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, util.ErrInvalidToken))
	assert.Nil(t, shortLink)

	// Error not found
	token_2 := util.RandomString(util.MaxLenToken)
	shortLink, err = control.DeleteLink(context.Background(), token_2)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
	assert.Nil(t, shortLink)

	//Token valid
	url := util.RandomURL(9)
	user := repository.HasUserID{ID: uuid.New(), Valid: true}
	newShortLink, err := control.CreateShortLink(context.Background(), url, user)
	assert.NoError(t, err)
	assert.NotNil(t, newShortLink)
	assert.Equal(t, newShortLink.Url, url)

	shortLink, err = control.DeleteLink(context.Background(), newShortLink.Token)
	assert.NoError(t, err)
	assert.NotNil(t, shortLink)
	assert.Equal(t, shortLink.url, repository.DeleteStringValue)

	//Check the recicle link has been created
	n_sK, err := repo.GetRecycleLink(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, n_sK)
	assert.NotZero(t, n_sK.SKey)
}
