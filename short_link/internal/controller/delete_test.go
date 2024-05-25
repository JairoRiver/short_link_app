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
