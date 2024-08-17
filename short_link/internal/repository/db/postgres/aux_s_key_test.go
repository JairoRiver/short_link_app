package postgres

import (
	"context"
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
