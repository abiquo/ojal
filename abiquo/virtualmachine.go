package abiquo

import (
	"fmt"
	"time"

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

// AttachDisk add a disk link to the *VirtualMachine in the last position
func (v *VirtualMachine) AttachDisk(disk Disk) {
	diskLink := disk.Link().SetRel(fmt.Sprintf("disk%v", len(v.Disks())))
	diskLink.DiskControllerType = disk.ControllerType()
	diskLink.DiskController = disk.Controller()
	v.Add(diskLink)
}

// AttachNIC add a nic link to the *VirtualMachine in the last position
func (v *VirtualMachine) AttachNIC(nic *core.Link) {
	nicLink := nic.SetRel(fmt.Sprintf("nic%v", len(v.NICs())))
	v.Add(nicLink)
}

// Delete removes an existing VirtualMachine from the API
func (v *VirtualMachine) Delete() (err error) {
	call := core.Delete(v)
	switch v.State {
	case "NOT_ALLOCATED":
		_, err = core.Rest(nil, call)
	default:
		err = NewTask(core.Delete(v))
		err, ok := err.(core.Error)
		if ok && err.Collection[0].Code == "VM-1" {
			return nil
		}
	}

	return
}

// DetachDisk ...
func (v *VirtualMachine) DetachDisk(disk Disk) {
	var diskLink *core.Link
	for i, l := range v.Links {
		if disk.URL() == l.URL() {
			diskLink = l
			v.Links[i] = v.Links[len(v.Links)-1]
			v.Links = v.Links[:len(v.Links)-1]
			break
		}
	}

	if diskLink == nil {
		return
	}

	index := func(l *core.Link) (i int) {
		fmt.Sscanf(l.Rel, "disk%d", &i)
		return
	}

	for _, l := range v.Links {
		if l.DiskControllerType != "" && index(l) < index(disk.Link()) {
			l.Rel = fmt.Sprintf("disk%d", index(disk.Link())-1)
		}
	}
}

// Deploy deploys v
func (v *VirtualMachine) Deploy() (err error) {
	return NewTask(core.Post(v.Rel("deploy"), acceptedRequest, core.Media("*/*"), nil))
}

// Disks returns an slice with the VM disk links
func (v *VirtualMachine) Disks() core.Links {
	return v.Links.Filter(func(l *core.Link) bool {
		return l.IsMedia("harddisk") || l.IsMedia("volume")
	})
}

// NICs returns an slice with the VM NIC links
func (v *VirtualMachine) NICs() core.Links {
	return v.Links.Filter(func(l *core.Link) bool {
		return l.IsMedia("privateip") || l.IsMedia("externalip") || l.IsMedia("publicip")
	})
}

// Off powers off the VM
func (v *VirtualMachine) Off() (err error) {
	return NewTask(core.Put(
		v.Rel("state"),
		acceptedRequest,
		virtualMachineState,
		&VirtualMachineState{
			State: "OFF",
		},
	))
}

// On powers on the VM
func (v *VirtualMachine) On() (err error) {
	return NewTask(core.Put(
		v.Rel("state"),
		acceptedRequest,
		virtualMachineState,
		&VirtualMachineState{
			State: "ON",
		},
	))
}

// Reboot the VM
func (v *VirtualMachine) Reboot() (err error) {
	return NewTask(core.Post(v.Rel("reset"), acceptedRequest, core.Media("application/json"), nil))
}

// Reconfigure reconfigures v
func (v *VirtualMachine) Reconfigure() (err error) {
	call := core.Put(v, acceptedRequest, v, v)
	switch v.State {
	case "NOT_ALLOCATED":
		_, err = core.Rest(nil, call)
	default:
		err = NewTask(call)
	}

	return
}

// SetVariables ...
func (v *VirtualMachine) SetVariables(variables map[string]string) (err error) {
	_, err = core.Rest(nil, core.Put(v, v, v, map[string]interface{}{
		"variables": variables,
	}))
	return
}

// SetVMMetadata sets the VM metadata as requested
func (v *VirtualMachine) SetVMMetadata(metadata *VirtualMachineMetadata) error {
	return v.Rel("metadata").Update(metadata)
}

// Shutdown powers off the VM
func (v *VirtualMachine) Shutdown() (err error) {
	return NewTask(core.Put(
		v.Rel("state"),
		acceptedRequest,
		virtualMachineState,
		&VirtualMachineState{
			GracefulShutdown: true,
			State:            "OFF",
		},
	))
}

// Synchronized ...
func (v *VirtualMachine) Synchronized() (done bool, err error) {
	tmp := new(VirtualMachine)
	err = v.Read(tmp)
	if err != nil {
		return
	}

	if v.LastSynchronize < tmp.LastSynchronize {
		done = true
		*v = *tmp
	}
	return
}

// Synchronize ...
func (v *VirtualMachine) Synchronize() (err error) {
	for {
		time.Sleep(5 * time.Second)
		done, err := v.Synchronized()
		if err != nil || done {
			return err
		}
	}
}

// Undeploy undeploys v
func (v *VirtualMachine) Undeploy(vmt *VirtualMachineTask) (err error) {
	return NewTask(core.Post(
		v.Rel("undeploy"),
		acceptedRequest,
		virtualMachineTask,
		vmt,
	))
}

// VirtualMachineTask ...
type VirtualMachineTask struct {
	ForceVDCLimits            bool `json:"forceVdcLimits,omitempty"`
	ForceUndeploy             bool `json:"forceUndeploy,omitempty"`
	ForceEnterpriseSoftLimits bool `json:"forceEnterpriseSoftLimits,omitempty"`
}
