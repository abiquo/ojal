package abiquo

import (
	"sync"

	"github.com/abiquo/opal/core"
)

var (
	initialize sync.Once
	collection = map[string]string{
		"category":               "categories",
		"datacenter":             "datacenters",
		"datacenterrepository":   "datacenterrepositories",
		"datastoreloadrule":      "datastoreloadrules",
		"datastore":              "datastores",
		"datastoretier":          "datastoretiers",
		"enterprise":             "enterprises",
		"fitpolicyrule":          "fitpolicyrules",
		"hardwareprofile":        "hardwareprofiles",
		"ip":                     "ips",
		"license":                "licenses",
		"machineloadrule":        "machineloadrules",
		"machine":                "machines",
		"networkservicetype":     "networkservicetypes",
		"privilege":              "privileges",
		"publiccloudregion":      "publiccloudregions",
		"rack":                   "racks",
		"remoteservice":          "remoteservices",
		"role":                   "roles",
		"scope":                  "scopes",
		"user":                   "users",
		"vlan":                   "vlans",
		"virtualappliance":       "virtualappliances",
		"virtualdatacenter":      "virtualdatacenters",
		"virtualmachine":         "virtualmachines",
		"virtualmachinetemplate": "virtualmachinetemplates",
	}

	resource = map[string]func() core.Resource{
		"category":               func() core.Resource { return new(Category) },
		"datacenter":             func() core.Resource { return new(Datacenter) },
		"datacenterrepository":   func() core.Resource { return new(DatacenterRepository) },
		"datastore":              func() core.Resource { return new(Datastore) },
		"datastoreloadrule":      func() core.Resource { return new(DatastoreLoadRule) },
		"datastoretier":          func() core.Resource { return new(DatastoreTier) },
		"enterprise":             func() core.Resource { return new(Enterprise) },
		"fitpolicyrule":          func() core.Resource { return new(FitPolicy) },
		"hardwareprofile":        func() core.Resource { return new(HardwareProfile) },
		"ip":                     func() core.Resource { return new(IP) },
		"license":                func() core.Resource { return new(License) },
		"machine":                func() core.Resource { return new(Machine) },
		"machineloadrule":        func() core.Resource { return new(MachineLoadRule) },
		"networkservicetype":     func() core.Resource { return new(NetworkServiceType) },
		"privilege":              func() core.Resource { return new(Privilege) },
		"publiccloudregion":      func() core.Resource { return new(Location) },
		"rack":                   func() core.Resource { return new(Rack) },
		"remoteservice":          func() core.Resource { return new(RemoteService) },
		"role":                   func() core.Resource { return new(Role) },
		"scope":                  func() core.Resource { return new(Scope) },
		"user":                   func() core.Resource { return new(User) },
		"vlan":                   func() core.Resource { return new(Network) },
		"virtualappliance":       func() core.Resource { return new(VirtualAppliance) },
		"virtualdatacenter":      func() core.Resource { return new(VirtualDatacenter) },
		"virtualmachine":         func() core.Resource { return new(VirtualMachine) },
		"virtualmachinetemplate": func() core.Resource { return new(VirtualMachineTemplate) },
	}
)

// Abiquo initializes the Abiquo API client and registers the known collections
func Abiquo(api string, credentials interface{}) (err error) {
	initialize.Do(func() {
		if err = core.Init(api, credentials); err != nil {
			return
		}
		for media, factory := range resource {
			core.RegisterMedia(media, collection[media], factory)
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
