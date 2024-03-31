package memory

import (
	"context"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

// PutRecycleLink adds a new recycle link.
func (r *Repository) PutRecycleLink(ctx context.Context, recycleLink model.RecycleLink) error {
	if _, ok := r.recycleLinkData[recycleLink.SKey]; !ok {
		r.recycleLinkData[recycleLink.SKey] = recycleLink
	}

	return nil
}

// GetRecycleLink retrieves a recycle link by SKey.
func (r *Repository) GetRecycleLink(ctx context.Context, recycleLinkID model.RecycleLinkId) (*model.RecycleLink, error) {
	if _, ok := r.recycleLinkData[recycleLinkID]; !ok {
		return nil, repository.ErrNotFound
	}

	recycleLinkValue := r.recycleLinkData[recycleLinkID]

	return &recycleLinkValue, nil
}

// DeleteRecycleLink delete a recycle link.
func (r *Repository) DeleteRecycleLink(ctx context.Context, recycleLinkID model.RecycleLinkId) error {
	if _, ok := r.recycleLinkData[recycleLinkID]; !ok {
		return repository.ErrNotFound
	}

	delete(r.recycleLinkData, recycleLinkID)

	return nil
}
