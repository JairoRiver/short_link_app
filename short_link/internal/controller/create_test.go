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
	assert.Equal(t, newShortLink.Url, url)

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
	assert.Equal(t, otherShortLink.Url, url)
	assert.Equal(t, otherShortLink.Token, util.ToBase62(uint64(testSK)))
}

func TestCreateCustomLink(t *testing.T) {
	repo := memory.New()
	control := New(repo)

	//test invalid token len
	url := util.RandomURL(11)
	user := repository.HasUserID{}
	errToken := util.RandomString(6)
	customToken, err := control.CreateCustomLink(context.Background(), url, user, errToken)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrInvalidCustomToken))
	assert.Nil(t, customToken)

	//test valid token len
	url_2 := util.RandomURL(11)
	user_2 := repository.HasUserID{}
	token := util.RandomString(10)
	customToken_2, err := control.CreateCustomLink(context.Background(), url_2, user_2, token)

	assert.NoError(t, err)
	assert.NotNil(t, customToken_2)
	assert.Equal(t, customToken_2.Token, token)
	assert.Equal(t, customToken_2.Url, url_2)
}
