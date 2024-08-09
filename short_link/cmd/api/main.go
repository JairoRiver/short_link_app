package main

import (
	"github.com/JairoRiver/short_link_app/short_link/internal/api"
	"github.com/JairoRiver/short_link_app/short_link/internal/api/handler/rest"
	"github.com/JairoRiver/short_link_app/short_link/internal/controller"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/memory"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
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

	if config.Environment == "development" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	repo := memory.New()
	control := controller.New(repo)
	handler := rest.New(control, config)

	server := api.New(handler)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server:")
	}
	log.Info().Msgf("start HTTP gateway server at %s", config.ServerAddress)
}
