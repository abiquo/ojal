package abiquo_test

import (
	"fmt"
	"sort"
	"time"

	. "github.com/abiquo/opal/abiquo"
	. "github.com/abiquo/opal/core"
)

var name = fmt.Sprint("ztest", -time.Now().Unix())

func ExampleInit() {
	err := Abiquo("https://testing:443/api", Basic{
		Username: "admin",
		Password: "xabiquo",
	})

	fmt.Println(err)
	fmt.Println()

	// Output:
	// <nil>
}

// ExampleCollection shows how to retrieve the users name from an enterpirse
func ExampleCollection() {
	users := []string{}
	collection := NewLinker("admin/enterprises/1/users", "users").Collection(nil)
	for collection.Next() {
		users = append(users, collection.Item().(*User).Name)
	}
	sort.Strings(users)
	fmt.Println(users)

	// Output:
	// [Cloud Default User for Outbound API Events Standard]
}

// ExampleCategories shows how to list all the categories
func ExampleCategories() {
	category := Categories(nil).Find(func(r Resource) bool {
		return r.(*Category).Name == "opal"
	})
	fmt.Println(category != nil)

	// JIRA-10108
	// categoriesSize := categories.Size()
	// categoriesLen := len(categories.List())
	// fmt.Println(categoriesSize == categoriesLen)

	// Output:
	// true
}

// ExampleCategory shows how to create a category
func ExampleCategory() {
	category := &Category{Name: name, Erasable: true}
	fmt.Println(category.Create())
	fmt.Println(Remove(category))

	// Output:
	// <nil>
	// <nil>
}

func ExampleLogin() {
	user := Login()
	enterprise := user.Enterprise()

	fmt.Println(user == nil)
	fmt.Println(enterprise == nil)
	fmt.Println(user.Name)
	fmt.Println(enterprise.Name)

	// Output:
	// false
	// false
	// Cloud
	// Abiquo
}

func ExampleNetwork() {
	location := NewLinkType("admin/datacenters/1", "location")
	datacenter := NewLinkType("cloud/locations/1", "location")

	enterprise := &Enterprise{Name: name}
	err0 := enterprise.Create()
	err1 := enterprise.CreateLimit(&Limit{DTO: NewDTO(
		enterprise.Link().SetRel("enterprise"),
		location.SetRel("location"),
	)})

	vdc := &VirtualDatacenter{
		Name:   name,
		HVType: "KVM",
		Network: &Network{
			Mask:    24,
			Address: "192.168.0.0",
			Gateway: "192.168.0.1",
			Name:    name,
			TypeNet: "INTERNAL",
		},
		DTO: NewDTO(
			datacenter.SetRel("location"),
			enterprise.Link().SetRel("enterprise"),
		),
	}
	err2 := vdc.Create()
	err3 := vdc.Network.CreateIP(&IP{IP: "192.168.0.253"})

	fmt.Println(err0)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)

	// Output:
	// <nil>
	// <nil>
	// <nil>
	// <nil>
}

// ExampleDatacenter shows the Datacenter functionality
func ExampleDatacenter() {
	dc := new(Datacenter)
	endpoint := NewLinkType("admin/datacenters/1", "datacenter")
	read := Read(endpoint, dc)
	netsStart := dc.Networks(nil).Size()
	network := &Network{
		Mask:    24,
		Address: "172.16.45.0",
		Gateway: "172.16.45.1",
		Name:    name,
		Tag:     3743,
		TypeNet: "EXTERNAL",
		DTO: NewDTO(
			NewLinkType("admin/enterprises/1", "enterprise").SetRel("enterprise"),
			NewLinkType("admin/datacenters/1/networkservicetypes/1", "networkservicetype").SetRel("networkservicetype"),
		),
	}
	create := dc.CreateExternal(network)
	remove := Remove(network)
	netsEnd := dc.Networks(nil).Size()

	fmt.Println(read)
	fmt.Println(create)
	fmt.Println(remove)
	fmt.Println(netsStart == netsEnd)

	// Output:
	// <nil>
	// <nil>
	// <nil>
	// true
}

func ExampleVirtualMachine_Deploy() {
	endpoint := NewLink("cloud/virtualdatacenters/22/virtualappliances/21")
	template := NewLinkType("admin/enterprises/1/datacenterrepositories/1/virtualmachinetemplates/1", "virtualmachinetemplate")
	vapp := endpoint.SetType("virtualappliance").Walk().(*VirtualAppliance)
	vm := &VirtualMachine{
		DTO: NewDTO(template.SetRel("virtualmachinetemplate")),
	}

	vapp.CreateVM(vm)
	vm.Deploy()
	vm.Undeploy()
	vm.Delete()
}
