package memory

import (
	"context"
	"testing"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	customLinkLenMin   = 1
	customLinkLenMax   = 100000
	linkCustomTokenLen = 8
)

// createRandomCustomLink create a random custom link for test propouses
func createRandomCustomLink(t *testing.T, repo *Repository) model.CustomLink {
	customLink := repository.CreateCustomLinkParams{
		UserId:       repository.HasUserID{ID: uuid.New(), Valid: true},
		Url:          util.RandomURL(linkURLLen),
		Token:        model.CustomLinkToken(util.RandomString(linkCustomTokenLen)),
		IsSuggestion: false,
		SuggestionId: 0,
		Deleted:      false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	newCustomLink, err := repo.PutCustomLink(context.Background(), customLink)

	assert.NoError(t, err)

	return newCustomLink
}

func TestPutCustomLink(t *testing.T) {
	repo := New()
	customLink := createRandomCustomLink(t, repo)

	assert.NotEmpty(t, customLink)
}

func TestGetCustomLinkByID(t *testing.T) {
	repo := New()
	customLink := createRandomCustomLink(t, repo)

	// Test not found error
	getCustomLinkError, err := repo.GetCustomLinkByID(context.Background(), model.CustomLinkId(util.RandomInt(customLinkLenMax+1, customLinkLenMax+50)))
	assert.Nil(t, getCustomLinkError)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test get link
	getCustomLink, err := repo.GetCustomLinkByID(context.Background(), customLink.Id)
	assert.NotNil(t, getCustomLink)
	assert.Equal(t, *getCustomLink, customLink)
	assert.NoError(t, err)
}

func TestGetCustomLinkByToken(t *testing.T) {
	repo := New()
	customLink := createRandomCustomLink(t, repo)

	// Test not found error
	getCustomLinkError, err := repo.GetCustomLinkByToken(context.Background(), model.CustomLinkToken(util.RandomString(linkCustomTokenLen)))
	assert.Nil(t, getCustomLinkError)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test get link
	getCustomLink, err := repo.GetCustomLinkByToken(context.Background(), customLink.Token)
	assert.NotNil(t, getCustomLink)
	assert.Equal(t, *getCustomLink, customLink)
	assert.NoError(t, err)
}

func TestListAllCustomLink(t *testing.T) {
	repo := New()
	for i := 0; i < listTestLen; i++ {
		_ = createRandomCustomLink(t, repo)
	}

	customLinkList, err := repo.ListAllCustomLink(context.Background())
	assert.NoError(t, err)
	assert.Len(t, customLinkList, listTestLen)
}

func TestListActiveCustomLink(t *testing.T) {
	repo := New()
	for i := 0; i < listTestLen; i++ {
		_ = createRandomCustomLink(t, repo)
	}

	customLinkList, err := repo.ListActiveCustomLink(context.Background())
	assert.NoError(t, err)

	// All elements must active, Deleted == false
	for _, link := range customLinkList {
		assert.False(t, link.Deleted)
	}
}

func TestListCustomLinkByUser(t *testing.T) {
	repo := New()
	linkList := make([]model.CustomLink, 0, listTestLen)
	for i := 0; i < listTestLen; i++ {
		linkList = append(linkList, createRandomCustomLink(t, repo))
	}

	for _, customLink := range linkList {
		getCustomLinks, err := repo.ListCustomLinkByUser(context.Background(), customLink.UserId)
		assert.NoError(t, err)
		for _, testLink := range getCustomLinks {
			assert.Equal(t, customLink.UserId, testLink.UserId)
			assert.False(t, testLink.Deleted)
		}
	}
}

func TestDeleteCustomLink(t *testing.T) {
	repo := New()
	customLink := createRandomCustomLink(t, repo)

	// Test Not Found error
	errorLink, err := repo.DeleteCustomLink(context.Background(), model.CustomLinkId(util.RandomInt(customLinkLenMax+1, customLinkLenMax+50)))
	assert.Nil(t, errorLink)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test logical delete
	deletedLink, err := repo.DeleteCustomLink(context.Background(), customLink.Id)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedLink)
	assert.Equal(t, deletedLink.Url, repository.DeleteStringValue)
	assert.Zero(t, deletedLink.Token)
	assert.True(t, deletedLink.Deleted)
	assert.Greater(t, deletedLink.UpdatedAt, deletedLink.CreatedAt)

	deletedLinkToken, err := repo.GetCustomLinkByToken(context.Background(), customLink.Token)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
	assert.Nil(t, deletedLinkToken)
}
