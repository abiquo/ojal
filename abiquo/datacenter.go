package abiquo

import (
	"net/url"

	"github.com/abiquo/ojal/core"
)

// Datacenter represents an Abiquo API datacentercore.Resource
//
// Collections:
// - networks
// - hardwareprofiles
// - discover
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
