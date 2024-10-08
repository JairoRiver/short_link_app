package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
)

// PutCustomLink adds a new link .
func (r *Repository) PutCustomLink(ctx context.Context, customLinkParams repository.CreateCustomLinkParams) (model.CustomLink, error) {

	customLink := model.CustomLink{
		Id:           model.CustomLinkId(r.lastIds["customLink"] + 1),
		Url:          customLinkParams.Url,
		Token:        customLinkParams.Token,
		IsSuggestion: customLinkParams.IsSuggestion,
		SuggestionId: customLinkParams.SuggestionId,
		Deleted:      customLinkParams.Deleted,
		CreatedAt:    customLinkParams.CreatedAt,
		UpdatedAt:    customLinkParams.UpdatedAt,
	}

	r.lastIds["customLink"] = r.lastIds["customLink"] + 1

	if customLinkParams.UserId.Valid {
		customLink.UserId = customLinkParams.UserId.ID
	}

	r.customLinkData[customLink.Id] = customLink

	if _, ok := r.customLinkTokenData[customLink.Token]; !ok {
		r.customLinkTokenData[customLink.Token] = customLink
	}

	return customLink, nil
}

// GetCustomLinkByID retrieves a custom link by Id.
func (r *Repository) GetCustomLinkByID(ctx context.Context, customLinkID model.CustomLinkId) (*model.CustomLink, error) {
	if _, ok := r.customLinkData[customLinkID]; !ok {
		return nil, fmt.Errorf("Repository memory GetCustomLinkByID method error: %w", repository.ErrNotFound)
	}

	customLinkValue := r.customLinkData[customLinkID]

	return &customLinkValue, nil
}

// GetCustomLinkByToken retrieves a custom link by token.
func (r *Repository) GetCustomLinkByToken(ctx context.Context, customLinkToken model.CustomLinkToken) (*model.CustomLink, error) {
	if _, ok := r.customLinkTokenData[customLinkToken]; !ok {
		return nil, fmt.Errorf("Repository memory GetCustomLinkByToken method error: %w", repository.ErrNotFound)
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

// ListActiveCustomLink retrieves all Custom Links not deleted.
func (r *Repository) ListActiveCustomLink(ctx context.Context) ([]model.CustomLink, error) {
	customLinks := make([]model.CustomLink, 0, len(r.customLinkData))

	for _, link := range r.customLinkData {
		if !link.Deleted {
			customLinks = append(customLinks, link)
		}
	}

	return customLinks, nil
}

// ListCustomLinkByUser retrieves all Custom Links by User.
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
		return nil, fmt.Errorf("Repository memory DeleteCustomLink method error: %w", repository.ErrNotFound)
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

// DeleteCustomLinkByToken logic delete for a custom link by token.
func (r *Repository) DeleteCustomLinkByToken(ctx context.Context, customLinkToken model.CustomLinkToken) (*model.CustomLink, error) {
	if _, ok := r.customLinkTokenData[customLinkToken]; !ok {
		return nil, fmt.Errorf("Repository memory DeleteCustomLinkByToken method error: %w", repository.ErrNotFound)
	}

	id := r.customLinkTokenData[customLinkToken].Id

	newCustomLink := model.CustomLink{
		Id:           id,
		UserId:       r.customLinkData[id].UserId,
		Url:          repository.DeleteStringValue,
		Token:        "",
		IsSuggestion: r.customLinkData[id].IsSuggestion,
		SuggestionId: r.customLinkData[id].SuggestionId,
		Deleted:      true,
		CreatedAt:    r.customLinkData[id].CreatedAt,
		UpdatedAt:    time.Now(),
	}

	r.customLinkData[id] = newCustomLink

	delete(r.customLinkTokenData, customLinkToken)

	return &newCustomLink, nil
}
