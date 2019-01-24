package abiquo

import (
	"net/url"

	"github.com/abiquo/ojal/core"
)

// Categories returns the platform categories collection
func Categories(query url.Values) *core.Collection {
	return core.NewLinker("config/categories", "categories").Collection(query)
}

// Create creates a new Category in the Abiquo API
func (c *Category) Create() error {
	return core.Create(core.NewLinker("config/categories", "category"), c)
}

// Datacenters returns the Abiquo API datacenters collection
func Datacenters(query url.Values) *core.Collection {
	return core.NewLinker("admin/datacenters", "datacenters").Collection(query)
}

// DeviceTypes returns the API supported SDN device types collection
func DeviceTypes(q url.Values) *core.Collection {
	return core.NewLinker("config/devicetypes", "devicetypes").Collection(q)
}

// Enterprises returns a slice of enterprises
func Enterprises(query url.Values) *core.Collection {
	return core.NewLinker("admin/enterprises", "enterprises").Collection(query)
}

// Create creates the requested enterprise
func (e *Enterprise) Create() error {
	return core.Create(core.NewLinker("admin/enterprises", "enterprise"), e)
}

// Licenses returns the collection of licenses in the platform
func Licenses(query url.Values) *core.Collection {
	return core.NewLinker("config/licenses", "licenses").Collection(query)
}

// Login returns the User resource for the client credentials
func Login() (user *User, err error) {
	resource, err := core.NewLinker("login", "user").Walk()
	if err != nil {
		return
	}
	user = resource.(*User)
	return
}

// VMs returns a load balancer node list
func (l *LoadBalancer) VMs() (vms core.Links, err error) {
	resource, err := l.Walk("virtualmachines")
	if err != nil {
		return
	}
	vms = resource.(*core.DTO).Links
	return
}

// Privileges retuns the API privileges collection
func Privileges(query url.Values) *core.Collection {
	return core.NewLinker("config/privileges", "privileges").Collection(query)
}

// RemoteServices returns the Remote Services collection in the platform
func RemoteServices(query url.Values) *core.Collection {
	return core.NewLinker("admin/remoteservices", "remoteservices").Collection(query)
}

// Roles returns the roles collection in the platform
func Roles(query url.Values) *core.Collection {
	return core.NewLinker("admin/roles", "roles").Collection(query)
}

// Create posts the *Role r to the Abiquo API roles endpoint
func (r *Role) Create() error {
	return core.Create(core.NewLinker("admin/roles", "role"), r)
}

// AddPrivilege adds the *Privilege rel privilege link to the *Role
func (r *Role) AddPrivilege(p *Privilege) {
	r.Add(p.Rel("privilege"))
}

// Create posts the *Scope s to the scopes endpoint
func (s *Scope) Create() error {
	return core.Create(core.NewLinker("admin/scopes", "scope"), s)
}

// Scopes returns the API scopes collection
func Scopes(query url.Values) *core.Collection {
	return core.NewLinker("admin/scopes", "scopes").Collection(query)
}

// VirtualDatacenters returns the Abiquo virtual datacenters collection
func VirtualDatacenters(query url.Values) *core.Collection {
	return core.NewLinker("cloud/virtualdatacenters", "virtualdatacenters").Collection(query)
}

// Create creates a new VDC
func (v *VirtualDatacenter) Create() error {
	endpoint := core.NewLinker("cloud/virtualdatacenters", "virtualdatacenter")
	return core.Create(endpoint, v)
}
