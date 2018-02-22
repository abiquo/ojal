package abiquo

import "github.com/abiquo/opal/core"

// VirtualAppliance represents a VAPP dto
type VirtualAppliance struct {
	Name string `json:"name"`
	core.DTO
}

// NewVirtualAppliance returns a VirtualAppliance DTO
func NewVirtualAppliance() core.Resource { return new(VirtualAppliance) }

// CreateVM creates a VM inside v
func (v *VirtualAppliance) CreateVM(vm *VirtualMachine) error {
	return core.Create(v.Rel("virtualmachines").SetType("virtualmachine"), vm)
}
