package postgres

import (
	"context"
	"errors"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRandomRecycleLink(t *testing.T) {
	arg := model.RecycleLink{
		SKey: model.RecycleLinkId(util.RandomInt(minRandomSk, maxRandomSk)),
	}
	err := testStore.PutRecycleLink(context.Background(), arg)
	assert.NoError(t, err)

}

func emptyRecycleLink() {
	getRecycleLink, err := testStore.GetRecycleLink(context.Background())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return
		}
		log.Fatal().Err(err).Msg("failed to get recycle link")
	}
	err = testStore.DeleteRecycleLink(context.Background(), getRecycleLink.SKey)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return
		}
		log.Fatal().Err(err).Msg("failed to get recycle link")
	}

	emptyRecycleLink()
}

// createRandomShortLink create short link for test proposes
func TestPutRecycleLink(t *testing.T) {
	createRandomRecycleLink(t)
	emptyRecycleLink()
}

func TestGetRecycleLink(t *testing.T) {
	// Test not found error
	getRecycleLinkError, err := testStore.GetRecycleLink(context.Background())
	assert.Empty(t, getRecycleLinkError)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test get link
	createRandomRecycleLink(t)
	getRecycleLink, err := testStore.GetRecycleLink(context.Background())
	assert.NotNil(t, getRecycleLink)
	assert.NotEmpty(t, getRecycleLink.SKey)
	assert.NoError(t, err)

	emptyRecycleLink()
}

func TestDeleteRecycleLink(t *testing.T) {
	// Test not found error
	sKey := model.RecycleLinkId(util.RandomInt(minRandomSk, maxRandomSk))
	err := testStore.DeleteRecycleLink(context.Background(), sKey)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNotFound))

	// Test delete link
	createRandomRecycleLink(t)
	getRecycleLink, err := testStore.GetRecycleLink(context.Background())
	assert.NotNil(t, getRecycleLink)
	assert.NotEmpty(t, getRecycleLink.SKey)
	assert.NoError(t, err)

	err = testStore.DeleteRecycleLink(context.Background(), getRecycleLink.SKey)
	assert.NoError(t, err)
}
