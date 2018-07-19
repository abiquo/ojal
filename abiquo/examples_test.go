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
	location    core.Resource
	datacenter  core.Resource
	enterprise  core.Resource
)

func init() {
	environment = map[string]string{
		"ABQ_ENDPOINT": os.Getenv("ABQ_ENDPOINT"),
		"ABQ_USERNAME": os.Getenv("ABQ_USERNAME"),
		"ABQ_PASSWORD": os.Getenv("ABQ_PASSWORD"),
	}

	for k, v := range environment {
		if v == "" {
			panic(k + " environment variable should not be empty")
		}
	}

	if err := core.Init(environment["ABQ_ENDPOINT"], core.Basic{
		Username: environment["ABQ_USERNAME"],
		Password: environment["ABQ_PASSWORD"],
	}); err != nil {
		panic(err)
	}

	if location = abiquo.Datacenters(nil).Find(func(r core.Resource) bool {
		return r.Link().Title == "datacenter 1"
	}); location == nil {
		panic("location not found")
	}

	if datacenter = abiquo.PrivateLocations(nil).Find(func(r core.Resource) bool {
		return r.Link().Title == "datacenter 1"
	}); datacenter == nil {
		panic("datacenter not found")
	}

	if enterprise = abiquo.Enterprises(nil).Find(func(r core.Resource) bool {
		return r.Link().Title == "Abiquo"
	}); enterprise == nil {
		panic("enterprise not found")
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
	enterprise := user.Rel("enterprise").Walk()

	fmt.Println(user == nil)
	fmt.Println(enterprise == nil)
	fmt.Println(user.Name)
	fmt.Println(enterprise.Link().Title)

	// Output:
	// false
	// false
	// Cloud
	// Abiquo
}

func ExampleNetwork() {
	enterprise := &abiquo.Enterprise{Name: name}
	err0 := enterprise.Create()

	endpoint := enterprise.Rel("limits").SetType("limit")
	err1 := core.Create(endpoint, &abiquo.Limit{
		DTO: core.NewDTO(
			enterprise.Link().SetRel("enterprise"),
			location.Link().SetRel("location"),
		)},
	)

	vdc := &abiquo.VirtualDatacenter{
		Name:   name,
		HVType: "KVM",
		Network: &abiquo.Network{
			Mask:    24,
			Address: "192.168.0.0",
			Gateway: "192.168.0.1",
			Name:    name,
			Type:    "INTERNAL",
		},
		DTO: core.NewDTO(
			datacenter.Rel("location"),
			enterprise.Link().SetRel("enterprise"),
		),
	}
	err2 := vdc.Create()

	endpoint = vdc.Network.Rel("ips").SetType("privateip")
	err3 := core.Create(endpoint, &abiquo.IP{
		IP: "192.168.0.253",
	})

	err4 := core.Remove(vdc)
	err5 := core.Remove(enterprise)

	fmt.Println(err0)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(err4)
	fmt.Println(err5)

	// Output:
	// <nil>
	// <nil>
	// <nil>
	// <nil>
	// <nil>
	// <nil>
}

// ExampleDatacenter shows the Datacenter functionality
func ExampleDatacenter() {
	nst := datacenter.Rel("networkservicetypes").Collection(nil).Find(func(r core.Resource) bool {
		return r.Link().Title == "Service Network"
	})
	external := &abiquo.Network{
		Mask:    24,
		Address: "172.16.45.0",
		Gateway: "172.16.45.1",
		Name:    name,
		Tag:     3743,
		Type:    "EXTERNAL",
		DTO: core.NewDTO(
			enterprise.Link().SetRel("enterprise"),
			nst.Link().SetRel("networkservicetype"),
		),
	}

	endpoint := datacenter.Rel("network").SetType("vlan")
	err0 := core.Create(endpoint, external)
	err1 := core.Remove(external)

	fmt.Println(err0)
	fmt.Println(err1)

	// Output:
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

	endpoint = vapp.Rel("virtualmachines").SetType("virtualmachine")
	core.Create(endpoint, vm)
	vm.Deploy()
	vm.Undeploy()
	vm.Delete()
}

// ExampleEnterprise show all enterprises names
func ExampleEnterprise() {
	for _, e := range abiquo.Enterprises(nil).List() {
		fmt.Println(e.URL())
	}
}
