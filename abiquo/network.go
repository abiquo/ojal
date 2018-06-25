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
	TypeNet string `json:"type"`
	core.DTO
}
