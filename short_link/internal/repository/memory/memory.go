package memory

import "github.com/JairoRiver/short_link_app/short_link/pkg/model"

// ShortLinkRepository defines a Shot link repository.
type Repository struct {
	shortLinkData       map[model.ShortLinkId]model.ShortLink
	customLinkData      map[model.CustomLinkId]model.CustomLink
	customLinkTokenData map[model.CustomLinkToken]model.CustomLink
	recycleLinkData     map[model.RecycleLinkId]model.RecycleLink
	auxSKey             model.AuxSKey
	//aux map
	lastIds map[string]int
}

// New creates a new memory Repository.
func New() *Repository {
	return &Repository{
		shortLinkData:       map[model.ShortLinkId]model.ShortLink{},
		customLinkData:      map[model.CustomLinkId]model.CustomLink{},
		customLinkTokenData: map[model.CustomLinkToken]model.CustomLink{},
		recycleLinkData:     map[model.RecycleLinkId]model.RecycleLink{},
		lastIds:             map[string]int{"shortLink": 0, "customLink": 0},
		auxSKey:             model.AuxSKey{N: 60, Init: 931151403, End: 56800235583, Step: 931151403, A0: 0, N0: 1},
	}
}
