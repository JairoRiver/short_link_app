package controller

import (
	"errors"
	"fmt"

	"github.com/JairoRiver/short_link_app/short_link/internal/repository"
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
)

// Controller defines a short link service controller.
type Controller struct {
	repo repository.Storer
}

// New creates a short link service controller.
func New(repo repository.Storer) *Controller {
	return &Controller{repo}
}

var ErrInvalidCustomToken = errors.New("error invalid custom token")

func decodingToken(token string) (uint64, error) {
	if len(token) > util.MaxLenToken {
		return 0, fmt.Errorf("Controller GetByToken token len must be lower than 7, error: %w", util.ErrInvalidToken)
	}

	s_k, err := util.FromBase62(token)
	if err != nil {
		return 0, fmt.Errorf("Controller GetByToken util.FromBase62 error: %w", err)
	}

	return s_k, nil
}
