package memory

import (
	"github.com/JairoRiver/short_link_app/short_link/internal/util"
	"github.com/JairoRiver/short_link_app/short_link/pkg/model"
)

// ShortLinkRepository defines a Shot link repository.
type Repository struct {
	shortLinkData       map[model.ShortLinkId]model.ShortLink
	customLinkData      map[model.CustomLinkId]model.CustomLink
	customLinkTokenData map[model.CustomLinkToken]model.CustomLink
	recycleLinkData     []model.RecycleLink
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
		lastIds:             map[string]int{"shortLink": 0, "customLink": 0},
		auxSKey:             model.AuxSKey{N: util.SknValue, Init: util.SkinitValue, End: util.SkendValue, Step: util.SkstepValue, A0: util.Ska0Value, N0: util.Skn0Value},
	}
}
