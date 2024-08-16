package postgres

import (
	"context"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"time"
)

// PutShortLink add a new short link
const createShortLinkQuery = `
INSERT INTO short_link (
                        user_id,
                        url,
                        token,
                        s_key,
                        deleted
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, user_id, url, token, s_key, deleted, created_at, updated_at
`

func (q *Queries) PutShortLink(ctx context.Context, shotLinkParams repository.CreateShortLinkParams) (model.ShortLink, error) {
	var userId *uuid.UUID
	if shotLinkParams.UserId.Valid {
		userId = &shotLinkParams.UserId.ID
	}
	row := q.db.QueryRow(ctx, createShortLinkQuery, userId, shotLinkParams.Url, shotLinkParams.Token, shotLinkParams.SKey, shotLinkParams.Deleted)
	var i model.ShortLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.SKey,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

// GetShortLinkByID retrieves a shot link by Id.
const getShortLinkByIdQuery = `
SELECT id, user_id, url, token, s_key, deleted, created_at, updated_at 
FROM short_link
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetShortLinkByID(ctx context.Context, shotLinkID model.ShortLinkId) (*model.ShortLink, error) {
	row := q.db.QueryRow(ctx, getShortLinkByIdQuery, shotLinkID)
	var i model.ShortLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.SKey,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres GetShortLinkByID method error: %w", repository.ErrNotFound)
	}

	return &i, err
}

// GetShortLinkBySKey retrieves a shot link by S_key.
const getShortLinkBySkQuery = `
SELECT id, user_id, url, token, s_key, deleted, created_at, updated_at 
FROM short_link
WHERE s_key = $1
LIMIT 1
`

func (q *Queries) GetShortLinkBySKey(ctx context.Context, sKeyID model.ShortLinkId) (*model.ShortLink, error) {
	row := q.db.QueryRow(ctx, getShortLinkBySkQuery, sKeyID)
	var i model.ShortLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.SKey,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres GetShortLinkBySKey method error: %w", repository.ErrNotFound)
	}

	return &i, err
}

// ListAllShortLink retrieves all Short Links.
const listAllShortLinkQuery = `
SELECT id, user_id, url, token, s_key, deleted, created_at, updated_at 
FROM short_link
`

func (q *Queries) ListAllShortLink(ctx context.Context) ([]model.ShortLink, error) {
	rows, err := q.db.Query(ctx, listAllShortLinkQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.ShortLink
	for rows.Next() {
		var i model.ShortLink
		if err := rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Url,
			&i.Token,
			&i.SKey,
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

// ListActiveShortLink retrieves all Short Links not deleted.
const listActiveShortLinkQuery = `
SELECT id, user_id, url, token, s_key, deleted, created_at, updated_at 
FROM short_link
WHERE deleted IS false
`

func (q *Queries) ListActiveShortLink(ctx context.Context) ([]model.ShortLink, error) {
	rows, err := q.db.Query(ctx, listActiveShortLinkQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.ShortLink
	for rows.Next() {
		var i model.ShortLink
		if err := rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Url,
			&i.Token,
			&i.SKey,
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

// ListShortLinkByUser retrieves all Short Links by User.
const listShortLinkByUserQuery = `
SELECT id, user_id, url, token, s_key, deleted, created_at, updated_at 
FROM short_link
WHERE user_id = $1
`

type ListShortLinkByUserParams struct {
	UserId uuid.UUID `json:"user_id"`
}

func (q *Queries) ListShortLinkByUser(ctx context.Context, arg ListShortLinkByUserParams) ([]model.ShortLink, error) {
	rows, err := q.db.Query(ctx, listShortLinkByUserQuery, arg.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.ShortLink
	for rows.Next() {
		var i model.ShortLink
		if err := rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Url,
			&i.Token,
			&i.SKey,
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

// DeleteShortLink logic delete for a short link.
const deleteShortLinkQuery = `
UPDATE short_link
SET url = $1, token = $2, s_key= $3, deleted = $4, updated_at = $5
WHERE Id = $6 
RETURNING id, user_id, url, token, s_key, deleted, created_at, updated_at
`

type DeleteShortLinkParams struct {
	ID int64 `json:"id"`
}

func (q *Queries) DeleteShortLink(ctx context.Context, args DeleteShortLinkParams) (*model.ShortLink, error) {
	row := q.db.QueryRow(ctx, deleteShortLinkQuery, repository.DeleteStringValue, "", 0, true, time.Now(), args.ID)
	var i model.ShortLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.SKey,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres DeleteShortLink method error: %w", repository.ErrNotFound)
	}
	return &i, err
}

// DeleteShortLinkBySKey retrieves a shot link by S_key.
const deleteShortLinkBySkQuery = `
UPDATE short_link
SET url = $1, token = $2, s_key= $3, deleted = $4, updated_at = $5
WHERE s_key = $6 
RETURNING id, user_id, url, token, s_key, deleted, created_at, updated_at
`

type DeleteShortLinkBySkParams struct {
	SKey int64 `json:"id"`
}

func (q *Queries) DeleteShortLinkBySK(ctx context.Context, args DeleteShortLinkBySkParams) (*model.ShortLink, error) {
	row := q.db.QueryRow(ctx, deleteShortLinkBySkQuery, repository.DeleteStringValue, "", 0, true, time.Now(), args.SKey)
	var i model.ShortLink
	err := row.Scan(
		&i.Id,
		&i.UserId,
		&i.Url,
		&i.Token,
		&i.SKey,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres DeleteShortLinkBySK method error: %w", repository.ErrNotFound)
	}
	return &i, err
}
