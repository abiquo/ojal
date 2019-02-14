package abiquo

import (
	"github.com/abiquo/ojal/core"
)

// Categories returns the platform categories collection
func Categories() *core.Link {
	return core.NewLink("config/categories").SetType("categories")
}

// Datacenters returns the Abiquo API datacenters collection
func Datacenters() *core.Link {
	return core.NewLink("admin/datacenters").SetType("datacenters")
}

// DeviceTypes returns the API supported SDN device types collection
func DeviceTypes() *core.Link {
	return core.NewLink("config/devicetypes").SetType("devicetypes")
}

// Enterprises returns a slice of enterprises
func Enterprises() *core.Link {
	return core.NewLink("admin/enterprises").SetType("enterprises")
}

// Licenses returns the collection of licenses in the platform
func Licenses() *core.Link {
	return core.NewLink("config/licenses").SetType("licenses")
}

// Login returns the User resource for the client credentials
func Login() (user *User, err error) {
	user = new(User)
	err = core.NewLink("login").SetType("user").Read(user)
	return
}

// Privileges retuns the API privileges collection
func Privileges() *core.Link {
	return core.NewLink("config/privileges").SetType("privileges")
}

// RemoteServices returns the Remote Services collection in the platform
func RemoteServices() *core.Link {
	return core.NewLink("admin/remoteservices").SetType("remoteservices")
}

// Roles returns the roles collection in the platform
func Roles() *core.Link {
	return core.NewLink("admin/roles").SetType("roles")
}

// Scopes returns the API scopes collection
func Scopes() *core.Link {
	return core.NewLink("admin/scopes").SetType("scopes")
}

// VirtualDatacenters returns the Abiquo virtual datacenters collection
func VirtualDatacenters() *core.Link {
	return core.NewLink("cloud/virtualdatacenters").SetType("virtualdatacenters")
}
