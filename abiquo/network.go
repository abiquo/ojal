package abiquo

import (
	"fmt"

	"github.com/abiquo/opal/core"
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

// CreateIP creates the provided IP in the Network
func (n *Network) CreateIP(i *IP) error {
	var media string
	switch n.TypeNet {
	case "INTERNAL":
		media = "privateip"
	default:
		return fmt.Errorf("CreateIP: %v ip type not implemented", n.TypeNet)
	}
	return core.Create(n.Rel("ips").SetType(media), i)
}
