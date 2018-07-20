package abiquo

import (
	"github.com/abiquo/ojal/core"
)

type Firewall struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProviderId  string `json:"providerId,omitempty"`
	core.DTO
}

type FirewallRule struct {
	Id         int      `json:"id,omitempty"`
	FromPort   int      `json:"fromPort"`
	ToPort     int      `json:"toPort"`
	Protocol   string   `json:"protocol"`
	ProviderId string   `json:"providerId"`
	Sources    []string `json:"sources"`
	Targets    []string `json:"targets"`
}

type FirewallRules struct {
	Collection []FirewallRule `json:"collection"`
}

type LoadBalancer struct {
	Algorithm    string                 `json:"algorithm"`
	HealthChecks []interface{}          `json:"healthChecks,omitempty"`
	Name         string                 `json:"name"`
	Addresses    *LoadBalancerAddresses `json:"loadBalancerAddresses,omitempty"`
	Rules        *LoadBalancerRules     `json:"routingRules,omitempty"`
	core.DTO
}

type LoadBalancerAddress struct {
	Internal bool `json:"internal"`
}

type LoadBalancerAddresses struct {
	Collection []LoadBalancerAddress `json:"collection"`
}

type LoadBalancerRule struct {
	PortIn      string `json:"portIn"`
	PortOut     string `json:"portOut"`
	ProtocolIn  string `json:"protocolIn"`
	ProtocolOut string `json:"protocolOut"`
}

type LoadBalancerRules struct {
	Collection []LoadBalancerRule `json:"collection"`
}

// VirtualDatacenter represents an Abiquo API Virtual datacenter DTO
//
// Collections:
// - virtualappliances
type VirtualDatacenter struct {
	CPUHard     int      `json:"cpuHard"`
	CPUSoft     int      `json:"cpuSoft"`
	DiskHard    int      `json:"diskHardLimitInMb"`
	DiskSoft    int      `json:"diskSoftLimitInMb"`
	HVType      string   `json:"hypervisorType"`
	Name        string   `json:"name"`
	Network     *Network `json:"network"`
	PublicHard  int      `json:"publicIpsHard"`
	PublicSoft  int      `json:"publicIpsSoft"`
	RAMHard     int      `json:"ramHard"`
	RAMSoft     int      `json:"ramSoft"`
	StorageHard int      `json:"storageHardInMb"`
	StorageSoft int      `json:"storageSoftInMb"`
	VLANHard    int      `json:"vlansHard"`
	VLANSoft    int      `json:"vlansSoft"`
	core.DTO
}
