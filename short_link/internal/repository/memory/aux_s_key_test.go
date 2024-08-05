package memory

import (
	"context"
	"testing"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/stretchr/testify/assert"
)

const (
	nValue    = util.SknValue
	initValue = util.SkinitValue
	endValue  = util.SkendValue
	stepValue = util.SkstepValue
	a0Value   = util.Ska0Value
	n0Value   = util.Skn0Value
)

func TestGetAuxSKey(t *testing.T) {
	repo := New()
	auxSKey, err := repo.GetAuxSKey(context.Background())

	// Test not error
	assert.NoError(t, err)
	assert.NotEmpty(t, auxSKey)

	// Test the init value
	assert.Equal(t, auxSKey.N, uint(nValue))
	assert.Equal(t, auxSKey.End, uint(endValue))
	assert.Equal(t, auxSKey.Step, uint(stepValue))
	assert.Equal(t, auxSKey.A0, uint(a0Value))
	assert.Equal(t, auxSKey.N0, uint(n0Value))
}

func TestUpdateAuxSKey(t *testing.T) {
	repo := New()

	// Test to update N0
	newN0 := util.RandomInt(1, nValue)
	N0UpdateParams := repository.AuxSKeyParams{
		N0: repository.IsIntValid{Value: uint(newN0), Valid: true},
	}
	newN0SKey, err := repo.UpdateAuxSKey(context.Background(), N0UpdateParams)

	assert.NoError(t, err)
	assert.Equal(t, newN0SKey.N0, uint(newN0))

	// Test to update A0
	newA0 := util.RandomInt(1, stepValue)
	A0UpdateParams := repository.AuxSKeyParams{
		A0: repository.IsIntValid{Value: uint(newA0), Valid: true},
	}
	newA0SKey, err := repo.UpdateAuxSKey(context.Background(), A0UpdateParams)
	assert.NoError(t, err)
	assert.Equal(t, newA0SKey.A0, uint(newA0))
}
