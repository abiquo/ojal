package abiquo

import (
	"fmt"

	"github.com/abiquo/ojal/core"
)

// ActionPlan represents an Abiquo action plan DTO
type ActionPlan struct {
	CreatedBy   string            `json:"createdBy,omitempty"`
	Description string            `json:"description"`
	Entries     []ActionPlanEntry `json:"entries"`
	Name        string            `json:"name"`
	core.DTO
}

// ActionPlanEntry represents an Abiquo action plan entry
type ActionPlanEntry struct {
	Parameter     string `json:"parameter"`
	ParameterType string `json:"parameterType"`
	Sequence      int    `json:"sequence"`
	Type          string `json:"type"`
	core.DTO
}

// Alarm represents an Abiquo alert DTO
type Alarm struct {
	Active           bool    `json:"active"`
	TimeRangeMinutes int     `json:"timeRangeMinutes"`
	DataPointsLimit  int     `json:"datapointslimit"`
	Formula          string  `json:"formula"`
	Name             string  `json:"name"`
	Statistic        string  `json:"statistic"`
	Threshold        float64 `json:"threshold"`
	core.DTO
}

// VirtualMachine represents an Abiquo API VMcore.Resource
type VirtualMachine struct {
	Backups         []BackupPolicy    `json:"backupPolicies,omitempty"`
	CPU             int               `json:"cpu,omitempty"`
	ID              int               `json:"id,omitempty"`
	FQDN            string            `json:"fqdn,omitempty"`
	Label           string            `json:"label,omitempty"`
	LastSynchronize int64             `json:"lastSynchronize"`
	Layer           string            `json:"layer,omitempty"`
	Metadata        string            `json:"metadata,omitempty"`
	Monitored       bool              `json:"monitored"`
	Name            string            `json:"name,omitempty"`
	Password        string            `json:"password,omitempty"`
	RAM             int               `json:"ram,omitempty"`
	State           string            `json:"state,omitempty"`
	Variables       map[string]string `json:"variables,omitempty"`
	VdrpEnabled     bool              `json:"vdrpEnabled,omitempty"`
	VdrpPort        int               `json:"vdrpPort,omitempty"`
	UUID            string            `json:"uuid,omitempty"`
	core.DTO
}

// VirtualMachineMetadata is used to configure a VM metadata section
type VirtualMachineMetadata struct {
	Metadata VirtualMachineMetadataFields `json:"metadata"`
	core.DTO
}

// VirtualMachineMetadataFields is used to configure a VM metadata fields
type VirtualMachineMetadataFields struct {
	StartupScript string `json:"startup-script,omitempty"`
}

// VirtualMachineState is used to change a VM state
type VirtualMachineState struct {
	State            string `json:"state,omitempty"`
	GracefulShutdown bool   `json:"gracefulShutdown,omitempty"`
	core.DTO
}

// Reconfigure reconfigures v
func (v *VirtualMachine) Reconfigure() (err error) {
	_, err = core.Rest(nil, core.Put(v.URL(), "acceptedrequest", v.Media(), v))
	return
}

// Deploy deploys v
func (v *VirtualMachine) Deploy() (err error) {
	return NewTask(core.Post(
		v.Rel("deploy").Href,
		"acceptedrequest",
		"virtualmachinetask",
		v,
	))
}

// Undeploy undeploys v
func (v *VirtualMachine) Undeploy() (err error) {
	return NewTask(core.Post(
		v.Rel("undeploy").Href,
		"acceptedrequest",
		"virtualmachinetask",
		v,
	))
}

// Reboot the VM
func (v *VirtualMachine) Reboot() (err error) {
	return NewTask(core.Post(
		v.Rel("reset").Href,
		"acceptedrequest",
		"application/json",
		nil,
	))
}

// Off powers off the VM
func (v *VirtualMachine) Off() (err error) {
	return NewTask(core.Put(
		v.Rel("state").Href,
		"acceptedrequest",
		"virtualmachinestate",
		&VirtualMachineState{
			State: "OFF",
		},
	))
}

// On powers on the VM
func (v *VirtualMachine) On() (err error) {
	return NewTask(core.Put(
		v.Rel("state").Href,
		"acceptedrequest",
		"virtualmachinestate",
		&VirtualMachineState{
			State: "ON",
		},
	))
}

// SetMetadata sets the VM metadata as requested
func (v *VirtualMachine) SetMetadata(metadata *VirtualMachineMetadata) error {
	return core.Update(v.Rel("metadata"), metadata)
}

// Disks returns an slice with the VM disk links
func (v *VirtualMachine) Disks() []*core.Link {
	return v.Links.Filter(func(l *core.Link) bool {
		return l.IsMedia("harddisk") || l.IsMedia("volume")
	})
}

// NICs returns an slice with the VM NIC links
func (v *VirtualMachine) NICs() (nics []*core.Link) {
	return v.Links.Filter(func(l *core.Link) bool {
		return l.IsMedia("privateip") || l.IsMedia("externalip") || l.IsMedia("publicip")
	})
}

// AttachDisk add a disk link to the *VirtualMachine in the last position
func (v *VirtualMachine) AttachDisk(disk Disk) error {
	diskLink := disk.Link().SetRel(fmt.Sprintf("disk%v", len(v.Disks())))
	diskLink.DiskControllerType = disk.ControllerType()
	diskLink.DiskController = disk.Controller()
	v.Add(diskLink)
	return nil
}

// AttachNIC add a nic link to the *VirtualMachine in the last position
func (v *VirtualMachine) AttachNIC(nic *core.Link) error {
	nicLink := nic.SetRel(fmt.Sprintf("nic%v", len(v.NICs())))
	v.Add(nicLink)
	return nil
}

// Delete removes an existing VirtualMachine from the API
func (v *VirtualMachine) Delete() error {
	return core.Remove(v)
}

// Shutdown powers off the VM
func (v *VirtualMachine) Shutdown() (err error) {
	return NewTask(core.Put(
		v.Rel("state").Href,
		"acceptedrequest",
		"virtualmachinestate",
		&VirtualMachineState{
			GracefulShutdown: true,
			State:            "OFF",
		},
	))
}

// GetState return v VirtualMachine state
func (v *VirtualMachine) GetState() (state VirtualMachineState, err error) {
	err = core.Read(v.Rel("state"), &state)
	return
}
