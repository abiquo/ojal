package abiquo

import (
	"github.com/abiquo/opal/core"
)

// VirtualDatacenter represents an Abiquo API Virtual datacenter DTO
//
// Collections:
// - virtualappliances
type VirtualDatacenter struct {
	// Soft limits
	CPUSoft     int `json:"cpuSoft"`
	DiskSoft    int `json:"diskSoftLimitInMb"`
	PublicSoft  int `json:"publicIpsSoft"`
	RAMSoft     int `json:"ramSoft"`
	StorageSoft int `json:"storageSoftInMb"`
	VLANSoft    int `json:"vlansSoft"`
	//Hard limits
	CPUHard     int `json:"cpuHard"`
	DiskHard    int `json:"diskHardLimitInMb"`
	PublicHard  int `json:"publicIpsHard"`
	RAMHard     int `json:"ramHard"`
	StorageHard int `json:"storageHardInMb"`
	VLANHard    int `json:"vlansHard"`
	//
	Name    string   `json:"name"`
	HVType  string   `json:"hypervisorType"`
	Network *Network `json:"network"`
	core.DTO
}

// Create creates a new VDC
func (v *VirtualDatacenter) Create() error {
	endpoint := core.NewLinker("cloud/virtualdatacenters", "virtualdatacenter")
	return core.Create(endpoint, v)
}

// CreateNetwork creates a new private network in the VDC
func (v *VirtualDatacenter) CreateNetwork(network *Network) error {
	return core.Create(v.Rel("privatenetworks").SetType("vlan"), network)
}
