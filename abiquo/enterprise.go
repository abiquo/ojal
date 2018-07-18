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
	IPSoft   int `json:"publicIpsSoft"`
	RAMSoft  int `json:"ramSoft"`
	RepoSoft int `json:"repositorySoftInMb"`
	VolSoft  int `json:"volSoft"`
	VLANSoft int `json:"vlansSoft"`

	CPUHard  int `json:"cpuHard"` // Enterprise global hard limits
	HDHard   int `json:"HDHard"`
	IPHard   int `json:"publicIpsHard"`
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

// Limit represents an abiquo enterprise limit
type Limit struct {
	CPUHard   int  `json:"cpuHard"`
	CPUSoft   int  `json:"cpuSoft"`
	HDHard    int  `json:"HDHard"`
	HDSoft    int  `json:"HDSoft"`
	EnableHPs bool `json:"enabledHardwareProfiles"`
	IPHard    int  `json:"ipHard"`
	IPSoft    int  `json:"ipSoft"`
	RAMHard   int  `json:"ramHard"`
	RAMSoft   int  `json:"ramSoft"`
	RepoHard  int  `json:"repositoryHardInMb"`
	RepoSoft  int  `json:"repositorySoftInMb"`
	VLANHard  int  `json:"vlansHard"`
	VLANSoft  int  `json:"vlansSoft"`
	VolHard   int  `json:"volHard"`
	VolSoft   int  `json:"volSoft"`
	core.DTO
}

// User represents an Abiquo enterprise user
type User struct {
	Active      bool   `json:"active,omitempty"`
	AuthType    string `json:"authType,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email"`
	FirstLogin  bool   `json:"firstLogin,omitempty"`
	Locale      string `json:"locale"`
	Locked      bool   `json:"locked,omitempty"`
	Name        string `json:"name,omitempty"`
	Nick        string `json:"nick,omitempty"`
	Password    string `json:"password,omitempty"`
	Surname     string `json:"surname,omitempty"`
	core.DTO
}
