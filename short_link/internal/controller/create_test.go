package controller

import (
	"context"
	"math/big"
	"testing"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/memory"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateShortLink(t *testing.T) {
	repo := memory.New()
	control := New(repo)

	var testSK int64 = 1000

	// Test whitout recicle link
	url := util.RandomURL(9)
	user := repository.HasUserID{ID: uuid.New(), Valid: true}

	newShortLink, err := control.CreateShortLink(context.Background(), url, user)

	assert.NoError(t, err)
	assert.NotNil(t, newShortLink)
	assert.Equal(t, newShortLink.url, url)

	// Test with recicle link
	testRLink := model.RecycleLink{
		SKey: model.RecycleLinkId(testSK),
	}
	repo.PutRecycleLink(context.Background(), testRLink)
	url = util.RandomURL(9)
	user = repository.HasUserID{}

	otherShortLink, err := control.CreateShortLink(context.Background(), url, user)

	assert.NoError(t, err)
	assert.NotNil(t, otherShortLink)
	assert.Equal(t, otherShortLink.url, url)
	assert.Equal(t, otherShortLink.token, big.NewInt(testSK).Text(62))
}
