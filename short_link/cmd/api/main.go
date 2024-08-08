package main

import (
	"github.com/JairoRiver/short_link_app/short_link/internal/api"
	"github.com/JairoRiver/short_link_app/short_link/internal/api/handler/rest"
	"github.com/JairoRiver/short_link_app/short_link/internal/controller"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/memory"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	repo := memory.New()
	control := controller.New(repo)
	handler := rest.New(control, config)

	server := api.New(handler)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
