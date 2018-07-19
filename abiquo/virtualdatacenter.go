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
	Algorithm             string                `json:"algorithm"`
	HealthChecks          []interface{}         `json:"healthChecks,omitempty"`
	Name                  string                `json:"name"`
	LoadBalancerAddresses LoadBalancerAddresses `json:"loadBalancerAddresses"`
	RoutingRules          LoadBalancerRules     `json:"routingRules"`
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
