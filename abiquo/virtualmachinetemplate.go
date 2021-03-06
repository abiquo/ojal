package abiquo

import "github.com/abiquo/ojal/core"

// VirtualMachineTemplate represents an Abiquo API virtual machine template DTO
type VirtualMachineTemplate struct {
	ChefEnabled                      bool        `json:"chefEnabled"`
	CPURequired                      int         `json:"cpuRequired,omitempty"`
	EnableCPUHotAdd                  bool        `json:"enableCpuHotAdd"`
	EnableDisksHotReconfigure        bool        `json:"enableDisksHotReconfigure"`
	EnableNicsHotReconfigure         bool        `json:"enableNicsHotReconfigure"`
	EnableOnlyHPRecommended          bool        `json:"enableOnlyHPRecommended"`
	EnableRAMHotAdd                  bool        `json:"enableRamHotAdd"`
	EnableRemoteAccessHotReconfigure bool        `json:"enableRemoteAccessHotReconfigure"`
	GenerateGuestInitialPassword     bool        `json:"generateGuestInitialPassword"`
	GuestSetup                       string      `json:"guestSetup,omitempty"`
	EthernetDriverType               string      `json:"ethernetDriverType,omitempty"`
	ID                               string      `json:"id,omitempty"`
	LoginUser                        string      `json:"loginUser,omitempty"`
	Name                             string      `json:"name"`
	IconURL                          string      `json:"iconUrl,omitempty"`
	Description                      string      `json:"description"`
	Label                            string      `json:"label,omitempty"`
	OsType                           string      `json:"osType,omitempty"`
	RAMRequired                      int         `json:"ramRequired,omitempty"`
	State                            string      `json:"state,omitempty"`
	Variables                        interface{} `json:"variables,omitempty"`
	core.DTO
}
