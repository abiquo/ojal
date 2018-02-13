package core_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/abiquo/opal/core"
)

type battery []*struct {
	Message  string
	Current  interface{}
	Expected interface{}
}

func (b battery) Run(name string, t *testing.T) {
	for _, v := range b {
		if v.Current != v.Expected {
			t.Fatalf("%v %v ? %v != %v", name, v.Message, v.Current, v.Expected)
		}
	}
}

type dto struct {
	Name string `json:"name"`
	DTO
}

func newTenant() Resource { return new(dto) }

var (
	none        *Link
	self        *Link
	edit        *Link
	enterprise  Linker
	enterprises Linker
	name0       string
	name1       string
	enterprise0 *dto
	enterprise1 *dto
	result      *dto
)

func TestInit(t *testing.T) {
	basic := Basic{
		Username: "admin",
		Password: "xabiquo",
	}

	oauth := Oauth{
		APIKey:      "5336cd80-d17b-488a-8917-518a12ee366a",
		APISecret:   "nuDmkp1t4qmcyxGVfVsujmVqJ5VexeLIymvBA5Oy",
		Token:       "7ea0959c-82f1-4013-ab2b-6648999f3915",
		TokenSecret: "TgYSC9Y4TX3r+p9q3F8DhcJ3J9FFXOCmPD6pAKw1G31wTUAtlTgZTMJjDT/jS2F4K2DUYX6Py641PLeBkKMntS+GdKkO09ajkil9ZH67Fa0=",
	}

	err0 := Init("https://missing:443/api", Basic{})
	err1 := Init("https://testing:443/api", Basic{})
	err2 := Init("https://testing:443/api", oauth)
	err3 := Init("https://testing:443/api", basic)

	battery{
		{"err0", err0 == nil, false},
		{"err1", err1 == nil, false},
		{"err2", err2 == nil, true},
		{"err3", err3 == nil, true},
		{"Version()", Version(), "4.2"},
	}.Run("Init", t)

	ts := time.Now().Unix()
	name0 = fmt.Sprint("ztest A ", ts)
	name1 = fmt.Sprint("ztest B ", ts)
	none = NewLinkType("none", "none").SetRel("none").SetTitle("none")
	self = NewLinkType("self", "self").SetRel("self").SetTitle("self")
	edit = NewLinkType("edit", "edit").SetRel("edit").SetTitle("edit")
	enterprise = NewLinker("admin/enterprises", "enterprise")
	enterprises = NewLinker("admin/enterprises", "enterprises")
	enterprise0 = &dto{Name: name0}
	enterprise1 = &dto{Name: name1}
	result = &dto{}
	RegisterResource("enterprise", newTenant)
	RegisterResource("user", newTenant)
	RegisterCollection("enterprises", newTenant)
	RegisterCollection("user", newTenant)
}

func ExampleLink() {
	fmt.Println(none.URL())
	fmt.Println(none.Rel)
	fmt.Println(none.Title)
	fmt.Println(none.Media())

	// Output:
	// https://testing:443/api/none
	// none
	// none
	// application/vnd.abiquo.none+json
}

func TestCall(t *testing.T) {
	post, err := Rest(result, Post(
		"admin/enterprises",
		"enterprise",
		"enterprise",
		enterprise0,
	))
	battery{
		{"err", err, nil},
		{"Ok()", post.Ok(), true},
		{"Status()", post.Status(), http.StatusCreated},
		{"enterprise.name", result.Name, name0},
	}.Run("post", t)
	href := post.Location()

	put, err := Rest(result, Put(
		href,
		"enterprise",
		"enterprise",
		enterprise1,
	))
	battery{
		{"err", err, nil},
		{"Ok()", put.Ok(), true},
		{"Status()", put.Status(), http.StatusOK},
		{"enterprise.Name", result.Name, name1},
	}.Run("put", t)

	get, err := Rest(result, Get(href, "enterprise"))
	battery{
		{"err", err, nil},
		{"Ok()", get.Ok(), true},
		{"Status()", get.Status(), http.StatusOK},
		{"enterprise.Name", result.Name, name1},
	}.Run("get", t)

	delete1, err := Rest(nil, Delete(href))
	battery{
		{"err", err, nil},
		{"Ok()", delete1.Ok(), true},
		{"Status()", delete1.Status(), http.StatusNoContent},
	}.Run("delete1", t)

	delete2, err := Rest(nil, Delete(href))
	battery{
		{"err", err == nil, false},
		{"Ok()", delete2.Ok(), false},
		{"Status()", delete2.Status(), http.StatusNotFound},
	}.Run("delete2", t)

	values := url.Values{"idDatacenter": {"1"}}
	query, err := Rest(nil, Get("admin/rules", "rules").Query(values))
	battery{
		{"err", err, nil},
		{"Ok()", query.Ok(), true},
		{"Status()", query.Status(), http.StatusOK},
		//		{"call.href", call.href, "https://testing:443/api/admin/rules?idDatacenter=1"},
	}.Run("query", t)
}

