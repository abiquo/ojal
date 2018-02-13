package abiquo

import (
	"net/url"

	. "github.com/abiquo/opal/core"
)

type PublicLocation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DTO
}

type PrivateLocation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DTO
}

func newPrivateLocation() Resource { return new(PrivateLocation) }
func newPublicLocation() Resource  { return new(PublicLocation) }

// PrivateLocations returns the private cloud locations from the Abiquo API
func PrivateLocations(query url.Values) *Collection {
	return NewLinker("cloud/locations", "datacenters").Collection(query)
}

// PublicLocations returns the public cloud regions from the Abiquo API
func PublicLocations(query url.Values) *Collection {
	return NewLinker("cloud/locations", "publiccloudregions").Collection(query)
}
