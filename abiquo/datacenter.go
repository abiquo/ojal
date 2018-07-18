package abiquo

import (
	"net/url"

	"github.com/abiquo/ojal/core"
)

// Datacenter represents an Abiquo API datacentercore.Resource
//
// Collections:
// - backuppolicies
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

// Rack represents an Abiquo Datacenter rack
type Rack struct {
	Description string `json:"shortDescription"`
	HAEnabled   bool   `json:"haEnabled,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	NRSQ        int    `json:"nrsq,omitempty"`
	VlanIDMax   int    `json:"vlanIdMax,omitempty"`
	VlanIDMin   int    `json:"vlanIdMin,omitempty"`
	Reserved    int    `json:"vlanPerVdcReserved,omitempty"`
	core.DTO
}

type NetworkServiceType struct {
	Default bool   `json:"defaultNST"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	core.DTO
}

type Machine struct {
	AgentUser  string `json:"agentUser"`
	CPU        int    `json:"cpu"`
	CPUUsed    int    `json:"cpuUsed"`
	Datastores struct {
		Collection []*Datastore `json:"collection"`
		core.DTO
	} `json:"datastores"`
	Description string `json:"description"`
	Hostname    string `json:"hostname"`
	Initiator   string `json:"initiatorIQN"`
	IP          string `json:"ip"`
	IPService   string `json:"ipService"`
	ManagerIP   string `json:"managerIp,omitempty"`
	ManagerPass string `json:"managerPassword,omitempty"`
	ManagerUser string `json:"managerUser,omitempty"`
	Name        string `json:"name"`
	Interfaces  struct {
		Collection []*struct {
			MAC  string `json:"mac"`
			Name string `json:"name"`
			core.DTO
		} `json:"collection"`
		core.DTO
	} `json:"networkInterfaces"`
	Password string `json:"password,omitempty"`
	Port     int    `json:"port,omitempty"`
	RAM      int    `json:"ram"`
	RAMUsed  int    `json:"ramUsed"`
	State    string `json:"state"`
	Type     string `json:"type"`
	User     string `json:"user"`
	core.DTO
}
