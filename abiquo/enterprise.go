package abiquo

import (
	"fmt"
	"net/url"

	"github.com/abiquo/ojal/core"
)

// Enterprise represents an Abiquo Enterprisecore.Resource
//
// Collections:
// - users
// - virtualappliances
// - virtualdatacenters
// - virtualmachines
// - datacenterrepositories
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

// Enterprises returns a slice of enterprises
func Enterprises(query url.Values) *core.Collection {
	return core.NewLinker("admin/enterprises", "enterprises").Collection(query)
}

// Create creates the requested enterprise
func (e *Enterprise) Create() error {
	return core.Create(core.NewLinker("admin/enterprises", "enterprise"), e)
}

// CreateLimit
func (e *Enterprise) CreateLimit(l *Limit) error {
	return core.Create(e.Rel("limits").SetType("limit"), l)
}

// ExampleEnterprise show all enterprises names
func ExampleEnterprise() {
	for _, e := range Enterprises(nil).List() {
		fmt.Println(e.URL())
	}
}

type EnterpriseProperties struct {
	Properties map[string]string `json:"properties"`
	core.DTO
}
