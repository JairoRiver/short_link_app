package postgres

import (
	"context"
	"errors"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	minRandomSk = 100
	maxRandomSk = 100000
	urlLength   = 10
	listTestLen = 10
)

// createRandomShortLink create short link for test proposes
func createRandomShortLink(t *testing.T, args ...bool) model.ShortLink {
	status := false
	if len(args) > 0 {
		status = args[0]
	}
	arg := repository.CreateShortLinkParams{
		Url:     util.RandomURL(urlLength),
		Token:   util.RandomString(util.MaxLenToken),
		SKey:    model.ShortLinkId(util.RandomInt(minRandomSk, maxRandomSk)),
		Deleted: status,
	}
	shortLink, err := testStore.PutShortLink(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, shortLink)
	assert.NotEmpty(t, shortLink.Id)
	assert.Equal(t, arg.Url, shortLink.Url)
	assert.Equal(t, arg.Token, shortLink.Token)
	assert.Equal(t, arg.SKey, shortLink.SKey)
	assert.Equal(t, arg.Deleted, shortLink.Deleted)
	assert.NotEmpty(t, shortLink.CreatedAt)
	assert.NotEmpty(t, shortLink.UpdatedAt)

	return shortLink
}

func TestPutShortLink(t *testing.T) {
	_ = createRandomShortLink(t)
}

func TestGetShortLinkByID(t *testing.T) {
	shortLink := createRandomShortLink(t)

	// Test not found error
	getShortLinkError, err := testStore.GetShortLinkByID(context.Background(), shortLink.Id+1)
	assert.Empty(t, getShortLinkError)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test get link
	getShortLink, err := testStore.GetShortLinkByID(context.Background(), shortLink.Id)
	assert.NotNil(t, getShortLink)
	assert.Equal(t, *getShortLink, shortLink)
	assert.NoError(t, err)
}

func TestGetShortLinkBySKey(t *testing.T) {
	shortLink := createRandomShortLink(t)

	// Test not found error
	getShortLinkError, err := testStore.GetShortLinkBySKey(context.Background(), model.ShortLinkId(util.RandomInt(1, maxRandomSk+1)))
	assert.Empty(t, getShortLinkError)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test get link
	getShortLink, err := testStore.GetShortLinkBySKey(context.Background(), shortLink.SKey)
	assert.NotNil(t, getShortLink)
	assert.Equal(t, *getShortLink, shortLink)
	assert.NoError(t, err)
}

func TestListAllShortLink(t *testing.T) {

	for i := 0; i < listTestLen; i++ {
		_ = createRandomShortLink(t)
	}

	shotLinkList, err := testStore.ListAllShortLink(context.Background())
	assert.NoError(t, err)
	// The len should be <= listTestLen
	assert.LessOrEqual(t, listTestLen, len(shotLinkList))

	for _, v := range shotLinkList {
		assert.NotEmpty(t, v)
	}
}

func TestListActiveShortLink(t *testing.T) {

	for i := 0; i < listTestLen; i++ {
		_ = createRandomShortLink(t, false)
		_ = createRandomShortLink(t, true)
	}

	//test dont get deleted links
	shotLinkList, err := testStore.ListAllShortLink(context.Background())
	assert.NoError(t, err)
	activeLinksList, err := testStore.ListActiveShortLink(context.Background())
	assert.NoError(t, err)
	assert.NotEqual(t, len(shotLinkList), len(activeLinksList))

	// The filed deleted must be false
	for _, v := range activeLinksList {
		assert.False(t, v.Deleted)
	}
}

func TestListShortLinkByUser(t *testing.T) {
	//TODO add this test when implement user repository
}

func TestDeleteShortLink(t *testing.T) {
	shortLink := createRandomShortLink(t)
	args := DeleteShortLinkParams{
		ID: int64(shortLink.Id),
	}
	// Test delete link
	deletedLink, err := testStore.DeleteShortLink(context.Background(), args)
	assert.NotNil(t, deletedLink)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedLink)
	assert.Equal(t, deletedLink.Url, repository.DeleteStringValue)
	assert.Zero(t, deletedLink.Token)
	assert.Zero(t, deletedLink.SKey)
	assert.True(t, deletedLink.Deleted)
	assert.Greater(t, deletedLink.UpdatedAt, deletedLink.CreatedAt)

	deletedLinkSk, err := testStore.GetShortLinkBySKey(context.Background(), shortLink.SKey)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
	assert.Empty(t, deletedLinkSk)

	// Test Not Found error
	argsErr := DeleteShortLinkParams{
		ID: util.RandomInt(maxRandomSk, maxRandomSk*2),
	}
	errorLink, err := testStore.DeleteShortLink(context.Background(), argsErr)
	assert.Empty(t, errorLink)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
}

func TestDeleteShortLinkBySK(t *testing.T) {
	shortLink := createRandomShortLink(t)
	args := DeleteShortLinkBySkParams{
		SKey: int64(shortLink.SKey),
	}
	// Test delete link
	deletedLink, err := testStore.DeleteShortLinkBySK(context.Background(), args)
	assert.NotNil(t, deletedLink)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedLink)
	assert.Equal(t, deletedLink.Url, repository.DeleteStringValue)
	assert.Zero(t, deletedLink.Token)
	assert.Zero(t, deletedLink.SKey)
	assert.True(t, deletedLink.Deleted)
	assert.Greater(t, deletedLink.UpdatedAt, deletedLink.CreatedAt)

	deletedLinkSk, err := testStore.GetShortLinkBySKey(context.Background(), shortLink.SKey)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
	assert.Empty(t, deletedLinkSk)

	// Test Not Found error
	argsErr := DeleteShortLinkBySkParams{
		SKey: util.RandomInt(maxRandomSk, maxRandomSk*2),
	}
	errorLink, err := testStore.DeleteShortLinkBySK(context.Background(), argsErr)
	assert.Empty(t, errorLink)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
}
