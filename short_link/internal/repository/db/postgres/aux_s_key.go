package postgres

import (
	"context"
	"fmt"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
	"github.com/jackc/pgx/v5"
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
