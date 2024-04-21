package memory

import (
	"context"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

// GetAuxSKey retrieves data to generate S Keys.
func (r *Repository) GetAuxSKey(ctx context.Context) (*model.AuxSKey, error) {

	auxSKeyData := r.auxSKey

	return &auxSKeyData, nil
}

// UpdateAuxSKey retrieves data to generate S Keys.
func (r *Repository) UpdateAuxSKey(ctx context.Context, params repository.AuxSKeyParams) (*model.AuxSKey, error) {

	// Validate if need update N0
	if params.N0.Valid {
		r.auxSKey.N0 = uint8(params.N0.Value)
	}

	// Validate if nedd update A0
	if params.A0.Valid {
		r.auxSKey.A0 = params.A0.Value
	}
	auxSKeyData := r.auxSKey

	return &auxSKeyData, nil
}
