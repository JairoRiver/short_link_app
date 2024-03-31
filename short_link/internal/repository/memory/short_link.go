package memory

import (
	"context"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
)

// PutShortLink adds a new link .
func (r *Repository) PutShortLink(ctx context.Context, shotLink model.ShortLink) error {
	if _, ok := r.shortLinkData[shotLink.Id]; !ok {
		r.shortLinkData[shotLink.Id] = shotLink
	}

	if _, ok := r.shortLinkData[shotLink.SKey]; !ok {
		r.shortLinkData[shotLink.SKey] = shotLink
	}

	return nil
}

// GetShortLinkByID retrieves a shot link by Id.
func (r *Repository) GetShortLinkByID(ctx context.Context, shotLinkID model.ShortLinkId) (*model.ShortLink, error) {
	if _, ok := r.shortLinkData[shotLinkID]; !ok {
		return nil, repository.ErrNotFound
	}

	shortLinkValue := r.shortLinkData[shotLinkID]

	return &shortLinkValue, nil
}

// GetShortLinkBySKey retrieves a shot link by S_key.
func (r *Repository) GetShortLinkBySKey(ctx context.Context, sKeyID model.ShortLinkId) (*model.ShortLink, error) {
	if _, ok := r.shortLinkData[sKeyID]; !ok {
		return nil, repository.ErrNotFound
	}

	shortLinkValue := r.shortLinkData[sKeyID]

	return &shortLinkValue, nil
}

// ListAllShortLink retrieves all Short Links.
func (r *Repository) ListAllShortLink(ctx context.Context) ([]model.ShortLink, error) {
	shortLinks := make([]model.ShortLink, 0, len(r.shortLinkData))

	for _, link := range r.shortLinkData {
		shortLinks = append(shortLinks, link)
	}

	return shortLinks, nil
}

// ListActiveShortLink retreives all Short Likns not deleted.
func (r *Repository) ListActiveShortLink(ctx context.Context) ([]model.ShortLink, error) {
	shortLinks := make([]model.ShortLink, 0, len(r.shortLinkData))

	for _, link := range r.shortLinkData {
		if !link.Deleted {
			shortLinks = append(shortLinks, link)
		}
	}

	return shortLinks, nil
}

// ListShortLinkByUser retreives all Short Links by User.
func (r *Repository) ListShortLinkByUser(ctx context.Context, userId uuid.UUID) ([]model.ShortLink, error) {
	shortLinks := make([]model.ShortLink, 0, len(r.shortLinkData))

	for _, link := range r.shortLinkData {
		if userId == link.UserId && !link.Deleted {
			shortLinks = append(shortLinks, link)
		}
	}

	return shortLinks, nil
}

// DeleteShortLink logic delete for a short link.
func (r *Repository) DeleteShortLink(ctx context.Context, shotLinkID model.ShortLinkId) (*model.ShortLink, error) {
	if _, ok := r.shortLinkData[shotLinkID]; !ok {
		return nil, repository.ErrNotFound
	}

	sKey := r.shortLinkData[shotLinkID].SKey

	newShortLink := model.ShortLink{
		Id:        shotLinkID,
		UserId:    r.shortLinkData[shotLinkID].UserId,
		Url:       repository.DeleteStringValue,
		Token:     "",
		SKey:      0,
		Deleted:   true,
		CreatedAt: r.shortLinkData[shotLinkID].CreatedAt,
		UpdatedAt: time.Now(),
	}

	r.shortLinkData[shotLinkID] = newShortLink

	delete(r.shortLinkData, sKey)

	return &newShortLink, nil
}
