package postgres

import (
	"context"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"strconv"
	"time"
)

// PutCustomLink adds a new link .
const createCustomLinkQuery = `
INSERT INTO custom_link (
                        user_id,
                        url,
                        token,
                        is_suggestion,
                        suggestion_id,
                        deleted
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at
`

func (q *Queries) PutCustomLink(ctx context.Context, customLinkParams repository.CreateCustomLinkParams) (model.CustomLink, error) {
	var userId *uuid.UUID
	if customLinkParams.UserId.Valid {
		userId = &customLinkParams.UserId.ID
	}
	row := q.db.QueryRow(ctx, createCustomLinkQuery, userId, customLinkParams.Url, customLinkParams.Token, customLinkParams.IsSuggestion, customLinkParams.SuggestionId, customLinkParams.Deleted)
	var i model.CustomLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.IsSuggestion,
		&i.SuggestionId,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

// GetCustomLinkByID retrieves a custom link by Id.
const getCustomLinkByIdQuery = `
SELECT id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at 
FROM custom_link
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCustomLinkByID(ctx context.Context, customLinkID model.CustomLinkId) (*model.CustomLink, error) {
	row := q.db.QueryRow(ctx, getCustomLinkByIdQuery, customLinkID)
	var i model.CustomLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.IsSuggestion,
		&i.SuggestionId,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres GetCustomLinkByID method error: %w", repository.ErrNotFound)
	}

	return &i, err
}

// GetCustomLinkByToken retrieves a custom link by token.
const getCustomLinkByTokenQuery = `
SELECT id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at
FROM custom_link
WHERE token = $1
LIMIT 1
`

func (q *Queries) GetCustomLinkByToken(ctx context.Context, customLinkToken model.CustomLinkToken) (*model.CustomLink, error) {
	row := q.db.QueryRow(ctx, getCustomLinkByTokenQuery, customLinkToken)
	var i model.CustomLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.IsSuggestion,
		&i.SuggestionId,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres GetCustomLinkByToken method error: %w", repository.ErrNotFound)
	}

	return &i, err
}

// ListAllCustomLink retrieves all Custom Links.
const listAllCustomLinkQuery = `
SELECT id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at 
FROM custom_link
`

func (q *Queries) ListAllCustomLink(ctx context.Context) ([]model.CustomLink, error) {
	rows, err := q.db.Query(ctx, listAllCustomLinkQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.CustomLink
	for rows.Next() {
		var i model.CustomLink
		if err := rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Url,
			&i.Token,
			&i.IsSuggestion,
			&i.SuggestionId,
			&i.Deleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// ListActiveCustomLink retrieves all Custom Links not deleted.
const listActiveCustomLinkQuery = `
SELECT id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at 
FROM custom_link
WHERE deleted IS false
`

func (q *Queries) ListActiveCustomLink(ctx context.Context) ([]model.CustomLink, error) {
	rows, err := q.db.Query(ctx, listActiveCustomLinkQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.CustomLink
	for rows.Next() {
		var i model.CustomLink
		if err := rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Url,
			&i.Token,
			&i.IsSuggestion,
			&i.SuggestionId,
			&i.Deleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// ListCustomLinkByUser retrieves all Custom Links by User.
const listCustomLinkByUserQuery = `
SELECT id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at 
FROM custom_link
WHERE user_id = $1
`

func (q *Queries) ListCustomLinkByUser(ctx context.Context, userId uuid.UUID) ([]model.CustomLink, error) {
	rows, err := q.db.Query(ctx, listCustomLinkByUserQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.CustomLink
	for rows.Next() {
		var i model.CustomLink
		if err := rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Url,
			&i.Token,
			&i.IsSuggestion,
			&i.SuggestionId,
			&i.Deleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// DeleteCustomLink logic delete for a custom link.
const deleteCustomLinkQuery = `
UPDATE custom_link
SET url = $1, token = $2, deleted = $3, updated_at = $4
WHERE Id = $5
RETURNING id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at
`

func (q *Queries) DeleteCustomLink(ctx context.Context, customLinkID model.CustomLinkId) (*model.CustomLink, error) {
	deleteToken := repository.DeleteCustomTokenValue + strconv.Itoa(int(customLinkID))
	row := q.db.QueryRow(ctx, deleteCustomLinkQuery, repository.DeleteStringValue, deleteToken, true, time.Now().UTC(), customLinkID)
	var i model.CustomLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.IsSuggestion,
		&i.SuggestionId,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres DeleteCustomLink method error: %w", repository.ErrNotFound)
	}
	return &i, err
}

// DeleteCustomLinkByToken logic delete for a custom link by token.
const deleteCustomLinkByTokenQuery = `
UPDATE custom_link
SET url = $1, token = $2, deleted = $3, updated_at = $4
WHERE token = $5
RETURNING id, user_id, url, token, is_suggestion, suggestion_id, deleted, created_at, updated_at
`

func (q *Queries) DeleteCustomLinkByToken(ctx context.Context, customLinkToken model.CustomLinkToken) (*model.CustomLink, error) {
	rowLink := q.db.QueryRow(ctx, getCustomLinkByTokenQuery, customLinkToken)
	var customLink model.CustomLink
	err := rowLink.Scan(
		&customLink.Id,
		&customLink.UserId,
		&customLink.Url,
		&customLink.Token,
		&customLink.IsSuggestion,
		&customLink.SuggestionId,
		&customLink.Deleted,
		&customLink.CreatedAt,
		&customLink.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres GetCustomLinkByToken method error: %w", repository.ErrNotFound)
		return &customLink, err
	}

	deleteToken := repository.DeleteCustomTokenValue + strconv.Itoa(int(customLink.Id))
	row := q.db.QueryRow(ctx, deleteCustomLinkByTokenQuery, repository.DeleteStringValue, deleteToken, true, time.Now().UTC(), customLinkToken)
	var i model.CustomLink
	err = row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.IsSuggestion,
		&i.SuggestionId,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	return &i, err
}