func TestDTO(t *testing.T) {
	dto := NewDTO(none)
	battery{
		{"Href()", dto.URL(), ""},
		{"Type()", dto.Media(), Media("")},
	}.Run("none", t)

	dto.Add(self)
	battery{
		{"Href()", dto.URL(), self.URL()},
		{"Type()", dto.Media(), self.Media()},
	}.Run("self", t)

	dto.Add(edit)
	battery{
		{"Href()", dto.URL(), edit.URL()},
		{"Type()", dto.Media(), edit.Media()},
	}.Run("edit", t)
}

func ExampleHref() {
	fmt.Println(Resolve("", nil))
	fmt.Println(Resolve("admin/rules", nil))
	fmt.Println(Resolve("admin/rules", url.Values{"idDatacenter": {"1"}}))

	// Output:
	// https://testing:443/api/
	// https://testing:443/api/admin/rules
	// https://testing:443/api/admin/rules?idDatacenter=1
}

func ExampleType() {
	fmt.Println(Media(""))
	fmt.Println(Media("text/plain"))
	fmt.Println(Media("enterprise"))

	// Output:
	//
	// text/plain
	// application/vnd.abiquo.enterprise+json
}

func ExampleUpload() {
	ova := "/home/antxon/Downloads/test.ova"
	templates := "https://testing:443/am/erepos/1/templates"
	reply, err := Upload(templates, ova)

	fmt.Println(err)
	fmt.Println(reply.Ok())
	fmt.Println(reply.Location() != "")

	// Output:
	// <nil>
	// true
	// true
}

func TestCollection(t *testing.T) {
	collection := enterprises.Collection(nil)
	count := 0
	for collection.Next() {
		collection.Item()
		count++
	}

	battery{
		{"Size()", count, collection.Size()},
	}.Run("collection", t)
}

func ExampleCollection_List() {
	for _, e := range enterprises.Collection(nil).List() {
		fmt.Println(e.(*dto).Name)
		break
	}

	// Output:
	// Abiquo
}

func ExampleCollection_Find() {
	for _, name := range []string{"Abiquo", "abq"} {
		finder := func(r Resource) bool {
			return r.(*dto).Name == name
		}
		result := enterprises.Collection(nil).Find(finder)
		fmt.Println(result == nil)
	}

	// Output:
	// false
	// true
}

func ExampleCollection_First() {
	e := enterprises.Collection(nil).First()
	fmt.Println(e.(*dto).Name)

	// Output:
	// Abiquo
}

func ExampleCrud() {
	var (
		create1 = Create(enterprise, enterprise0)
		create2 = Create(enterprise, enterprise0)
		read    = Read(enterprise0, enterprise0)
		update  = Update(enterprise0, enterprise0)
		remove1 = Remove(enterprise0)
		remove2 = Remove(enterprise0)
	)
	fmt.Println(create1)
	fmt.Println(create2)
	fmt.Println(read)
	fmt.Println(update)
	fmt.Println(remove1)
	fmt.Println(remove2)

	// Output:
	// <nil>
	// 409 Unexpected status code: ENTERPRISE-4 Duplicate name for an enterprise
	// <nil>
	// <nil>
	// <nil>
	// 404 Unexpected status code: EN-0 The requested enterprise does not exist
}

func TestWalk(t *testing.T) {
	endpoint := NewLinker("admin/enterprises/1/users/1", "user")
	enterprise := endpoint.Walk().Walk("enterprise")

	battery{
		{"nil", enterprise == nil, false},
		{"name", enterprise.(*dto).Name, "Abiquo"},
	}.Run("walk", t)
}

func ExampleResources_Map() {
	var count int
	counter := func(r Resource) { count++ }
	collection := enterprises.Collection(nil)
	collection.List().Map(counter)
	fmt.Println(collection.Size() == count)

	// Output:
	// true
}

func ExampleResources_Filter() {
	filter := func(r Resource) bool { return r.(*dto).Name == "Abiquo" }
	collection := enterprises.Collection(nil).List().Filter(filter)
	fmt.Println(collection[0].(*dto).Name)
	fmt.Println(len(collection))

	// Output:
	// Abiquo
	// 1
}
