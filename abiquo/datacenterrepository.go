package abiquo

import (
	"net/url"
	"strings"

	"github.com/abiquo/opal/core"
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

// Upload uploads an OVA to the *DatacenterRepository, and returns the *VirtualMachineTemplate DTO
func (d *DatacenterRepository) Upload(file string) (v *VirtualMachineTemplate, err error) {
	endpoint := d.Rel("applianceManagerRepositoryUri").Href + "/templates"
	reply, err := core.Upload(endpoint, file)
	if err == nil {
		path := strings.Join(strings.Split(reply.Location(), "/")[7:], "/")
		templates := d.Rel("virtualmachinetemplates").Collection(url.Values{"path": {path}})
		if vmt := templates.First(); vmt != nil {
			v = vmt.(*VirtualMachineTemplate)
		}
	}
	return
}
