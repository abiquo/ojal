package abiquo

import (
	"github.com/abiquo/ojal/core"
)

// Firewall defines a firewall policy resource
type Firewall struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProviderIS  string `json:"providerId,omitempty"`
	core.DTO
}

// FirewallRule defines firewall policy rule
type FirewallRule struct {
	ID         int      `json:"id,omitempty"`
	FromPort   int      `json:"fromPort"`
	ToPort     int      `json:"toPort"`
	Protocol   string   `json:"protocol"`
	ProviderID string   `json:"providerId"`
	Sources    []string `json:"sources"`
	Targets    []string `json:"targets"`
}

// FirewallRules defines a firewall policy rules
type FirewallRules struct {
	Collection []FirewallRule `json:"collection"`
}

// LoadBalancer defines a load balancer resource
type LoadBalancer struct {
	Algorithm    string                    `json:"algorithm"`
	HealthChecks *LoadBalancerHealthChecks `json:"healthChecks,omitempty"`
	Name         string                    `json:"name"`
	Addresses    *LoadBalancerAddresses    `json:"loadBalancerAddresses,omitempty"`
	Rules        *LoadBalancerRules        `json:"routingRules,omitempty"`
	core.DTO
}

// VMs returns a load balancer node list
func (l *LoadBalancer) VMs() (vms core.Links, err error) {
	resource, err := l.Walk("virtualmachines")
	if err != nil {
		return
	}
	vms = resource.(*core.DTO).Links
	return
}

// LoadBalancerAddress defines a load balancer address resource
type LoadBalancerAddress struct {
	Endpoint   string `json:"endpoint,omitempty"`
	Internal   bool   `json:"internal"`
	ProviderID string `json:"providerId,omitempty"`
	core.DTO
}

// Endpoints returns a load balancer endpoint list
func (l *LoadBalancerAddresses) Endpoints(internal bool) (ips []string) {
	if l != nil {
		for _, ip := range l.Collection {
			if ip.Internal == internal {
				ips = append(ips, ip.Endpoint)
			}
		}
	}
	return
}

// LoadBalancerAddresses represents a load balancer addresses collection
type LoadBalancerAddresses struct {
	Collection []LoadBalancerAddress `json:"collection"`
}

// LoadBalancerHealthCheck defines a load balancer health check
type LoadBalancerHealthCheck struct {
	Value        string `json:"value"`
	Protocol     string `json:"protocol"`
	Port         string `json:"port"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	IntervalInMs int    `json:"intervalInMs"`
	TimeoutInMs  int    `json:"timeoutInMs"`
	Attempts     int    `json:"attempts"`
}

// LoadBalancerHealthChecks defines a load balancer health checks collection
type LoadBalancerHealthChecks struct {
	Collection []LoadBalancerHealthCheck `json:"collection"`
}

// LoadBalancerRule represents a load balancer rule
type LoadBalancerRule struct {
	PortIn      string `json:"portIn"`
	PortOut     string `json:"portOut"`
	ProtocolIn  string `json:"protocolIn"`
	ProtocolOut string `json:"protocolOut"`
}

// LoadBalancerRules represents a load balancer rule set
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
