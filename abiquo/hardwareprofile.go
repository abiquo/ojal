package abiquo

import (
	"github.com/abiquo/opal/core"
)

type HardwareProfile struct {
	Active  bool   `json:"active"`
	Name    string `json:"name"`
	CPU     int    `json:"cpu"`
	RAMInMB int    `json:"ramInMb"`
	core.DTO
}

func newHardwareProfile() core.Resource { return new(HardwareProfile) }
