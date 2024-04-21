package memory

import (
	"context"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

// Aux func to find the position of a recycle key
func (r *Repository) recycleKeyIndex(value model.RecycleLinkId) int {
	for i, v := range r.recycleLinkData {
		if v.SKey == value {
			return i
		}
	}
	return -1
}

// Aux func to delete element of recycle link slice from one index
func (r *Repository) deleteElementRecycleLink(index int) []model.RecycleLink {
	newSlice := append(r.recycleLinkData[:index], r.recycleLinkData[index+1:]...)
	return newSlice
}

// PutRecycleLink adds a new recycle link.
func (r *Repository) PutRecycleLink(ctx context.Context, recycleLink model.RecycleLink) error {
	r.recycleLinkData = append(r.recycleLinkData, recycleLink)

	return nil
}

// GetRecycleLink retrieves a recycle link by SKey.
func (r *Repository) GetRecycleLink(ctx context.Context) (*model.RecycleLink, error) {
	if len(r.recycleLinkData) == 0 {
		return nil, repository.ErrNotFound
	}

	recycleLinkValue := r.recycleLinkData[0]

	return &recycleLinkValue, nil
}

// DeleteRecycleLink delete a recycle link.
func (r *Repository) DeleteRecycleLink(ctx context.Context, recycleLinkID model.RecycleLinkId) error {
	recycleLinkIndex := r.recycleKeyIndex(recycleLinkID)
	if recycleLinkIndex == -1 {
		return repository.ErrNotFound
	}

	newRecycleLinkSlice := r.deleteElementRecycleLink(recycleLinkIndex)
	r.recycleLinkData = newRecycleLinkSlice

	return nil
}
