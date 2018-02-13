package abiquo

import (
	"net/url"

	. "github.com/abiquo/opal/core"
)

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
	DTO
}

func NewVirtualDatacenter() Resource { return new(VirtualDatacenter) }

// Create creates a new VDC
func (v *VirtualDatacenter) Create() error {
	endpoint := NewLinker("cloud/virtualdatacenters", "virtualdatacenter")
	return Create(endpoint, v)
}

// CreateNetwork creates a new private network in the VDC
func (v *VirtualDatacenter) CreateNetwork(network *Network) error {
	return Create(v.Rel("privatenetworks").SetType("vlan"), network)
}

// VirtualAppliances returns the list of virtualappliances inside a VirtualDatacenter
func (v *VirtualDatacenter) VirtualAppliances(query url.Values) *Collection {
	return v.Rel("virtualappliances").Collection(query)
}
