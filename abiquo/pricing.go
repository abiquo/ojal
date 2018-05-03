package abiquo

import "github.com/abiquo/ojal/core"

type Currency struct {
	Digits int    `json:"digits"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	core.DTO
}

type price struct {
	Price float64 `json:"price"`
	core.DTO
}

type CurrencyPrice price

type CostCode struct {
	CurrencyPrices []CurrencyPrice `json:"currencyPrices"`
	Description    string          `json:"description"`
	Name           string          `json:"name"`
	core.DTO
}

type DatastoreTierPrice price
type HardwareProfilePrice price
type TierPrice price

type AbstractDCPrice struct {
	DatastoreTiers []DatastoreTierPrice   `json:"pricingDatastoreTiers,omitempty"`
	HPAbstractDC   []HardwareProfilePrice `json:"pricingHardwareProfilesAbsDc,omitempty"`
	Tiers          []TierPrice            `json:"pricingTiers,omitempty"`
	Firewall       float64                `json:"firewall"`
	HardDiskGB     float64                `json:"hdGB"`
	Layer          float64                `json:"layer"`
	LoadBalancer   float64                `json:"loadBalancer"`
	MemoryGB       float64                `json:"memoryGB"`
	MemoryOnGB     float64                `json:"memoryOnGB"`
	MemoryOffGB    float64                `json:"memoryOffGB"`
	NatIP          float64                `json:"natIp"`
	PublicIP       float64                `json:"publicIp"`
	RepositoryGB   float64                `json:"repositoryGB"`
	VCPU           float64                `json:"vcpu"`
	VCPUOn         float64                `json:"vcpuOn"`
	VCPUOff        float64                `json:"vcpuOff"`
	VLAN           float64                `json:"vlan"`
	core.DTO
}

type PricingTemplate struct {
	AbstractDCPrices     []AbstractDCPrice `json:"pricingAbstractDatacenters,omitempty"`
	CostCodes            []CurrencyPrice   `json:"pricingCostCodes,omitempty"`
	ChargingPeriod       int               `json:"chargingPeriod"`
	DeployMessage        string            `json:"deployMessage"`
	DefaultTemplate      bool              `json:"defaultTemplate"`
	Description          string            `json:"description"`
	ID                   int               `json:"id,omitempty"`
	MinimumCharge        int               `json:"minimumCharge"`
	MinimumChargePeriod  int               `json:"minimumChargePeriod"`
	Name                 string            `json:"name"`
	ShowChangesBefore    bool              `json:"showChangesBefore"`
	StandingChargePeriod int               `json:"standingChargePeriod"`
	core.DTO
}
