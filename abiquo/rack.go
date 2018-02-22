package abiquo

import "github.com/abiquo/opal/core"

type Rack struct {
	HAEnabled bool   `json:"haEnabled,omitempty"`
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	NRSQ      int    `json:"nrsq,omitempty"`
	VlanIDMax int    `json:"vlanIdMax,omitempty"`
	VlanIDMin int    `json:"vlanIdMin,omitempty"`
	Reserved  int    `json:"vlanPerVdcReserved,omitempty"`
	core.DTO
}
