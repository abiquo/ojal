package abiquo

import (
	"github.com/abiquo/ojal/core"
)

// Network represents an Abiquo API Network DTO
//
// Collections:
// - ips
type Network struct {
	Address string `json:"address"`
	DNS1    string `json:"primaryDNS,omitempty"`
	DNS2    string `json:"secondaryDNS,omitempty"`
	Mask    int    `json:"mask"`
	Gateway string `json:"gateway"`
	Name    string `json:"name"`
	Suffix  string `json:"sufixDNS,omitempty"`
	Tag     int    `json:"tag,omitempty"`
	Type    string `json:"type"`
	core.DTO
}

// IP represents an Abiquo API IP DTO
type IP struct {
	Available  bool   `json:"available"`
	ProviderID string `json:"providerId"`
	ID         int    `json:"id,omitempty"`
	IPv6       bool   `json:"ipv6"`
	IP         string `json:"ip"`
	core.DTO
}
