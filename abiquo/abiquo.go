package abiquo

import (
	"github.com/abiquo/ojal/core"
)

func init() {
	collections := map[string]string{
		"backuppolicy":           "backuppolicies",
		"category":               "categories",
		"datacenter":             "datacenters",
		"datacenterrepository":   "datacenterrepositories",
		"datastoreloadrule":      "datastoreloadrules",
		"datastore":              "datastores",
		"datastoretier":          "datastoretiers",
		"device":                 "devices",
		"devicetype":             "devicetypes",
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

	resources := map[string]func() core.Resource{
		"backuppolicy":           func() core.Resource { return new(BackupPolicy) },
		"category":               func() core.Resource { return new(Category) },
		"datacenter":             func() core.Resource { return new(Datacenter) },
		"datacenterrepository":   func() core.Resource { return new(DatacenterRepository) },
		"datastore":              func() core.Resource { return new(Datastore) },
		"datastoreloadrule":      func() core.Resource { return new(DatastoreLoadRule) },
		"datastoretier":          func() core.Resource { return new(DatastoreTier) },
		"device":                 func() core.Resource { return new(Device) },
		"devicetype":             func() core.Resource { return new(DeviceType) },
		"enterprise":             func() core.Resource { return new(Enterprise) },
		"enterpriseproperties":   func() core.Resource { return new(EnterpriseProperties) },
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

	for media, factory := range resources {
		core.RegisterMedia(media, collections[media], factory)
	}
}

// Login returns the User resource for the client credentials
func Login() (user *User) {
	if resource := core.NewLinker("login", "user").Walk(); resource != nil {
		user = resource.(*User)
	}
	return
}
