package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

// Datacenter represents an Abiquo API datacentercore.Resource
type Datacenter struct {
	ID             int           `json:"id"`
	Location       string        `json:"location"`
	Name           string        `json:"name"`
	RemoteServices []interface{} `json:"remoteServices"`
	core.DTO
}

// NewDatacenter returns a new Abiquo API Datacenter DTO
func NewDatacenter() core.Resource { return new(Datacenter) }

// Datacenters returns the Abiquo API datacenters collection
func Datacenters(query url.Values) *core.Collection {
	return core.NewLinker("admin/datacenters", "datacenters").Collection(query)
}

// CreateExternal creates a new external network in the datacenter
func (d *Datacenter) CreateExternal(external *Network) (err error) {
	return core.Create(d.Rel("network").SetType("vlan"), external)
}

// Networks returns a Datacenter Networks
func (d *Datacenter) Networks(query url.Values) *core.Collection {
	return d.Rel("network").Collection(query)
}

// HardwareProfiles returns a location hardware profiles list
func (d *Datacenter) HardwareProfiles(query url.Values) *core.Collection {
	return d.Rel("hardwareprofiles").Collection(query)
}
