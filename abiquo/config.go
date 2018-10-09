package abiquo

import "github.com/abiquo/ojal/core"

// Category represents an Abiquo category
type Category struct {
	Name     string `json:"name"`
	Default  bool   `json:"defaultCategory"`
	ID       int    `json:"id,omitempty"`
	Erasable bool   `json:"erasable"`
	core.DTO
}

// CostCode represents an Abiquo cost code
type CostCode struct {
	CurrencyPrices []PricingResource `json:"currencyPrices"`
	Description    string            `json:"description"`
	Name           string            `json:"name"`
	core.DTO
}

// Currency represents an Abiquo currency
type Currency struct {
	Digits int    `json:"digits"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	core.DTO
}

// DeviceInterface represents an Abiquo SDN device type interface
type DeviceInterface struct {
	Interface   string            `json:"deviceInterface"`
	RealName    string            `json:"realName"`
	Constraints map[string]string `json:"constraints"`
	Operations  []interface{}     `json:"operations"`
}

// DeviceType represents an Abiquo SDN device type
type DeviceType struct {
	Name       string            `json:"name"`
	Interfaces []DeviceInterface `json:"deviceInterfaces"`
	core.DTO
}

// License represents an Abiquo License resource
type License struct {
	Code                 string `json:"code"`
	Expiration           string `json:"expiration,omitempty"`
	ID                   int    `json:"id,omitempty"`
	NumCores             int    `json:"numcores,omitempty"`
	ScalingGroupsEnabled bool   `json:"scalingGroupsEnabled,omitempty"`
	core.DTO
}

// PricingTemplate represents an Abiquo Pricing Template resource
type PricingTemplate struct {
	AbstractDCPrices     []PricingDatacenter `json:"pricingAbstractDatacenters,omitempty"`
	CostCodes            []PricingResource   `json:"pricingCostCodes,omitempty"`
	ChargingPeriod       int                 `json:"chargingPeriod"`
	DeployMessage        string              `json:"deployMessage"`
	DefaultTemplate      bool                `json:"defaultTemplate"`
	Description          string              `json:"description"`
	ID                   int                 `json:"id,omitempty"`
	MinimumCharge        int                 `json:"minimumCharge"`
	MinimumChargePeriod  int                 `json:"minimumChargePeriod"`
	Name                 string              `json:"name"`
	ShowChangesBefore    bool                `json:"showChangesBefore"`
	StandingChargePeriod int                 `json:"standingChargePeriod"`
	core.DTO
}

// PricingDatacenter represents an Abiquo Pricing Template resource datacenter entry
type PricingDatacenter struct {
	DatastoreTiers []PricingResource `json:"pricingDatastoreTiers,omitempty"`
	HPAbstractDC   []PricingResource `json:"pricingHardwareProfilesAbsDc,omitempty"`
	Tiers          []PricingResource `json:"pricingTiers,omitempty"`
	Firewall       float64           `json:"firewall"`
	HardDiskGB     float64           `json:"hdGB"`
	Layer          float64           `json:"layer"`
	LoadBalancer   float64           `json:"loadBalancer"`
	MemoryGB       float64           `json:"memoryGB"`
	MemoryOnGB     float64           `json:"memoryOnGB"`
	MemoryOffGB    float64           `json:"memoryOffGB"`
	NatIP          float64           `json:"natIp"`
	PublicIP       float64           `json:"publicIp"`
	RepositoryGB   float64           `json:"repositoryGB"`
	VCPU           float64           `json:"vcpu"`
	VCPUOn         float64           `json:"vcpuOn"`
	VCPUOff        float64           `json:"vcpuOff"`
	VLAN           float64           `json:"vlan"`
	core.DTO
}

// PricingResource represents an Abiquo Resource pricing entry
type PricingResource struct {
	Price float64 `json:"price"`
	core.DTO
}

// Privilege represents an Abiquo API Privilege DTO
type Privilege struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Events []string `json:"events"`
	core.DTO
}
