package abiquo

import (
	"sync"

	"github.com/abiquo/opal/core"
)

var (
	initialize  sync.Once
	collections = map[string]func() core.Resource{
		"categories":              NewCategory,
		"datacenters":             newDatacenter,
		"datacenterrepositories":  newDatacenterRepository,
		"datastoreloadrules":      NewDatastoreLoadRule,
		"datastores":              newDatastore,
		"datastoretiers":          newDatastoreTier,
		"enterprises":             NewEnterprise,
		"fitpolicyrules":          NewFitPolicy,
		"hardwareprofiles":        newHardwareProfile,
		"ips":                     newIP,
		"licenses":                newLicense,
		"machineloadrules":        NewMachineLoadRule,
		"machines":                newMachine,
		"networkservicetypes":     newNetworkServiceType,
		"privileges":              newPrivilege,
		"publiccloudregions":      newPublicLocation,
		"racks":                   newRack,
		"remoteservices":          newRemoteService,
		"roles":                   NewRole,
		"scopes":                  newScope,
		"users":                   NewUser,
		"vlans":                   newNetwork,
		"virtualappliances":       NewVirtualAppliance,
		"virtualdatacenters":      NewVirtualDatacenter,
		"virtualmachines":         NewVirtualMachine,
		"virtualmachinetemplates": NewVirtualMachineTemplate,
	}

	resources = map[string]func() core.Resource{
		"category":               NewCategory,
		"datacenter":             newDatacenter,
		"datacenterrepository":   newDatacenterRepository,
		"datastore":              newDatastore,
		"datastoreloadrule":      NewDatastoreLoadRule,
		"datastoretier":          newDatastoreTier,
		"enterprise":             NewEnterprise,
		"fitpolicyrule":          NewFitPolicy,
		"hardwareprofile":        newHardwareProfile,
		"ip":                     newIP,
		"license":                newLicense,
		"machine":                newMachine,
		"machineloadrule":        NewMachineLoadRule,
		"networkservicetype":     newNetworkServiceType,
		"privilege":              newPrivilege,
		"publiccloudregion":      newPublicLocation,
		"rack":                   newRack,
		"remoteservice":          newRemoteService,
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
		if err = core.Init(api, credentials); err == nil {
			for media, constructor := range collections {
				core.RegisterCollection(media, constructor)
			}
			for resource, factory := range resources {
				core.RegisterResource(resource, factory)
			}
		}
	})
	return
}

// Login returns the User resource for the client credentials
func Login() (user *User) {
	u := new(User)
	if err := core.Read(core.NewLinker("login", "user"), u); err == nil {
		user = u
	}
	return
}
