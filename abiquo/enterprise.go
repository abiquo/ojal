package abiquo

import (
	"fmt"
	"net/url"

	"github.com/abiquo/opal/core"
)

// Enterprise represents an Abiquo Enterprisecore.Resource
type Enterprise struct {
	ID   int    `json:"id,omitempty"` // The enterprise unique ID
	Name string `json:"name"`         // The name of the enterprise

	CPUSoft  int `json:"cpuSoft"` // Enterprise global soft limits
	HDSoft   int `json:"HDSoft"`
	IPSoft   int `json:"ipSoft"`
	RAMSoft  int `json:"ramSoft"`
	RepoSoft int `json:"repositorySoftInMb"`
	VolSoft  int `json:"volSoft"`
	VLANSoft int `json:"vlansSoft"`

	CPUHard  int `json:"cpuHard"` // Enterprise global hard limits
	HDHard   int `json:"HDHard"`
	IPHard   int `json:"ipHard"`
	RAMHard  int `json:"ramHard"`
	RepoHard int `json:"repositoryHardInMb"`
	VolHard  int `json:"volHard"`
	VLANHard int `json:"vlansHard"`
	core.DTO
}

func NewEnterprise() core.Resource { return new(Enterprise) }

// Enterprises returns a slice of enterprises
func Enterprises(query url.Values) *core.Collection {
	return core.NewLinker("admin/enterprises", "enterprises").Collection(query)
}

// Create creates the requested enterprise
func (e *Enterprise) Create() error {
	return core.Create(core.NewLinker("admin/enterprises", "enterprise"), e)
}

// Users returns the *Enterprise users collection
func (e *Enterprise) Users(query url.Values) *core.Collection {
	return e.Rel("users").Collection(query)
}

// VirtualAppliances returns the *Enterprise virtualappliances collection
func (e *Enterprise) VirtualAppliances(query url.Values) *core.Collection {
	return e.Rel("virtualappliances").Collection(query)
}

// VirtualDatacenters returns the *Enterprise virtualdatacenters collection
func (e *Enterprise) VirtualDatacenters(query url.Values) *core.Collection {
	return e.Rel("cloud/virtualdatacenters").Collection(query)
}

// VirtualMachines returns the *Enterprise virtualmachines collection
func (e *Enterprise) VirtualMachines(query url.Values) *core.Collection {
	return e.Rel("virtualmachines").Collection(query)
}

// CreateLimit
func (e *Enterprise) CreateLimit(l *Limit) error {
	return core.Create(e.Rel("limits").SetType("limit"), l)
}

// DatacenterRepositories returns the enterprise datacenter repositories collection
func (e *Enterprise) DatacenterRepositories(query url.Values) *core.Collection {
	return e.Rel("datacenterrepositories").Collection(query)
}

// ExampleEnterprise show all enterprises virtualmachines
func ExampleEnterprise() {
	for _, e := range Enterprises(nil).List() {
		enterprise := e.(*Enterprise)
		for _, v := range enterprise.VirtualMachines(nil).List() {
			fmt.Println(v.URL())
		}
	}
}
