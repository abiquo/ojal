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
}

// ExampleUser shows how to retrieve the users name from an enterprise
func ExampleUser() {
	users := []string{}
	collection := core.NewLink("admin/enterprises/1/users").SetType("users").Collection(nil)
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
	category := abiquo.Categories().Collection(nil).Find(func(r core.Resource) bool {
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
	err0 := abiquo.Categories().SetType("category").Create(category)
	err1 := category.Remove()

	fmt.Println(err0)
	fmt.Println(err1)

	// Output:
	// <nil>
	// <nil>
}

func ExampleLogin() {
	user, err1 := abiquo.Login()
	enterprise, err2 := user.Rel("enterprise").Walk()

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(user == nil)
	fmt.Println(enterprise == nil)
	fmt.Println(user.Name)
	fmt.Println(enterprise.Link().Title)

	// Output:
	// <nil>
	// <nil>
	// false
	// false
	// Cloud
	// Abiquo
}

func ExampleNetwork() {
	find := func(link *core.Link, find func(core.Resource) bool) core.Resource {
		return link.Collection(nil).Find(find)
	}

	datacenter := find(abiquo.PrivateLocations(), func(r core.Resource) bool {
		return r.Link().Title == "datacenter 1"
	})

	location := find(abiquo.Datacenters(), func(r core.Resource) bool {
		return r.Link().Title == "datacenter 1"
	})

	enterprise := &abiquo.Enterprise{Name: name}
	err0 := abiquo.Enterprises().SetType("enterprise").Create(enterprise)

	endpoint := enterprise.Rel("limits").SetType("limit")
	err1 := endpoint.Create(&abiquo.Limit{
		DTO: core.NewDTO(
			enterprise.Link().SetRel("enterprise"),
			location.Link().SetRel("location"),
		),
	})

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
	err2 := abiquo.VirtualDatacenters().SetType("virtualdatacenter").Create(vdc)
	err3 := vdc.Network.Rel("ips").SetType("privateip").Create(&abiquo.IP{
		IP: "192.168.0.253",
	})

	fmt.Println(err0)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(vdc.Remove())
	fmt.Println(enterprise.Remove())

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
	find := func(link *core.Link, find func(core.Resource) bool) core.Resource {
		return link.Collection(nil).Find(find)
	}

	enterprise := find(abiquo.Enterprises(), func(r core.Resource) bool {
		return r.Link().Title == "Abiquo"
	})

	datacenter := find(abiquo.PrivateLocations(), func(r core.Resource) bool {
		return r.Link().Title == "datacenter 1"
	})

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
	err0 := endpoint.Create(external)
	err1 := external.Remove()

	fmt.Println(err0)
	fmt.Println(err1)

	// Output:
	// <nil>
	// <nil>
}

func ExampleVirtualMachine_Deploy() {
	find := func(link *core.Link) core.Resource {
		return link.Collection(nil).Find(func(r core.Resource) bool {
			return r.Link().Title == "tests"
		})
	}

	vdc := find(abiquo.VirtualDatacenters())
	template := find(vdc.Rel("templates"))
	vapp := find(vdc.Rel("virtualappliances"))

	vm := &abiquo.VirtualMachine{
		DTO: core.NewDTO(
			template.Link().SetRel("virtualmachinetemplate"),
		),
	}
	err0 := vapp.Rel("virtualmachines").SetType("virtualmachine").Create(vm)
	err1 := vm.Deploy()
	err2 := vm.Undeploy(&abiquo.VirtualMachineTask{
		ForceUndeploy: true,
	})
	err3 := vm.Delete()

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

// ExampleEnterprise show all enterprises names
func ExampleEnterprise() {
	for _, e := range abiquo.Enterprises().Collection(nil).List() {
		fmt.Println(e.URL())
	}
}
