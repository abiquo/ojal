package abiquo

import (
	"github.com/abiquo/ojal/core"
)

// Enterprise represents an enterprise resource
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

// EnterpriseProperties represents an enterprise properties resource
type EnterpriseProperties struct {
	Properties map[string]string `json:"properties"`
	core.DTO
}

// Limit represents an enterprise limit
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
	Active       bool   `json:"active,omitempty"`
	AuthType     string `json:"authType,omitempty"`
	Description  string `json:"description,omitempty"`
	Email        string `json:"email"`
	FirstLogin   bool   `json:"firstLogin,omitempty"`
	Locale       string `json:"locale"`
	Locked       bool   `json:"locked,omitempty"`
	Name         string `json:"name,omitempty"`
	Nick         string `json:"nick,omitempty"`
	Password     string `json:"password,omitempty"`
	PublicSSHKey string `json:"publicSshKey,omitempty"`
	Surname      string `json:"surname,omitempty"`
	core.DTO
}
