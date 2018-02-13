package abiquo

import (
	"fmt"

	. "github.com/abiquo/opal/core"
)

// VirtualMachine represents an Abiquo API VMcore.Resource
type VirtualMachine struct {
	CPU       int               `json:"cpu,omitempty"`
	Label     string            `json:"label,omitempty"`
	Metadata  string            `json:"metadata,omitempty"`
	Monitored bool              `json:"monitored"`
	Name      string            `json:"name,omitempty"`
	RAM       int               `json:"ram,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
	UUID      string            `json:"uuid,omitempty"`
	DTO
}

type VirtualMachineMetadata struct {
	Metadata VirtualMachineMetadataFields `json:"metadata"`
	DTO
}

type VirtualMachineMetadataFields struct {
	StartupScript string `json:"startup-script,omitempty"`
}

func NewVirtualMachine() Resource { return new(VirtualMachine) }

// Reconfigure reconfigures v
func (v *VirtualMachine) Reconfigure() (err error) {
	_, err = Rest(v, Put(v.URL(), "acceptedrequest", v.Media(), v))
	return
}

// Deploy deploys v
func (v *VirtualMachine) Deploy() (err error) {
	return NewTask(v.Rel("deploy"), "virtualmachinetask", v)
}

// Undeploy undeploys v
func (v *VirtualMachine) Undeploy() (err error) {
	return NewTask(v.Rel("undeploy"), "virtualmachinetask", v)
}

// SetMetadata sets the VM metadata as requested
func (v *VirtualMachine) SetMetadata(metadata *VirtualMachineMetadata) error {
	return Update(v.Rel("metadata"), metadata)
}

func filterLinks(l []*Link, filter func(link *Link) bool) (links []*Link) {
	for _, link := range l {
		if filter(link) {
			links = append(links, link)
		}
	}
	return
}

func isLink(l *Link, media string) bool {
	return l.Type == Media(media)
}

// Disks returns an slice with the VM disk links
func (v *VirtualMachine) Disks() []*Link {
	return filterLinks(v.Links, func(l *Link) bool {
		return isLink(l, "harddisk") || isLink(l, "volume")
	})
}

// NICs returns an slice with the VM NIC links
func (v *VirtualMachine) NICs() (nics []*Link) {
	return filterLinks(v.Links, func(l *Link) bool {
		return isLink(l, "privateip") || isLink(l, "externalip") || isLink(l, "publicip")
	})
}

// AttachDisk add a disk link to the *VirtualMachine in the last position
func (v *VirtualMachine) AttachDisk(hd *HardDisk) error {
	diskLink := hd.Link().SetRel(fmt.Sprintf("disk%v", len(v.Disks())))
	diskLink.DiskControllerType = hd.DiskControllerType
	diskLink.DiskController = hd.DiskController
	v.Add(diskLink)
	return nil
}

// AttachNIC add a nic link to the *VirtualMachine in the last position
func (v *VirtualMachine) AttachNIC(nic *Link) error {
	nicLink := nic.SetRel(fmt.Sprintf("nic%v", len(v.NICs())))
	v.Add(nicLink)
	return nil
}

func (v *VirtualMachine) Delete() error {
	return Remove(v)
}
