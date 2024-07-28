package main

import (
	"github.com/JairoRiver/short_link_app/short_link/internal/api"
	"github.com/JairoRiver/short_link_app/short_link/internal/api/handler/rest"
	"github.com/JairoRiver/short_link_app/short_link/internal/controller"
	"github.com/JairoRiver/short_link_app/short_link/internal/repository/memory"
)

func main() {
	repo := memory.New()
	control := controller.New(repo)
	handler := rest.New(control)

	server := api.New(handler)
	server.Start("0.0.0.0:8080")
}
