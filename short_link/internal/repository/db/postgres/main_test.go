package postgres

import (
	"context"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
)

var testStore Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../..", "app")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	testStore = *New(connPool)
	os.Exit(m.Run())
}
