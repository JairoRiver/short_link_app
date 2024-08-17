package postgres

import (
	"context"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// GetAuxSKey retrieves data to generate S Keys.
const getAuxSKeyQuery = `
SELECT N, "End", Step, A0, N0 
FROM aux_s_key
LIMIT 1
`

func (q *Queries) GetAuxSKey(ctx context.Context) (*model.AuxSKey, error) {
	row := q.db.QueryRow(ctx, getAuxSKeyQuery)
	var i model.AuxSKey
	err := row.Scan(
		&i.N,
		&i.End,
		&i.Step,
		&i.A0,
		&i.N0,
	)

	if err == pgx.ErrNoRows {
		err = fmt.Errorf("Repository postgres GetAuxSKey method error: %w", repository.ErrNotFound)
	}

	return &i, err
}

// UpdateAuxSKey retrieves data to generate S Keys.
const updateAuxSKeyQuery = `
UPDATE aux_s_key
SET A0 = COALESCE($1, A0), 
    N0 = COALESCE($2, N0)
RETURNING N, "End", Step, A0, N0
`

type updateAuxSKeyParams struct {
	A0 pgtype.Int8
	N0 pgtype.Int8
}

func (q *Queries) UpdateAuxSKey(ctx context.Context, params repository.AuxSKeyParams) (*model.AuxSKey, error) {
	args := updateAuxSKeyParams{
		A0: pgtype.Int8{Int64: int64(params.A0.Value), Valid: params.A0.Valid},
		N0: pgtype.Int8{Int64: int64(params.N0.Value), Valid: params.N0.Valid},
	}
	row := q.db.QueryRow(ctx, updateAuxSKeyQuery, args.A0, args.N0)
	var i model.AuxSKey
	err := row.Scan(
		&i.N,
		&i.End,
		&i.Step,
		&i.A0,
		&i.N0,
	)

	return &i, err
}
