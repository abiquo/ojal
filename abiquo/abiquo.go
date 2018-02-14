package abiquo

import (
	"sync"

	. "github.com/abiquo/opal/core"
)

var (
	initialize  sync.Once
	collections = map[string]func() Resource{
		"categories":              NewCategory,
		"datacenters":             NewDatacenter,
		"datacenterrepositories":  newDatacenterRepository,
		"enterprises":             NewEnterprise,
		"hardwareprofiles":        newHardwareProfile,
		"ips":                     newIP,
		"licenses":                newLicense,
		"privileges":              newPrivilege,
		"publiccloudregions":      newPublicLocation,
		"roles":                   NewRole,
		"scopes":                  newScope,
		"users":                   NewUser,
		"vlans":                   newNetwork,
		"virtualappliances":       NewVirtualAppliance,
		"virtualdatacenters":      NewVirtualDatacenter,
		"virtualmachines":         NewVirtualMachine,
		"virtualmachinetemplates": NewVirtualMachineTemplate,
	}

	resources = map[string]func() Resource{
		"category":               NewCategory,
		"datacenter":             NewDatacenter,
		"datacenterrepository":   newDatacenterRepository,
		"enterprise":             NewEnterprise,
		"hardwareprofile":        newHardwareProfile,
		"ip":                     newIP,
		"license":                newLicense,
		"privilege":              newPrivilege,
		"publiccloudregion":      newPublicLocation,
		"role":                   NewRole,
		"scope":                  newScope,
		"user":                   NewUser,
		"vlan":                   newNetwork,
		"virtualappliance":       NewVirtualAppliance,
		"virtualdatacenter":      NewVirtualDatacenter,
		"virtualmachine":         NewVirtualMachine,
		"virtualmachinetemplate": NewVirtualMachineTemplate,
	}
)

// Abiquo initializes the Abiquo API client and registers the known collections
func Abiquo(api string, credentials interface{}) (err error) {
	initialize.Do(func() {
		if err = Init(api, credentials); err == nil {
			for media, constructor := range collections {
				RegisterCollection(media, constructor)
			}
			for resource, factory := range resources {
				RegisterResource(resource, factory)
			}
		}
	})
	return
}

// Login returns the User resource for the client credentials
func Login() (user *User) {
	u := new(User)
	if err := Read(NewLinker("login", "user"), u); err == nil {
		user = u
	}
	return
}
