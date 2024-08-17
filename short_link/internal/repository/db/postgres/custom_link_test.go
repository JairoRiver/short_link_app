package postgres

import (
	"context"
	"errors"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// createRandomShortLink create short link for test proposes
func createRandomCustomLink(t *testing.T, token string, args ...bool) model.CustomLink {
	status := false
	if len(args) > 0 {
		status = args[0]
	}
	arg := repository.CreateCustomLinkParams{
		Url:     util.RandomURL(urlLength),
		Token:   model.CustomLinkToken(util.RandomString(util.MaxLenToken + 1)),
		Deleted: status,
	}
	if len(token) > 0 {
		arg.Token = model.CustomLinkToken(token)
	}
	customLink, err := testStore.PutCustomLink(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, customLink)
	assert.NotEmpty(t, customLink.Id)
	assert.Equal(t, arg.Url, customLink.Url)
	assert.Equal(t, arg.Token, customLink.Token)
	assert.Equal(t, arg.Deleted, customLink.Deleted)
	assert.NotEmpty(t, customLink.CreatedAt)
	assert.NotEmpty(t, customLink.UpdatedAt)

	if len(token) > 0 {
		// Test error duplicated token
		customLinkDuplicate, err := testStore.PutCustomLink(context.Background(), arg)
		assert.Error(t, err)
		assert.Empty(t, customLinkDuplicate)
	}

	return customLink
}

func TestPutCustomLink(t *testing.T) {
	_ = createRandomCustomLink(t, "")
}

func TestGetCustomLinkByID(t *testing.T) {
	customLink := createRandomCustomLink(t, "")

	// Test not found error
	getCustomLinkError, err := testStore.GetCustomLinkByID(context.Background(), customLink.Id+1)
	assert.Empty(t, getCustomLinkError)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test get link
	getCustomLink, err := testStore.GetCustomLinkByID(context.Background(), customLink.Id)
	assert.NoError(t, err)
	assert.NotNil(t, getCustomLink)
	assert.Equal(t, *getCustomLink, customLink)
	assert.Equal(t, customLink.Id, getCustomLink.Id)
	assert.Equal(t, customLink.Url, getCustomLink.Url)
	assert.Equal(t, customLink.Token, getCustomLink.Token)
	assert.Equal(t, customLink.Deleted, getCustomLink.Deleted)
	assert.Equal(t, customLink.IsSuggestion, getCustomLink.IsSuggestion)
	assert.Equal(t, customLink.SuggestionId, getCustomLink.SuggestionId)
	assert.Equal(t, customLink.CreatedAt, getCustomLink.CreatedAt)
	assert.Equal(t, customLink.UpdatedAt, getCustomLink.UpdatedAt)

}

func TestGetCustomLinkByToken(t *testing.T) {
	customLink := createRandomCustomLink(t, "")

	// Test not found error
	getCustomLinkError, err := testStore.GetCustomLinkByToken(context.Background(), model.CustomLinkToken(util.RandomString(util.MaxLenToken+1)))
	assert.Empty(t, getCustomLinkError)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test get link
	getCustomLink, err := testStore.GetCustomLinkByToken(context.Background(), customLink.Token)
	assert.NoError(t, err)
	assert.NotNil(t, getCustomLink)
	assert.Equal(t, *getCustomLink, customLink)
	assert.Equal(t, customLink.Id, getCustomLink.Id)
	assert.Equal(t, customLink.Url, getCustomLink.Url)
	assert.Equal(t, customLink.Token, getCustomLink.Token)
	assert.Equal(t, customLink.Deleted, getCustomLink.Deleted)
	assert.Equal(t, customLink.IsSuggestion, getCustomLink.IsSuggestion)
	assert.Equal(t, customLink.SuggestionId, getCustomLink.SuggestionId)
	assert.Equal(t, customLink.CreatedAt, getCustomLink.CreatedAt)
	assert.Equal(t, customLink.UpdatedAt, getCustomLink.UpdatedAt)
}

func TestListAllCustomLink(t *testing.T) {

	for i := 0; i < listTestLen; i++ {
		_ = createRandomCustomLink(t, "")
	}

	customLinkList, err := testStore.ListAllCustomLink(context.Background())
	assert.NoError(t, err)
	// The len should be <= listTestLen
	assert.LessOrEqual(t, listTestLen, len(customLinkList))

	for _, v := range customLinkList {
		assert.NotEmpty(t, v)
	}
}

func TestListActiveCustomLink(t *testing.T) {

	for i := 0; i < listTestLen; i++ {
		_ = createRandomCustomLink(t, "", false)
		_ = createRandomCustomLink(t, "", true)
	}

	//test dont get deleted links
	customLinkList, err := testStore.ListAllCustomLink(context.Background())
	assert.NoError(t, err)
	activeLinksList, err := testStore.ListActiveCustomLink(context.Background())
	assert.NoError(t, err)
	assert.NotEqual(t, len(customLinkList), len(activeLinksList))

	// The filed deleted must be false
	for _, v := range activeLinksList {
		assert.False(t, v.Deleted)
	}
}

func TestListCustomLinkByUser(t *testing.T) {
	//TODO add this test when implement user repository
}

func TestDeleteCustomLink(t *testing.T) {
	customLink := createRandomCustomLink(t, "")
	deleteToken := model.CustomLinkToken(repository.DeleteCustomTokenValue + strconv.Itoa(int(customLink.Id)))
	// Test delete link
	deletedLink, err := testStore.DeleteCustomLink(context.Background(), customLink.Id)
	assert.NotNil(t, deletedLink)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedLink)
	assert.Equal(t, deletedLink.Url, repository.DeleteStringValue)
	assert.Equal(t, deletedLink.Token, deleteToken)
	assert.Equal(t, customLink.IsSuggestion, deletedLink.IsSuggestion)
	assert.Equal(t, customLink.SuggestionId, deletedLink.SuggestionId)
	assert.True(t, deletedLink.Deleted)
	assert.Greater(t, deletedLink.UpdatedAt, deletedLink.CreatedAt)

	// Test Not Found error
	errorLink, err := testStore.DeleteCustomLink(context.Background(), model.CustomLinkId(util.RandomInt(maxRandomSk, maxRandomSk*2)))
	assert.Empty(t, errorLink)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
}

func TestDeleteCustomLinkByToken(t *testing.T) {
	customLink := createRandomCustomLink(t, "")
	deleteToken := model.CustomLinkToken(repository.DeleteCustomTokenValue + strconv.Itoa(int(customLink.Id)))
	// Test delete link
	deletedLink, err := testStore.DeleteCustomLinkByToken(context.Background(), customLink.Token)
	assert.NotNil(t, deletedLink)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedLink)
	assert.Equal(t, deletedLink.Url, repository.DeleteStringValue)
	assert.Equal(t, deletedLink.Token, deleteToken)
	assert.Equal(t, customLink.IsSuggestion, deletedLink.IsSuggestion)
	assert.Equal(t, customLink.SuggestionId, deletedLink.SuggestionId)
	assert.True(t, deletedLink.Deleted)
	assert.Greater(t, deletedLink.UpdatedAt, deletedLink.CreatedAt)

	// Test Not Found error
	errorLink, err := testStore.DeleteCustomLinkByToken(context.Background(), model.CustomLinkToken(util.RandomString(minRandomSk)))
	assert.Empty(t, errorLink)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
}
