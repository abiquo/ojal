package abiquo

import (
	. "github.com/abiquo/opal/core"
)

type HardwareProfile struct {
	Active  bool   `json:"active"`
	Name    string `json:"name"`
	CPU     int    `json:"cpu"`
	RAMInMB int    `json:"ramInMb"`
	DTO
}

func newHardwareProfile() Resource { return new(HardwareProfile) }
