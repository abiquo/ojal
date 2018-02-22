package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

type PublicLocation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	core.DTO
}

type PrivateLocation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	core.DTO
}

func newPrivateLocation() core.Resource { return new(PrivateLocation) }
func newPublicLocation() core.Resource  { return new(PublicLocation) }

// PrivateLocations returns the private cloud locations from the Abiquo API
func PrivateLocations(query url.Values) *core.Collection {
	return core.NewLinker("cloud/locations", "datacenters").Collection(query)
}

// PublicLocations returns the public cloud regions from the Abiquo API
func PublicLocations(query url.Values) *core.Collection {
	return core.NewLinker("cloud/locations", "publiccloudregions").Collection(query)
}
