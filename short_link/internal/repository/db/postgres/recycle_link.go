package postgres

import (
	"context"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/jackc/pgx/v5"
)

// PutRecycleLink adds a new recycle link.
const createRecycleLinkQuery = `
INSERT INTO recycle_link (
                        s_key
) VALUES (
  $1
) RETURNING s_key
`

func (q *Queries) PutRecycleLink(ctx context.Context, recycleLink model.RecycleLink) error {
	row := q.db.QueryRow(ctx, createRecycleLinkQuery, recycleLink.SKey)
	var i model.RecycleLink
	err := row.Scan(
		&i.SKey,
	)
	return err
}

// GetRecycleLink retrieves a recycle link by SKey.
const getRecycleLinkQuery = `
SELECT s_key
FROM recycle_link
ORDER BY id
LIMIT 1
`

func (q *Queries) GetRecycleLink(ctx context.Context) (*model.RecycleLink, error) {
	row := q.db.QueryRow(ctx, getRecycleLinkQuery)
	var i model.RecycleLink
	err := row.Scan(
		&i.SKey,
	)

	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres GetRecycleLink method error: %w", repository.ErrNotFound)
	}

	return &i, err
}

// DeleteRecycleLink delete a recycle link.
const deleteRecycleLinkQuery = `
DELETE FROM recycle_link
WHERE s_key = $1
RETURNING s_key
`

func (q *Queries) DeleteRecycleLink(ctx context.Context, recycleLinkID model.RecycleLinkId) error {
	row := q.db.QueryRow(ctx, deleteRecycleLinkQuery, recycleLinkID)
	var i model.RecycleLink
	err := row.Scan(
		&i.SKey,
	)
	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres DeleteRecycleLink method error: %w", repository.ErrNotFound)
	}

	return err
}
