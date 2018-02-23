package abiquo

import (
	"net/url"

	"github.com/abiquo/ojal/core"
)

type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	core.DTO
}

// PrivateLocations returns the private cloud locations from the Abiquo API
func PrivateLocations(query url.Values) *core.Collection {
	return core.NewLinker("cloud/locations", "datacenters").Collection(query)
}

// PublicLocations returns the public cloud regions from the Abiquo API
func PublicLocations(query url.Values) *core.Collection {
	return core.NewLinker("cloud/locations", "publiccloudregions").Collection(query)
}
