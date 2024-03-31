package memory

import (
	"context"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
)

// PutCustomLink adds a new link .
func (r *Repository) PutCustomLink(ctx context.Context, customLink model.CustomLink) error {
	if _, ok := r.customLinkData[customLink.Id]; !ok {
		r.customLinkData[customLink.Id] = customLink
	}

	if _, ok := r.customLinkTokenData[customLink.Token]; !ok {
		r.customLinkTokenData[customLink.Token] = customLink
	}

	return nil
}

// GetCustomLinkByID retrieves a custom link by Id.
func (r *Repository) GetCustomLinkByID(ctx context.Context, customLinkID model.CustomLinkId) (*model.CustomLink, error) {
	if _, ok := r.customLinkData[customLinkID]; !ok {
		return nil, repository.ErrNotFound
	}

	customLinkValue := r.customLinkData[customLinkID]

	return &customLinkValue, nil
}

// GetCustomLinkByToken retrieves a custom link by token.
func (r *Repository) GetCustomLinkByToken(ctx context.Context, customLinkToken model.CustomLinkToken) (*model.CustomLink, error) {
	if _, ok := r.customLinkTokenData[customLinkToken]; !ok {
		return nil, repository.ErrNotFound
	}

	customLinkValue := r.customLinkTokenData[customLinkToken]

	return &customLinkValue, nil
}

// ListAllCustomLink retrieves all Custom Links.
func (r *Repository) ListAllCustomLink(ctx context.Context) ([]model.CustomLink, error) {
	customLinks := make([]model.CustomLink, 0, len(r.customLinkData))

	for _, link := range r.customLinkData {
		customLinks = append(customLinks, link)
	}

	return customLinks, nil
}

// ListActiveCustomLink retreives all Custom Likns not deleted.
func (r *Repository) ListActiveCustomLink(ctx context.Context) ([]model.CustomLink, error) {
	customLinks := make([]model.CustomLink, 0, len(r.customLinkData))

	for _, link := range r.customLinkData {
		if !link.Deleted {
			customLinks = append(customLinks, link)
		}
	}

	return customLinks, nil
}

// ListCustomLinkByUser retreives all Custom Links by User.
func (r *Repository) ListCustomLinkByUser(ctx context.Context, userId uuid.UUID) ([]model.CustomLink, error) {
	customLinks := make([]model.CustomLink, 0, len(r.customLinkData))

	for _, link := range r.customLinkData {
		if userId == link.UserId && !link.Deleted {
			customLinks = append(customLinks, link)
		}
	}

	return customLinks, nil
}

// DeleteCustomLink logic delete for a custom link.
func (r *Repository) DeleteCustomLink(ctx context.Context, customLinkID model.CustomLinkId) (*model.CustomLink, error) {
	if _, ok := r.customLinkData[customLinkID]; !ok {
		return nil, repository.ErrNotFound
	}

	token := r.customLinkData[customLinkID].Token

	newCustomLink := model.CustomLink{
		Id:           customLinkID,
		UserId:       r.customLinkData[customLinkID].UserId,
		Url:          repository.DeleteStringValue,
		Token:        "",
		IsSuggestion: r.customLinkData[customLinkID].IsSuggestion,
		SuggestionId: r.customLinkData[customLinkID].SuggestionId,
		Deleted:      true,
		CreatedAt:    r.customLinkData[customLinkID].CreatedAt,
		UpdatedAt:    time.Now(),
	}

	r.customLinkData[customLinkID] = newCustomLink

	delete(r.customLinkTokenData, token)

	return &newCustomLink, nil
}
