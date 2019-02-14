package abiquo

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/abiquo/ojal/core"
)

// DatacenterRepository represents an Abiquo API DatacenterRepository DTO
//
// Collections:
// - virtualmachinetemplates
type DatacenterRepository struct {
	Name     string `json:"name"`
	Location string `json:"repositoryLocation"`
	core.DTO
}

// DiskDefinition represents a disk being uploaded
type DiskDefinition struct {
	Bootable           bool   `json:"bootable"`
	DiskController     string `json:"diskController"`
	DiskControllerType string `json:"diskControllerType,omitempty"`
	DiskFileFormat     string `json:"diskFileFormat"`
	DiskFileSize       int    `json:"diskFileSize"`
	DiskFilePath       string `json:"diskFilePath"`
	Label              string `json:"label"`
	RequiredHDinMB     int    `json:"requiredHDInMB"`
	Sequence           int    `json:"sequence"`
}

// TemplateDefinition represents a template being uploaded
type TemplateDefinition struct {
	CategoryName             string           `json:"categoryName"`
	Description              string           `json:"description"`
	Disks                    []DiskDefinition `json:"disks"`
	EthernetDriverType       string           `json:"ethernetDriverType"`
	LoginUser                string           `json:"loginUser,omitempty"`
	LoginPassword            string           `json:"loginPassword,omitempty"`
	Name                     string           `json:"name"`
	NewEmptyDiskCapacityInMB string           `json:"newEmptyDiskCapacityInMb"`
	OSType                   string           `json:"osType,omitempty"`
	OSVersion                string           `json:"osVersion,omitempty"`
	RequiredCPU              int              `json:"requiredCpu"`
	RequiredRAMInMB          int              `json:"requiredRamInMB"`
}

func (d *DatacenterRepository) upload(file, info string) (*VirtualMachineTemplate, error) {
	href := d.Rel("applianceManagerRepositoryUri").Href + "/templates"
	reply, err := core.Upload(href, file, info)
	if err != nil {
		return nil, err
	}

	// Prevent the template from not being found inmediately after the upload
	time.Sleep(5 * time.Second)

	path := strings.Join(strings.Split(reply.URL(), "/")[7:], "/")
	templates := d.Rel("virtualmachinetemplates").Collection(url.Values{"path": {path}})
	resource := templates.First()
	if resource == nil {
		return nil, errors.New("template not found after upload")
	}
	return resource.(*VirtualMachineTemplate), nil
}

// UploadOVA uploads an OVA to the *DatacenterRepository, and returns the *VirtualMachineTemplate DTO
func (d *DatacenterRepository) UploadOVA(file string) (v *VirtualMachineTemplate, err error) {
	return d.upload(file, "")
}

// UploadTemplate uploads an OVA to the *DatacenterRepository, and returns the *VirtualMachineTemplate DTO
func (d *DatacenterRepository) UploadTemplate(file string, definition TemplateDefinition) (v *VirtualMachineTemplate, err error) {
	bytes, err := json.Marshal(definition)
	if err != nil {
		return nil, err
	}
	info := string(bytes)
	return d.upload(file, info)
}
