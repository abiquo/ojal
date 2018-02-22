package abiquo

import "github.com/abiquo/opal/core"

// VirtualMachineTemplate represents an Abiquo API virtual machine template DTO
type VirtualMachineTemplate struct {
	ChefEnabled                      bool        `json:"chefEnabled"`
	CPURequired                      int         `json:"cpuRequired,omitempty"`
	EnableCPUHotAdd                  bool        `json:"enableCpuHotAdd"`
	EnableDisksHotReconfigure        bool        `json:"enableDisksHotReconfigure"`
	EnableNicsHotReconfigure         bool        `json:"enableNicsHotReconfigure"`
	EnableOnlyHPRecommended          bool        `json:"enableOnlyHPRecommended"`
	EnableRamHotAdd                  bool        `json:"enableRamHotAdd"`
	EnableRemoteAccessHotReconfigure bool        `json:"enableRemoteAccessHotReconfigure"`
	GenerateGuestInitialPassword     bool        `json:"GenerateGuestInitialPassword"`
	EthernetDriverType               string      `json:"ethernetDriverType,omitempty"`
	ID                               string      `json:"id,omitempty"`
	Name                             string      `json:"name"`
	IconURL                          string      `json:"iconUrl,omitempty"`
	Description                      string      `json:"description"`
	Label                            string      `json:"label,omitempty"`
	OsType                           string      `json:"osType,omitempty"`
	RAMRequired                      int         `json:"ramRequired,omitempty"`
	Variables                        interface{} `json:"variables,omitempty"`
	core.DTO
}
