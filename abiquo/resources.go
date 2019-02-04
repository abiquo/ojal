package abiquo

import "github.com/abiquo/ojal/core"

// Disk is the shared volume/harddisk interface
type Disk interface {
	Link() *core.Link
	Controller() string
	ControllerType() string
}

// HardDisk represents an hard disk resource
type HardDisk struct {
	ID                 int    `json:"id,omitempty"`
	Label              string `json:"label,omitempty"`
	Bootable           bool   `json:"bootable,omitempty"`
	DiskController     string `json:"diskController,omitempty"`
	DiskControllerType string `json:"diskControllerType,omitempty"`
	Sequence           int    `json:"sequence,omitempty"`
	DiskFormatType     string `json:"diskFormatType,omitempty"`
	DiskFileSize       int    `json:"diskFileSize,omitempty"`
	SizeInMb           int    `json:"sizeInMb,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	core.DTO
}

// Controller ...
func (h *HardDisk) Controller() string { return h.DiskController }

// ControllerType ...
func (h *HardDisk) ControllerType() string { return h.DiskControllerType }

// Tier represents a tier resource
type Tier struct {
	ID                      int    `json:"id,omitempty"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Enabled                 bool   `json:"enabled"`
	DefaultAllowed          bool   `json:"defaultAllowed"`
	StorageAllocationPolicy string `json:"storageAllocationPolicy"`
	core.DTO
}

// Volume represents a volume resource
type Volume struct {
	ID                 int    `json:"id,omitempty"`
	AllowResize        bool   `json:"allowResize"`
	SizeInMB           int    `json:"sizeInMB,omitempty"`
	State              string `json:"state,omitempty"`
	Name               string `json:"name,omitempty"`
	Bootable           bool   `json:"bootable,omitempty"`
	Description        string `json:"description,omitempty"`
	DiskControllerType string `json:"diskControllerType,omitempty"`
	DiskController     string `json:"diskController"`
	core.DTO
}

// Controller ...
func (v *Volume) Controller() string { return v.DiskController }

// ControllerType ...
func (v *Volume) ControllerType() string { return v.DiskControllerType }
