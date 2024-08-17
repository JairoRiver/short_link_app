package postgres

import (
	"context"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuxSKey(t *testing.T) {
	// Test get aux sKey
	auxSkey, err := testStore.GetAuxSKey(context.Background())
	assert.NotNil(t, auxSkey)
	assert.Equal(t, auxSkey.N, uint(util.SknValue))
	assert.Equal(t, auxSkey.Step, uint(util.SkstepValue))
	assert.Equal(t, auxSkey.End, uint(util.SkendValue))
	assert.NoError(t, err)
}

func TestUpdateAuxSKey(t *testing.T) {
	// Test update skAux
	params := repository.AuxSKeyParams{
		A0: repository.IsIntValid{Valid: true, Value: uint(util.RandomInt(minRandomSk, maxRandomSk))},
		N0: repository.IsIntValid{Valid: true, Value: uint(util.RandomInt(minRandomSk, maxRandomSk))},
	}
	deletedLink, err := testStore.UpdateAuxSKey(context.Background(), params)
	assert.NotNil(t, deletedLink)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedLink)
	assert.Equal(t, deletedLink.A0, params.A0.Value)
	assert.Equal(t, deletedLink.N0, params.N0.Value)
}
