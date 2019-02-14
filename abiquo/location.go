package abiquo

import (
	"github.com/abiquo/ojal/core"
)

// Location represents a location resource
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	core.DTO
}

// PrivateLocations returns the private cloud locations from the Abiquo API
func PrivateLocations() *core.Link {
	return core.NewLink("cloud/locations").SetType("datacenters")
}

// PublicLocations returns the public cloud regions from the Abiquo API
func PublicLocations() *core.Link {
	return core.NewLink("cloud/locations").SetType("publiccloudregions")
}
