package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

// Datacenter represents an Abiquo API datacentercore.Resource
type Datacenter struct {
	ID             int    `json:"id,omitempty"`
	Location       string `json:"location"`
	Name           string `json:"name"`
	RemoteServices struct {
		Collection []RemoteService `json:"collection"`
		core.DTO
	} `json:"remoteServices"`
	core.DTO
}

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

// Discover returns the datacenter discover action Values
func (d *Datacenter) Discover(q url.Values) *core.Collection {
	return d.Rel("discover").Collection(q)
}
