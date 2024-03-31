package memory

import (
	"context"
	"testing"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/stretchr/testify/assert"
)

const (
	recycleLinkLenMin = 1
	recycleLinkLenMax = 100000
)

// createRandomCustomLink create a random custom link for test propouses
func createRandomRecycleLink(t *testing.T, repo *Repository) model.RecycleLink {
	recycleLink := model.RecycleLink{
		SKey: model.RecycleLinkId(util.RandomInt(recycleLinkLenMin, recycleLinkLenMax)),
	}

	newRecycleLink := repo.PutRecycleLink(context.Background(), recycleLink)

	assert.NoError(t, newRecycleLink)

	return recycleLink
}

func TestPutRecycleLink(t *testing.T) {
	repo := New()
	recycleLink := createRandomRecycleLink(t, repo)

	assert.NotEmpty(t, recycleLink)
}

func TestGetRecycleLink(t *testing.T) {
	repo := New()
	recycleLink := createRandomRecycleLink(t, repo)

	// Test not found error
	getRecycleLinkError, err := repo.GetRecycleLink(context.Background(), model.RecycleLinkId(util.RandomInt(recycleLinkLenMax+1, recycleLinkLenMax+50)))
	assert.Nil(t, getRecycleLinkError)
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound)

	// Test get link
	getRecycleLink, err := repo.GetRecycleLink(context.Background(), recycleLink.SKey)
	assert.NotNil(t, getRecycleLink)
	assert.Equal(t, *getRecycleLink, recycleLink)
	assert.NoError(t, err)
}

func TestDeleteRecycleLink(t *testing.T) {
	repo := New()
	recycleLink := createRandomRecycleLink(t, repo)

	// Test Not Found error
	err := repo.DeleteRecycleLink(context.Background(), model.RecycleLinkId(util.RandomInt(recycleLinkLenMax+1, recycleLinkLenMax+50)))
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound)

	// Test delete
	err = repo.DeleteRecycleLink(context.Background(), recycleLink.SKey)
	assert.NoError(t, err)

	deletedLink, err := repo.GetRecycleLink(context.Background(), recycleLink.SKey)
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound)
	assert.Nil(t, deletedLink)
}
