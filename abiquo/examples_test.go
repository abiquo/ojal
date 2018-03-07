package abiquo_test

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/abiquo/ojal/abiquo"
	"github.com/abiquo/ojal/core"
)

var (
	name        = fmt.Sprint("ztest", -time.Now().Unix())
	environment map[string]string
)

func init() {
	environment = map[string]string{
		"OJAL_ENDPOINT": os.Getenv("OJAL_ENDPOINT"),
		"OJAL_USERNAME": os.Getenv("OJAL_USERNAME"),
		"OJAL_PASSWORD": os.Getenv("OJAL_PASSWORD"),
	}

	for k, v := range environment {
		if v == "" {
			panic(k + " environment variable should not be empty")
		}
	}

	if err := core.Init(environment["OJAL_ENDPOINT"], core.Basic{
		Username: environment["OJAL_USERNAME"],
		Password: environment["OJAL_PASSWORD"],
	}); err != nil {
		panic(err)
	}
}

// ExampleCollection shows how to retrieve the users name from an enterpirse
func ExampleCollection() {
	users := []string{}
	collection := core.NewLinker("admin/enterprises/1/users", "users").Collection(nil)
	for collection.Next() {
		users = append(users, collection.Item().(*abiquo.User).Name)
	}
	sort.Strings(users)
	fmt.Println(users)

	// Output:
	// [Cloud Default User for Outbound API Events Standard]
}

// ExampleCategories shows how to list all the categories
func ExampleCategories() {
	category := abiquo.Categories(nil).Find(func(r core.Resource) bool {
		return r.(*abiquo.Category).Name == "Others"
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
	category := &abiquo.Category{Name: name, Erasable: true}
	fmt.Println(category.Create())
	fmt.Println(core.Remove(category))

	// Output:
	// <nil>
	// <nil>
}

func ExampleLogin() {
	user := abiquo.Login()
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
	location := core.NewLinkType("admin/datacenters/1", "location")
	datacenter := core.NewLinkType("cloud/locations/1", "location")

	enterprise := &abiquo.Enterprise{Name: name}
	err0 := enterprise.Create()
	err1 := enterprise.CreateLimit(&abiquo.Limit{DTO: core.NewDTO(
		enterprise.Link().SetRel("enterprise"),
		location.SetRel("location"),
	)})

	vdc := &abiquo.VirtualDatacenter{
		Name:   name,
		HVType: "KVM",
		Network: &abiquo.Network{
			Mask:    24,
			Address: "192.168.0.0",
			Gateway: "192.168.0.1",
			Name:    name,
			TypeNet: "INTERNAL",
		},
		DTO: core.NewDTO(
			datacenter.SetRel("location"),
			enterprise.Link().SetRel("enterprise"),
		),
	}
	err2 := vdc.Create()
	err3 := vdc.Network.CreateIP(&abiquo.IP{IP: "192.168.0.253"})

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
	dc := new(abiquo.Datacenter)
	endpoint := core.NewLinkType("admin/datacenters/1", "datacenter")
	read := core.Read(endpoint, dc)
	network := &abiquo.Network{
		Mask:    24,
		Address: "172.16.45.0",
		Gateway: "172.16.45.1",
		Name:    name,
		Tag:     3743,
		TypeNet: "EXTERNAL",
		DTO: core.NewDTO(
			core.NewLinkType("admin/enterprises/1", "enterprise").SetRel("enterprise"),
			core.NewLinkType("admin/datacenters/1/networkservicetypes/1", "networkservicetype").SetRel("networkservicetype"),
		),
	}
	fmt.Println(read)
	fmt.Println(dc.CreateExternal(network))
	fmt.Println(core.Remove(network))

	// Output:
	// <nil>
	// <nil>
	// <nil>
}

func ExampleVirtualMachine_Deploy() {
	endpoint := core.NewLinkType("cloud/virtualdatacenters/1/virtualappliances/1", "virtualappliance")
	template := core.NewLinkType("admin/enterprises/1/datacenterrepositories/1/virtualmachinetemplates/1", "virtualmachinetemplate")
	vapp := endpoint.Walk().(*abiquo.VirtualAppliance)
	vm := &abiquo.VirtualMachine{
		DTO: core.NewDTO(template.SetRel("virtualmachinetemplate")),
	}

	vapp.CreateVM(vm)
	vm.Deploy()
	vm.Undeploy()
	vm.Delete()
}
