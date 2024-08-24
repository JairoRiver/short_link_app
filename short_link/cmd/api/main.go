package main

import (
	"context"
	"github.com/JairoRiver/short_link_app/short_link/internal/api"
	"github.com/JairoRiver/short_link_app/short_link/internal/api/handler/rest"
	"github.com/JairoRiver/short_link_app/short_link/internal/controller"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/db/postgres"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/memory"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	var repo repository.Storer
	if config.Environment == "development" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		repo = memory.New()
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		connPool, err := pgxpool.New(context.Background(), config.DBSource)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot connect to db")
		}
		repo = postgres.New(connPool)
	}

	control := controller.New(repo)
	handler := rest.New(control, config)

	server := api.New(handler)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server:")
	}
	log.Info().Msgf("start HTTP gateway server at %s", config.ServerAddress)
}
