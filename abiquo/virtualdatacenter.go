package abiquo

import (
	"github.com/abiquo/ojal/core"
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

// VirtualAppliance represents a VAPP dto
type VirtualAppliance struct {
	Name string `json:"name"`
	core.DTO
}
