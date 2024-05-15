package controller

import "github.com/JairoRiver/short_link_app/short_link/internal/repository"

// Controller defines a short link service controller.
type Controller struct {
	repo repository.Storer
}

// New creates a short link service controller.
func New(repo repository.Storer) *Controller {
	return &Controller{repo}
}
