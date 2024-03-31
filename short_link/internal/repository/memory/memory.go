package memory

import "github.com/JairoRiver/short_link_app/short_link/pkg/model"

// ShortLinkRepository defines a Shot link repository.
type Repository struct {
	shortLinkData       map[model.ShortLinkId]model.ShortLink
	customLinkData      map[model.CustomLinkId]model.CustomLink
	customLinkTokenData map[model.CustomLinkToken]model.CustomLink
}

// New creates a new memory Repository.
func New() *Repository {
	return &Repository{
		shortLinkData:       map[model.ShortLinkId]model.ShortLink{},
		customLinkData:      map[model.CustomLinkId]model.CustomLink{},
		customLinkTokenData: map[model.CustomLinkToken]model.CustomLink{},
	}
}
