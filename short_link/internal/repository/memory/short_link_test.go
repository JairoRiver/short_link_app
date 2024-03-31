package memory

import (
	"context"
	"testing"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	shortLinkLenMin = 1
	shortLinkLenMax = 10000
	linkURLLen      = 20
	linkTokenLen    = 6
	sKeyLenMin      = 10001
	sKeyLenMax      = 100000
	listTestLen     = 20
)

// createRandomShortLink create short link for test propouses
func createRandomShortLink(t *testing.T, repo *Repository) model.ShortLink {
	shortLink := model.ShortLink{
		Id:        model.ShortLinkId(util.RandomInt(shortLinkLenMin, shortLinkLenMax)),
		UserId:    uuid.New(),
		Url:       util.RandomURL(linkURLLen),
		Token:     util.RandomString(linkTokenLen),
		SKey:      model.ShortLinkId(util.RandomInt(sKeyLenMin, sKeyLenMax)),
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newShortLink := repo.PutShortLink(context.Background(), shortLink)

	assert.NoError(t, newShortLink)

	return shortLink
}

func TestPutShortLink(t *testing.T) {
	repo := New()
	shortLink := createRandomShortLink(t, repo)

	assert.NotEmpty(t, shortLink)
}

func TestGetShortLinkByID(t *testing.T) {
	repo := New()
	shortLink := createRandomShortLink(t, repo)

	// Test not found error
	getShortLinkError, err := repo.GetShortLinkByID(context.Background(), model.ShortLinkId(util.RandomInt(shortLinkLenMax+1, shortLinkLenMax+50)))
	assert.Nil(t, getShortLinkError)
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound)

	// Test get link
	getShortLink, err := repo.GetShortLinkByID(context.Background(), shortLink.Id)
	assert.NotNil(t, getShortLink)
	assert.Equal(t, *getShortLink, shortLink)
	assert.NoError(t, err)
}

func TestGetShortLinkBySKey(t *testing.T) {
	repo := New()
	shortLink := createRandomShortLink(t, repo)

	// Test not found error
	getShortLinkError, err := repo.GetShortLinkBySKey(context.Background(), model.ShortLinkId(util.RandomInt(sKeyLenMax+1, sKeyLenMax+50)))
	assert.Nil(t, getShortLinkError)
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound)

	// Test get link
	getShortLink, err := repo.GetShortLinkBySKey(context.Background(), shortLink.SKey)
	assert.NotNil(t, getShortLink)
	assert.Equal(t, *getShortLink, shortLink)
	assert.NoError(t, err)
}

func TestListAllShortLink(t *testing.T) {
	repo := New()
	for i := 0; i < listTestLen; i++ {
		_ = createRandomShortLink(t, repo)
	}

	shotLinkList, err := repo.ListAllShortLink(context.Background())
	assert.NoError(t, err)
	// The len shoud by the double beacause by short link in memory model we created two entries, by id and by sKey
	assert.Len(t, shotLinkList, listTestLen*2)
}

func TestListActiveShortLink(t *testing.T) {
	repo := New()
	for i := 0; i < listTestLen; i++ {
		_ = createRandomShortLink(t, repo)
	}

	shortLinkList, err := repo.ListActiveShortLink(context.Background())
	assert.NoError(t, err)

	// All elements must active, Deleted == false
	for _, link := range shortLinkList {
		assert.False(t, link.Deleted)
	}
}

func TestListShortLinkByUser(t *testing.T) {
	repo := New()
	linkList := make([]model.ShortLink, 0, listTestLen)
	for i := 0; i < listTestLen; i++ {
		linkList = append(linkList, createRandomShortLink(t, repo))
	}

	for _, shortLink := range linkList {
		getShortLinks, err := repo.ListShortLinkByUser(context.Background(), shortLink.UserId)
		assert.NoError(t, err)
		for _, testLink := range getShortLinks {
			assert.Equal(t, shortLink.UserId, testLink.UserId)
			assert.False(t, testLink.Deleted)
		}
	}
}

func TestDeleteShortLink(t *testing.T) {
	repo := New()
	shortLink := createRandomShortLink(t, repo)

	// Test Not Found error
	errorLink, err := repo.DeleteShortLink(context.Background(), model.ShortLinkId(util.RandomInt(shortLinkLenMax+1, shortLinkLenMax+50)))
	assert.Nil(t, errorLink)
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound)

	// Test logical delete
	deletedLink, err := repo.DeleteShortLink(context.Background(), shortLink.Id)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedLink)
	assert.Equal(t, deletedLink.Url, repository.DeleteStringValue)
	assert.Zero(t, deletedLink.Token)
	assert.Zero(t, deletedLink.SKey)
	assert.True(t, deletedLink.Deleted)
	assert.Greater(t, deletedLink.UpdatedAt, deletedLink.CreatedAt)

	deletedLinkSk, err := repo.GetShortLinkBySKey(context.Background(), shortLink.SKey)
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound)
	assert.Nil(t, deletedLinkSk)
}
